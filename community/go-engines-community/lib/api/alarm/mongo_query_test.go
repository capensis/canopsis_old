package alarm

import (
	"context"
	"fmt"
	"testing"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/author"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/entity/dbquery"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/pagination"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/datetime"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pattern"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/savedpattern"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/view"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	mock_mongo "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/mongo"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenPaginationRequest_ShouldBuildQueryWithLookupsAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.Query{
			Page:     2,
			Limit:    10,
			Paginate: true,
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 10},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAlarmAndPbehaviorInfoPatterns_ShouldBuildQueryWithLookupsAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAlarmPatternWithDurationField_ShouldBuildQueryWithAddFieldsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	durationCond, err := pattern.NewDurationCondition(pattern.ConditionGT, datetime.DurationWithUnit{
		Value: 10,
		Unit:  "m",
	})
	if err != nil {
		panic(err)
	}
	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
					{
						Field:     "v.duration",
						Condition: durationCond,
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	delete(fields, "v.duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{"v.duration": durationField}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
			{"v.duration": bson.M{"$gt": 600}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAlarmPatternWithInfosField_ShouldBuildQueryWithInfosTransformBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
					{
						Field:     "v.infos.info_name",
						FieldType: pattern.FieldTypeInt,
						Condition: pattern.NewIntCondition(pattern.ConditionEqual, 3),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{"v.infos_array": bson.M{"$objectToArray": "$v.infos"}}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
			{"v.infos_array": bson.M{"$elemMatch": bson.M{"v.info_name": bson.M{"$eq": 3}}}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAllPatterns_ShouldBuildQueryWithEntityLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "category",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-category"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"e.category": bson.M{"$eq": "test-category"}},
		}}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		bson.M{"$project": bson.M{"e": 0}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithCategoryFilter_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Category: "test-category",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected,
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"e.enabled": true},
			{"e.category": bson.M{"$eq": "test-category"}},
		}}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithCategoryFilterForResolvedAlarms_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Category: "test-category",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"e.enabled": true},
			{"e.category": bson.M{"$eq": "test-category"}},
		}}},
		bson.M{"$project": bson.M{
			"e": 0,
		}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithInstructionsFilter_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	instructionId := "test-instruction"
	mockCursor := mock_mongo.NewMockCursor(ctrl)
	mockCursor.EXPECT().Next(gomock.Any()).Return(true)
	mockCursor.EXPECT().Next(gomock.Any()).Return(false)
	mockCursor.EXPECT().Close(gomock.Any())
	mockCursor.EXPECT().Decode(gomock.Any()).Do(func(instruction *Instruction) {
		durationCond, err := pattern.NewDurationCondition(pattern.ConditionGT, datetime.DurationWithUnit{
			Value: 10,
			Unit:  "m",
		})
		if err != nil {
			panic(err)
		}

		*instruction = Instruction{
			ActiveOnPbh:   []string{"maintenance"},
			DisabledOnPbh: []string{"pause"},
			AlarmPatternFields: savedpattern.AlarmPatternFields{
				AlarmPattern: pattern.Alarm{
					{
						{
							Field:     "v.duration",
							Condition: durationCond,
						},
						{
							Field:     "v.infos.info_name",
							FieldType: pattern.FieldTypeInt,
							Condition: pattern.NewIntCondition(pattern.ConditionEqual, 3),
						},
					},
				},
			},
			EntityPatternFields: savedpattern.EntityPatternFields{
				EntityPattern: pattern.Entity{
					{
						{
							Field:     "category",
							Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-category"),
						},
					},
				},
			},
		}
	})
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockInstructionDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockInstructionDbCollection.EXPECT().Find(gomock.Any(), gomock.Eq(bson.M{
		"_id":    bson.M{"$in": []string{instructionId}},
		"status": bson.M{"$in": bson.A{InstructionStatusApproved, nil}},
	})).Return(mockCursor, nil)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		case mongo.InstructionMongoCollection:
			return mockInstructionDbCollection
		default:
			return nil
		}
	}).AnyTimes()
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))

	hasRunningExecution := false
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Instructions: []InstructionFilterRequest{
						{
							Running: &hasRunningExecution,
							Include: []string{instructionId},
						},
					},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	delete(fields, "v.duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"v.duration":    durationField,
			"v.infos_array": bson.M{"$objectToArray": "$v.infos"},
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getInstructionExecutionLookup(false)...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"e.enabled": true},
			{"$and": []bson.M{
				{"instruction_execution.instruction": bson.M{"$nin": []string{instructionId}}},
				{"$or": []bson.M{
					{"$and": []bson.M{
						{"$or": []bson.M{{"$and": []bson.M{
							{"v.duration": bson.M{"$gt": 600}},
							{"v.infos_array": bson.M{"$elemMatch": bson.M{"v.info_name": bson.M{"$eq": 3}}}},
						}}}},
						{"$or": []bson.M{{"$and": []bson.M{
							{"e.category": bson.M{"$eq": "test-category"}},
						}}}},
						{"v.pbehavior_info.type": bson.M{"$in": []string{"maintenance"}}},
						{"v.pbehavior_info.type": bson.M{"$nin": []string{"pause"}}},
					}},
				}},
			}},
		}}},
		bson.M{"$project": bson.M{
			"instruction_execution": 0,
		}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithInstructionsFilterForResolvedAlarms_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	instructionId := "test-instruction"
	mockCursor := mock_mongo.NewMockCursor(ctrl)
	mockCursor.EXPECT().Next(gomock.Any()).Return(true)
	mockCursor.EXPECT().Next(gomock.Any()).Return(false)
	mockCursor.EXPECT().Close(gomock.Any())
	mockCursor.EXPECT().Decode(gomock.Any()).Do(func(instruction *Instruction) {
		durationCond, err := pattern.NewDurationCondition(pattern.ConditionGT, datetime.DurationWithUnit{
			Value: 10,
			Unit:  "m",
		})
		if err != nil {
			panic(err)
		}

		*instruction = Instruction{
			ActiveOnPbh:   []string{"maintenance"},
			DisabledOnPbh: []string{"pause"},
			AlarmPatternFields: savedpattern.AlarmPatternFields{
				AlarmPattern: pattern.Alarm{
					{
						{
							Field:     "v.duration",
							Condition: durationCond,
						},
						{
							Field:     "v.infos.info_name",
							FieldType: pattern.FieldTypeInt,
							Condition: pattern.NewIntCondition(pattern.ConditionEqual, 3),
						},
					},
				},
			},
			EntityPatternFields: savedpattern.EntityPatternFields{
				EntityPattern: pattern.Entity{
					{
						{
							Field:     "category",
							Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-category"),
						},
					},
				},
			},
		}
	})
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockInstructionDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockInstructionDbCollection.EXPECT().Find(gomock.Any(), gomock.Eq(bson.M{
		"_id":    bson.M{"$in": []string{instructionId}},
		"status": bson.M{"$in": bson.A{InstructionStatusApproved, nil}},
	})).Return(mockCursor, nil)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		case mongo.InstructionMongoCollection:
			return mockInstructionDbCollection
		default:
			return nil
		}
	}).AnyTimes()
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))

	hasRunningExecution := false
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Instructions: []InstructionFilterRequest{
						{
							Running: &hasRunningExecution,
							Include: []string{instructionId},
						},
					},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	delete(fields, "v.duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"v.duration":    durationField,
			"v.infos_array": bson.M{"$objectToArray": "$v.infos"},
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected, getInstructionExecutionLookup(false)...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"e.enabled": true},
			{"$and": []bson.M{
				{"instruction_execution.instruction": bson.M{"$nin": []string{instructionId}}},
				{"$or": []bson.M{
					{"$and": []bson.M{
						{"$or": []bson.M{{"$and": []bson.M{
							{"v.duration": bson.M{"$gt": 600}},
							{"v.infos_array": bson.M{"$elemMatch": bson.M{"v.info_name": bson.M{"$eq": 3}}}},
						}}}},
						{"$or": []bson.M{{"$and": []bson.M{
							{"e.category": bson.M{"$eq": "test-category"}},
						}}}},
						{"v.pbehavior_info.type": bson.M{"$in": []string{"maintenance"}}},
						{"v.pbehavior_info.type": bson.M{"$nin": []string{"pause"}}},
					}},
				}},
			}},
		}}},
		bson.M{"$project": bson.M{
			"e":                     0,
			"instruction_execution": 0,
		}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithEntitySort_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "name",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-entity"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
			SortRequest: SortRequest{
				MultiSort: []string{
					"entity._id,desc",
					"entity.category.name,asc",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{
			{Key: "e._id", Value: -1},
			{Key: "e.category.name", Value: 1},
			{Key: "_id", Value: 1},
		}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected,
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"e.name": bson.M{"$eq": "test-entity"}},
		}}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
	)
	expected = append(expected, dbquery.GetCategoryLookup("e")...)
	expected = append(expected,
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipelineForResolvedAlarms_GivenRequestWithEntitySort_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "name",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-entity"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
			SortRequest: SortRequest{
				MultiSort: []string{
					"entity._id,desc",
					"entity.category.name,asc",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{
			{Key: "e._id", Value: -1},
			{Key: "e.category.name", Value: 1},
			{Key: "_id", Value: 1},
		}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"e.name": bson.M{"$eq": "test-entity"}},
		}}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
	)
	expected = append(expected, dbquery.GetCategoryLookup("e")...)
	expected = append(expected,
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithDurationSort_ShouldBuildQueryWithAddFieldsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	durationCond, err := pattern.NewDurationCondition(pattern.ConditionGT, datetime.DurationWithUnit{
		Value: 10,
		Unit:  "m",
	})
	if err != nil {
		panic(err)
	}
	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.duration",
						Condition: durationCond,
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
			SortRequest: SortRequest{
				MultiSort: []string{
					"v.duration,desc",
					"v.active_duration,desc",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{
			{Key: "v.duration", Value: -1},
			{Key: "v.active_duration", Value: -1},
			{Key: "_id", Value: 1},
		}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	activeDurationField := fields["v.active_duration"]
	delete(fields, "v.duration")
	delete(fields, "v.active_duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{"v.duration": durationField}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.duration": bson.M{"$gt": 600}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$addFields": bson.M{"v.active_duration": activeDurationField}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearch_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	search := "test-search"
	searchRegexp := primitive.Regex{
		Pattern: fmt.Sprintf(".*%s.*", search),
		Options: "i",
	}
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: search,
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{
			{"v.meta": nil},
			{"$or": []bson.M{
				{"v.connector": searchRegexp},
				{"v.connector_name": searchRegexp},
				{"v.component": searchRegexp},
				{"v.resource": searchRegexp},
			}},
		}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchAndOnlyParents_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	search := "test-search"
	searchRegexp := primitive.Regex{
		Pattern: fmt.Sprintf(".*%s.*", search),
		Options: "i",
	}
	opened := true
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Opened:      &opened,
					Search:      search,
					OnlyParents: true,
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getMetaAlarmRuleLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getChildrenCountLookup()...)
	fields := getComputedFields(now, "")
	fields["is_meta_alarm"] = getIsMetaAlarmField()
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
			"resolved_children":  0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := getOnlyParentsSearchPipeline(bson.M{
		"v.resolved": nil,
		"$or": []bson.M{
			{"v.connector": searchRegexp},
			{"v.connector_name": searchRegexp},
			{"v.component": searchRegexp},
			{"v.resource": searchRegexp},
		},
	}, mongo.AlarmMongoCollection, bson.M{"v.resolved": nil}, false)
	expected = append(expected,
		bson.M{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"v.resolved": nil},
			{"$or": []bson.M{
				{"v.parents": nil},
				{"v.parents": bson.M{"$eq": bson.A{}}},
				{"v.meta": bson.M{"$ne": nil}},
			}},
		}}},
	)
	expected = append(expected,
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchByEntityInfosAndOnlyParents_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	search := "test-search"
	searchRegexp := primitive.Regex{
		Pattern: fmt.Sprintf(".*%s.*", search),
		Options: "i",
	}
	opened := true
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				SearchBy: []string{
					"entity.infos.info1.value",
				},
				BaseFilterRequest: BaseFilterRequest{
					Opened:      &opened,
					Search:      search,
					OnlyParents: true,
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getMetaAlarmRuleLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getChildrenCountLookup()...)
	fields := getComputedFields(now, "")
	fields["is_meta_alarm"] = getIsMetaAlarmField()
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
			"resolved_children":  0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := getOnlyParentsSearchPipeline(bson.M{
		"v.resolved": nil,
		"$or": []bson.M{
			{"e.infos.info1.value": searchRegexp},
		},
	}, mongo.AlarmMongoCollection, bson.M{"v.resolved": nil}, true)
	expected = append(expected,
		bson.M{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"v.resolved": nil},
			{"$or": []bson.M{
				{"v.parents": nil},
				{"v.parents": bson.M{"$eq": bson.A{}}},
				{"v.meta": bson.M{"$ne": nil}},
			}},
		}}},
	)
	expected = append(expected,
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchByEntityInfosAndOnlyParentsForResolvedAlarms_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	search := "test-search"
	searchRegexp := primitive.Regex{
		Pattern: fmt.Sprintf(".*%s.*", search),
		Options: "i",
	}
	opened := true
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				SearchBy: []string{
					"entity.infos.info1.value",
				},
				BaseFilterRequest: BaseFilterRequest{
					Opened:      &opened,
					Search:      search,
					OnlyParents: true,
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getMetaAlarmRuleLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getChildrenCountLookup()...)
	fields := getComputedFields(now, "")
	fields["is_meta_alarm"] = getIsMetaAlarmField()
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
			"resolved_children":  0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := getOnlyParentsSearchPipeline(bson.M{
		"v.resolved": nil,
		"$or": []bson.M{
			{"e.infos.info1.value": searchRegexp},
		},
	}, mongo.ResolvedAlarmMongoCollection, bson.M{"v.resolved": nil}, true)
	expected = append(expected,
		bson.M{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{
			{"v.resolved": nil},
			{"$or": []bson.M{
				{"v.parents": nil},
				{"v.parents": bson.M{"$eq": bson.A{}}},
				{"v.meta": bson.M{"$ne": nil}},
			}},
		}}},
	)
	expected = append(expected,
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpression_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: "connector LIKE \"test name\" AND state = 3",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{
			{"v.meta": nil},
			{"$and": []bson.M{
				{"v.connector": bson.M{"$regex": "test name"}},
				{"v.state.val": bson.M{"$eq": 3}},
			}},
		}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpression_ShouldBuildQueryWithLookupsBeforeMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: "entity.name LIKE \"test name\" AND v.duration > 100",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	delete(fields, "v.duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"v.duration": durationField,
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{
			{"v.meta": nil},
			{"$and": []bson.M{
				{"e.name": bson.M{"$regex": "test name"}},
				{"v.duration": bson.M{"$gt": 100}},
			}},
		}}},
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpressionForResolvedAlarms_ShouldBuildQueryWithLookupsBeforeMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: "entity.name LIKE \"test name\" AND v.duration > 100",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	durationField := fields["v.duration"]
	delete(fields, "v.duration")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"v.duration": durationField,
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{
			{"e.name": bson.M{"$regex": "test name"}},
			{"v.duration": bson.M{"$gt": 100}},
		}}},
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$project": bson.M{"e": 0}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpression_ShouldBuildQueryWithReplaceInfosAlias(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: "infos.test1.value LIKE \"test val\"",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	infosField := fields["infos"]
	delete(fields, "infos")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"infos": infosField,
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{
			{"v.meta": nil},
			{"e.infos.test1.value": bson.M{"$regex": "test val"}},
		}}},
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpressionForResolvedAlarms_ShouldBuildQueryWithReplaceInfosAlias(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Search: "infos.test1.value LIKE \"test val\"",
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	fields := getComputedFields(now, "")
	infosField := fields["infos"]
	delete(fields, "infos")
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": fields,
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$addFields": bson.M{
			"infos": infosField,
		}},
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected, []bson.M{
		{"$match": bson.M{"e.infos.test1.value": bson.M{"$regex": "test val"}}},
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$project": bson.M{"e": 0}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithMultipleWidgetFilters_ShouldBuildQueryWithAllMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter1 := view.WidgetFilter{
		ID: "test-filter-1",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "category",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-category"),
					},
				},
			},
		},
	}
	filter2 := view.WidgetFilter{
		ID: "test-filter-2",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.component",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-component"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.id",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-pbehavior"),
					},
				},
			},
		},
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "resource"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter1, filter2})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter1.ID, filter2.ID},
				},
			},
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"v.steps":            0,
			"pbehavior.comments": 0,
			"e.services":         0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.component": bson.M{"$eq": "test-component"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.id": bson.M{"$eq": "test-pbehavior"}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected,
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"e.category": bson.M{"$eq": "test-category"}},
		}}}}},
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"e.type": bson.M{"$eq": "resource"}},
		}}}}},
		bson.M{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithDependencies_ShouldBuildQueryWithLookupsServicesAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
			WithDependencies: true,
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetImpactsCountPipeline("e")...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetDependsCountPipeline("e")...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.AlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithDependenciesForResolvedAlarms_ShouldBuildQueryWithLookupsServicesAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filter := view.WidgetFilter{
		ID: "test-filter",
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "pbehavior_info.canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, []view.WidgetFilter{filter})
	authorProvider := author.NewProvider(config.NewApiConfigProvider(config.CanopsisConf{}, zerolog.Nop()))
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			FilterRequest: FilterRequest{
				BaseFilterRequest: BaseFilterRequest{
					Filters: []string{filter.ID},
				},
			},
			WithDependencies: true,
		},
	}
	now := datetime.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "t", Value: -1}, {Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getEntityLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetCategoryLookup("e")...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorLookup(authorProvider)...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetImpactsCountPipeline("e")...)
	expectedDataPipeline = append(expectedDataPipeline, dbquery.GetDependsCountPipeline("e")...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(now, ""),
	})
	expectedDataPipeline = append(expectedDataPipeline,
		bson.M{"$addFields": bson.M{
			"e.pbehavior_info": "$v.pbehavior_info",
		}},
		bson.M{"$project": bson.M{
			"bookmarks":          0,
			"e.services":         0,
			"v.steps":            0,
			"pbehavior.comments": 0,
		}},
		bson.M{"$addFields": bson.M{"entity": "$e"}},
		bson.M{"$project": bson.M{"e": 0}},
	)
	expected := []bson.M{
		{"$match": bson.M{"healthcheck": bson.M{"$in": bson.A{nil, false}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"v.pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$match": bson.M{"$and": []bson.M{{"v.meta": nil}}}},
	}
	expected = append(expected, getEntityLookup()...)
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"e.enabled": true}}}},
		{"$project": bson.M{"e": 0}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient, authorProvider, mongo.ResolvedAlarmMongoCollection)
	result, err := b.CreateListAggregationPipeline(ctx, request, now, "")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(author.StripAuthorRandomPrefix(result), author.StripAuthorRandomPrefix(expected)); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func createMockDbClient(ctrl *gomock.Controller) mongo.DbClient {
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockInstructionDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		case mongo.InstructionMongoCollection:
			return mockInstructionDbCollection
		default:
			return nil
		}
	}).AnyTimes()

	return mockDbClient
}

func createMockDbClientWithFilterFetching(ctrl *gomock.Controller, filters []view.WidgetFilter) mongo.DbClient {
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	for _, v := range filters {
		filter := v
		mockSingleResult := mock_mongo.NewMockSingleResultHelper(ctrl)
		mockSingleResult.EXPECT().Decode(gomock.Any()).Do(func(v *view.WidgetFilter) {
			*v = filter
		})
		mockFilterDbCollection.EXPECT().FindOne(gomock.Any(), gomock.Eq(bson.M{"_id": filter.ID})).Return(mockSingleResult)
	}

	mockInstructionDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		case mongo.InstructionMongoCollection:
			return mockInstructionDbCollection
		default:
			return nil
		}
	}).AnyTimes()

	return mockDbClient
}
