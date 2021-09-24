package pbehavior

import (
	"encoding/json"
	"errors"

	"git.canopsis.net/canopsis/go-engines/lib/api/pagination"
	"git.canopsis.net/canopsis/go-engines/lib/api/pbehaviorexception"
	"git.canopsis.net/canopsis/go-engines/lib/api/pbehaviorreason"
	"git.canopsis.net/canopsis/go-engines/lib/canopsis/pbehavior"
	"git.canopsis.net/canopsis/go-engines/lib/canopsis/types"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type ListRequest struct {
	pagination.FilteredQuery
	SortBy string `form:"sort_by" json:"sort_by" binding:"oneoforempty=name author enabled tstart tstop type.name reason.name created updated rrule type.icon_name"`
}

type EIDsListRequest struct {
	pagination.FilteredQuery
	SortBy string `form:"sort_by" json:"sort_by" binding:"oneoforempty=id"`
}

type EditRequest struct {
	Author     string                             `json:"author" binding:"required,max=255"`
	Enabled    *bool                              `json:"enabled" binding:"required"`
	Filter     interface{}                        `json:"filter" binding:"required"`
	Name       string                             `json:"name" binding:"required,max=255"`
	Reason     string                             `json:"reason" binding:"required"`
	RRule      string                             `json:"rrule"`
	Start      types.CpsTime                      `json:"tstart" binding:"required" swaggertype:"integer"`
	Stop       *types.CpsTime                     `json:"tstop" swaggertype:"integer"`
	Type       string                             `json:"type" binding:"required"`
	Exdates    []pbehaviorexception.ExdateRequest `json:"exdates" binding:"dive"`
	Exceptions []string                           `json:"exceptions"`
}

type CreateRequest struct {
	EditRequest
	ID string `json:"_id" binding:"max=255"`
}

type UpdateRequest struct {
	EditRequest
	ID string `json:"-"`
}

type PatchRequest struct {
	Author     *string                            `json:"author"`
	Enabled    *bool                              `json:"enabled"`
	Filter     interface{}                        `json:"filter"`
	Name       *string                            `json:"name"`
	Reason     *string                            `json:"reason"`
	RRule      *string                            `json:"rrule"`
	Start      *types.CpsTime                     `json:"tstart" swaggertype:"integer"`
	Stop       NullableTime                       `json:"tstop" swaggertype:"integer"`
	Type       *string                            `json:"type"`
	Exdates    []pbehaviorexception.ExdateRequest `json:"exdates" binding:"dive"`
	Exceptions []string                           `json:"exceptions"`
}

type FindByEntityIDRequest struct {
	ID string `form:"id" binding:"required"`
}

type PBehavior struct {
	ID         string                         `bson:"_id" json:"_id"`
	Author     string                         `bson:"author" json:"author"`
	Comments   pbehavior.Comments             `bson:"comments" json:"comments"`
	Enabled    bool                           `bson:"enabled" json:"enabled"`
	Filter     Filter                         `bson:"filter" json:"filter"`
	Name       string                         `bson:"name" json:"name"`
	Reason     *pbehaviorreason.Reason        `bson:"reason" json:"reason"`
	RRule      string                         `bson:"rrule" json:"rrule"`
	Start      *types.CpsTime                 `bson:"tstart" json:"tstart" swaggertype:"integer"`
	Stop       *types.CpsTime                 `bson:"tstop" json:"tstop" swaggertype:"integer"`
	Created    *types.CpsTime                 `bson:"created" json:"created" swaggertype:"integer"`
	Updated    *types.CpsTime                 `bson:"updated" json:"updated" swaggertype:"integer"`
	Type       *pbehavior.Type                `bson:"type" json:"type"`
	Exdates    []pbehaviorexception.Exdate    `bson:"exdates" json:"exdates"`
	Exceptions []pbehaviorexception.Exception `bson:"exceptions" json:"exceptions"`
	// IsActiveStatus represents if pbehavior is in action for current time.
	IsActiveStatus *bool `bson:"-" json:"is_active_status,omitempty"`
}

type Filter struct {
	v interface{}
}

func NewFilter(v interface{}) Filter {
	return Filter{v: v}
}

func (f Filter) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.v)
}

func (f *Filter) UnmarshalBSONValue(_ bsontype.Type, b []byte) error {
	v, _, ok := bsoncore.ReadString(b)
	if !ok {
		return errors.New("invalid value, expected string")
	}

	err := json.Unmarshal([]byte(v), &f.v)
	return err
}

type NullableTime struct {
	*types.CpsTime
	isSet bool
}

func (t *NullableTime) UnmarshalJSON(data []byte) error {
	t.isSet = true
	if string(data) == "null" {
		return nil
	}
	temp := &types.CpsTime{}
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}
	t.CpsTime = temp
	return nil
}

type AggregationResult struct {
	Data       []PBehavior `bson:"data" json:"data"`
	TotalCount int64       `bson:"total_count" json:"total_count"`
}

func (r *AggregationResult) GetData() interface{} {
	return r.Data
}

func (r *AggregationResult) GetTotal() int64 {
	return r.TotalCount
}

type EID struct {
	ID string `bson:"id" json:"id"`
}

type AggregationEIDsResult struct {
	Data       []EID `bson:"data" json:"data"`
	TotalCount int64 `bson:"total_count" json:"total_count"`
}

func (r *AggregationEIDsResult) GetData() interface{} {
	return r.Data
}

func (r *AggregationEIDsResult) GetTotal() int64 {
	return r.TotalCount
}
