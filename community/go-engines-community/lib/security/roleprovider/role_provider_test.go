package roleprovider_test

import (
	"context"
	"slices"
	"testing"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/security/roleprovider"
	mock_mongo "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/mongo"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestRoleProvider_GetValidRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	validRoleNameIDMap := map[string]string{
		"default-role": "default-role-id",
		"role-1":       "role-1-id",
		"role-2":       "role-2-id",
	}

	dbCollection := mock_mongo.NewMockDbCollection(ctrl)
	dbCollection.EXPECT().Aggregate(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, rawPipeline interface{}, _ ...*options.AggregateOptions) (mongo.Cursor, error) {
		pipeline, ok := rawPipeline.([]bson.M)
		if !ok || len(pipeline) == 0 {
			t.Fatal("invalid Aggregate pipeline")
		}

		match, ok := pipeline[0]["$match"].(bson.M)
		if !ok || len(match) == 0 {
			t.Fatal("invalid match stage")
		}

		nameCond, ok := match["name"].(bson.M)
		if !ok || len(nameCond) == 0 {
			t.Fatal("invalid name condition")
		}

		roles, ok := nameCond["$in"].([]string)
		if !ok {
			t.Fatal("invalid $in operator")
		}

		cursor := mock_mongo.NewMockCursor(ctrl)
		cursor.EXPECT().Next(gomock.Any()).Return(true)
		cursor.EXPECT().Decode(gomock.Any()).DoAndReturn(func(doc *struct {
			FoundRoles []string `bson:"found_roles"`
		}) error {
			for _, role := range roles {
				if id, ok := validRoleNameIDMap[role]; ok {
					doc.FoundRoles = append(doc.FoundRoles, id)
				}
			}

			return nil
		})
		cursor.EXPECT().Err().Return(nil)
		cursor.EXPECT().Close(gomock.Any()).Return(nil)

		return cursor, nil
	}).AnyTimes()
	dbCollection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, rawFilter interface{}, _ ...*options.FindOneOptions) mongo.SingleResultHelper {
		filter, ok := rawFilter.(bson.M)
		if !ok || len(filter) == 0 {
			t.Fatal("invalid FindOne filter")
		}

		name, ok := filter["name"].(string)
		if !ok {
			t.Fatal("invalid name condition")
		}

		helper := mock_mongo.NewMockSingleResultHelper(ctrl)
		helper.EXPECT().Decode(gomock.Any()).DoAndReturn(func(doc *struct {
			ID string `bson:"_id"`
		}) error {
			var err error

			if id, ok := validRoleNameIDMap[name]; ok {
				doc.ID = id
			} else {
				err = mongodriver.ErrNoDocuments
			}

			return err
		})

		return helper
	}).AnyTimes()

	dbClient := mock_mongo.NewMockDbClient(ctrl)
	dbClient.EXPECT().Collection(gomock.Any()).Return(dbCollection)

	p := roleprovider.NewRoleProvider(dbClient)

	dataSets := []struct {
		testName       string
		potentialRoles []string
		defaultRole    string
		expectedRoles  []string
		expectedError  bool
	}{
		{
			testName:       "given no potential roles, should return default role",
			potentialRoles: []string{},
			defaultRole:    "default-role",
			expectedRoles:  []string{"default-role-id"},
			expectedError:  false,
		},
		{
			testName:       "given invalid potential role, should return default role",
			potentialRoles: []string{"invalid"},
			defaultRole:    "default-role",
			expectedRoles:  []string{"default-role-id"},
			expectedError:  false,
		},
		{
			testName:       "given invalid potential roles, should return default role",
			potentialRoles: []string{"invalid-1", "invalid-2"},
			defaultRole:    "default-role",
			expectedRoles:  []string{"default-role-id"},
			expectedError:  false,
		},
		{
			testName:       "given valid potential role, should return potential role",
			potentialRoles: []string{"role-1"},
			defaultRole:    "default-role",
			expectedRoles:  []string{"role-1-id"},
			expectedError:  false,
		},
		{
			testName:       "given valid potential roles, should return potential roles",
			potentialRoles: []string{"role-1", "role-2"},
			defaultRole:    "default-role",
			expectedRoles:  []string{"role-1-id", "role-2-id"},
			expectedError:  false,
		},
		{
			testName:       "given valid and invalid potential roles, should return only valid roles",
			potentialRoles: []string{"role-0", "role-1", "role-2", "role-3"},
			defaultRole:    "default-role",
			expectedRoles:  []string{"role-1-id", "role-2-id"},
			expectedError:  false,
		},
		{
			testName:       "given no potential roles and invalid default role, should return error",
			potentialRoles: []string{},
			defaultRole:    "invalid",
			expectedRoles:  nil,
			expectedError:  true,
		},
		{
			testName:       "given invalid potential roles and invalid default role, should return error",
			potentialRoles: []string{"invalid-1", "invalid-2"},
			defaultRole:    "invalid",
			expectedRoles:  nil,
			expectedError:  true,
		},
	}

	for _, dataset := range dataSets {
		t.Run(dataset.testName, func(t *testing.T) {
			roles, err := p.GetValidRoleIDs(ctx, dataset.potentialRoles, dataset.defaultRole)
			if !slices.Equal(roles, dataset.expectedRoles) {
				t.Errorf("expected roles %v, got %v", dataset.expectedRoles, roles)
			}

			if dataset.expectedError && err == nil {
				t.Error("expected error, but got none")
			}

			if !dataset.expectedError && err != nil {
				t.Errorf("expected no error, but got %v", err.Error())
			}
		})
	}
}
