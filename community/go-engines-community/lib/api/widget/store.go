package widget

import (
	"context"
	"errors"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/view"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/utils"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	FindViewIds(ctx context.Context, ids []string) (map[string]string, error)
	FindViewIdByTab(ctx context.Context, tabId string) (string, error)
	GetOneBy(ctx context.Context, id string) (*view.Widget, error)
	Insert(ctx context.Context, r EditRequest) (*view.Widget, error)
	Update(ctx context.Context, r EditRequest) (*view.Widget, error)
	Delete(ctx context.Context, id string) (bool, error)
	Copy(ctx context.Context, widget view.Widget, r CopyRequest) (*view.Widget, error)
	UpdateGridPositions(ctx context.Context, items []EditGridPositionItemRequest) (bool, error)
}

func NewStore(dbClient mongo.DbClient) Store {
	return &store{
		collection:         dbClient.Collection(mongo.WidgetMongoCollection),
		tabCollection:      dbClient.Collection(mongo.ViewTabMongoCollection),
		userPrefCollection: dbClient.Collection(mongo.UserPreferencesMongoCollection),
		filterCollection:   dbClient.Collection(mongo.WidgetFiltersMongoCollection),
	}
}

type store struct {
	collection         mongo.DbCollection
	tabCollection      mongo.DbCollection
	userPrefCollection mongo.DbCollection
	filterCollection   mongo.DbCollection
}

func (s *store) FindViewIds(ctx context.Context, ids []string) (map[string]string, error) {
	results := make([]struct {
		ID   string `bson:"_id"`
		View string `bson:"view"`
	}, 0)
	cursor, err := s.collection.Aggregate(ctx, []bson.M{
		{"$match": bson.M{"_id": bson.M{"$in": ids}}},
		{"$lookup": bson.M{
			"from":         mongo.ViewTabMongoCollection,
			"localField":   "tab",
			"foreignField": "_id",
			"as":           "tab",
		}},
		{"$unwind": bson.M{"path": "$tab"}},
		{"$project": bson.M{
			"view": "$tab.view",
		}},
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	viewIds := make(map[string]string)
	for _, result := range results {
		if result.View != "" {
			viewIds[result.ID] = result.View
		}
	}

	return viewIds, nil
}

func (s *store) FindViewIdByTab(ctx context.Context, tabId string) (string, error) {
	result := struct {
		View string `bson:"view"`
	}{}
	err := s.tabCollection.FindOne(ctx, bson.M{"_id": tabId}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongodriver.ErrNoDocuments) {
			return "", nil
		}
		return "", err
	}

	return result.View, nil
}

func (s *store) GetOneBy(ctx context.Context, id string) (*view.Widget, error) {
	model := &view.Widget{}
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&model)
	if err != nil {
		if errors.Is(err, mongodriver.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return model, nil
}

func (s *store) Insert(ctx context.Context, r EditRequest) (*view.Widget, error) {
	now := types.CpsTime{Time: time.Now()}
	widget := transformEditRequestToModel(r)
	widget.ID = utils.NewID()
	widget.Created = &now
	widget.Updated = &now

	_, err := s.collection.InsertOne(ctx, widget)
	if err != nil {
		return nil, err
	}

	return &widget, nil
}

func (s *store) Update(ctx context.Context, r EditRequest) (*view.Widget, error) {
	oldWidget, err := s.GetOneBy(ctx, r.ID)
	if err != nil || oldWidget == nil {
		return nil, err
	}

	now := types.CpsTime{Time: time.Now()}
	widget := transformEditRequestToModel(r)
	widget.ID = oldWidget.ID
	widget.Created = oldWidget.Created
	widget.Updated = &now
	// Empty InternalParameters to remove from update query.
	widget.InternalParameters = view.InternalParameters{}
	update := bson.M{"$set": widget}

	if oldWidget.Type == view.WidgetTypeJunit &&
		(widget.Type != oldWidget.Type ||
			widget.Parameters.IsAPI != oldWidget.Parameters.IsAPI ||
			widget.Parameters.Directory != oldWidget.Parameters.Directory ||
			widget.Parameters.ReportFileRegexp != oldWidget.Parameters.ReportFileRegexp) {
		update["$unset"] = bson.M{"internal_parameters": ""}
	}

	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": widget.ID}, update)
	if err != nil {
		return nil, err
	}

	return &widget, nil
}

func (s *store) Delete(ctx context.Context, id string) (bool, error) {
	delCount, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	if delCount == 0 {
		return false, nil
	}

	err = s.deleteUserPreferences(ctx, id)
	if err != nil {
		return false, err
	}

	err = s.deleteFilters(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *store) Copy(ctx context.Context, widget view.Widget, r CopyRequest) (*view.Widget, error) {
	now := types.CpsTime{Time: time.Now()}
	id := widget.ID
	widget.ID = utils.NewID()
	widget.Tab = r.Tab
	widget.Author = r.Author
	widget.Created = &now
	widget.Updated = &now

	cursor, err := s.filterCollection.Find(ctx, bson.M{
		"widget": id,
		"user":   bson.M{"$in": bson.A{"", nil}},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	filters := make([]interface{}, 0)
	for cursor.Next(ctx) {
		filter := view.Filter{}
		err := cursor.Decode(&filter)
		if err != nil {
			return nil, err
		}

		newId := utils.NewID()
		if widget.Parameters.MainFilter == filter.ID {
			widget.Parameters.MainFilter = newId
		}

		filter.ID = newId
		filter.Widget = widget.ID
		filter.Author = r.Author
		filter.Created = &now
		filter.Updated = &now
		filters = append(filters, filter)
	}

	_, err = s.collection.InsertOne(ctx, widget)
	if err != nil {
		return nil, err
	}

	if len(filters) > 0 {
		_, err := s.filterCollection.InsertMany(ctx, filters)
		if err != nil {
			return nil, err
		}
	}

	return &widget, nil
}

func (s *store) UpdateGridPositions(ctx context.Context, items []EditGridPositionItemRequest) (bool, error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ID
	}
	widgets := make([]view.Widget, 0, len(items))
	cursor, err := s.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return false, err
	}

	err = cursor.All(ctx, &widgets)
	if err != nil || len(widgets) != len(items) {
		return false, err
	}

	tabId := ""
	for _, w := range widgets {
		if tabId == "" {
			tabId = w.Tab
		} else if tabId != w.Tab {
			return false, ValidationErr{error: errors.New("widgets are related to different view tabs")}
		}
	}

	count, err := s.collection.CountDocuments(ctx, bson.M{"tab": tabId})
	if err != nil {
		return false, err
	}
	if count != int64(len(items)) {
		return false, ValidationErr{error: errors.New("widgets are missing")}
	}

	writeModels := make([]mongodriver.WriteModel, len(widgets))
	for i, item := range items {
		writeModels[i] = mongodriver.NewUpdateOneModel().
			SetFilter(bson.M{"_id": item.ID}).
			SetUpdate(bson.M{"$set": bson.M{"grid_parameters": item.GridParameters}})
	}

	res, err := s.collection.BulkWrite(ctx, writeModels)
	if err != nil {
		return false, err
	}

	return res.MatchedCount > 0, nil
}

func (s *store) deleteUserPreferences(ctx context.Context, widgetID string) error {
	_, err := s.userPrefCollection.DeleteMany(ctx, bson.M{
		"widget": widgetID,
	})

	return err
}

func (s *store) deleteFilters(ctx context.Context, widgetID string) error {
	_, err := s.filterCollection.DeleteMany(ctx, bson.M{
		"widget": widgetID,
	})

	return err
}

func transformEditRequestToModel(request EditRequest) view.Widget {
	return view.Widget{
		Tab:            request.Tab,
		Title:          request.Title,
		Type:           request.Type,
		GridParameters: request.GridParameters,
		Parameters:     request.Parameters,
		Author:         request.Author,
	}
}
