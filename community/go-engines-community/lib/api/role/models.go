package role

import (
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/author"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/pagination"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/security"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/datetime"
)

const (
	TypeUI  = "ui"
	TypeAPI = "api"

	PermissionAPIPrefix = "api_"
)

type ListRequest struct {
	pagination.FilteredQuery
	SortBy     string `form:"sort_by" binding:"oneoforempty=name"`
	Permission string `form:"permission"`
}

type EditRequest struct {
	Name        string              `json:"name" binding:"required,max=255"`
	Type        string              `json:"type" binding:"required,oneof=ui api"`
	Description string              `json:"description" binding:"max=255"`
	DefaultView string              `json:"defaultview"`
	Permissions map[string][]string `json:"permissions"`

	AuthConfig security.AuthMethodConf `json:"auth_config"`
	Author     string                  `json:"author" swaggerignore:"true"`
}

type BulkUpdatePermissionsRequestItem struct {
	ID          string              `json:"_id" binding:"required"`
	Permissions map[string][]string `json:"permissions"`
}

type Response struct {
	ID          string       `bson:"_id" json:"_id"`
	Name        string       `bson:"name" json:"name"`
	Type        string       `bson:"type" json:"type"`
	Description string       `bson:"description" json:"description"`
	DefaultView *View        `bson:"defaultview" json:"defaultview"`
	Permissions []Permission `bson:"permissions" json:"permissions"`
	Editable    *bool        `bson:"editable,omitempty" json:"editable,omitempty"`
	Deletable   *bool        `bson:"deletable,omitempty" json:"deletable,omitempty"`

	AuthConfig security.AuthMethodConf `bson:"auth_config" json:"auth_config"`

	Author  *author.Author    `bson:"author,omitempty" json:"author,omitempty"`
	Created *datetime.CpsTime `bson:"created,omitempty" json:"created,omitempty" swaggertype:"integer"`
	Updated *datetime.CpsTime `bson:"updated,omitempty" json:"updated,omitempty" swaggertype:"integer"`
}

type Permission struct {
	ID      string   `bson:"_id" json:"_id"`
	Name    string   `bson:"name" json:"name"`
	View    string   `bson:"view" json:"view,omitempty"`
	Type    string   `bson:"type" json:"type"`
	Actions []string `bson:"actions" json:"actions"`

	Bitmask               int8                     `bson:"bitmask" json:"-"`
	ApiPermissions        map[string]int8          `bson:"api_permissions" json:"-"`
	ApiPermissionsBitmask map[int8]map[string]int8 `bson:"api_permissions_bitmask" json:"-"`
}

type View struct {
	ID    string `bson:"_id" json:"_id"`
	Title string `bson:"title" json:"title"`
}

type AggregationResult struct {
	Data       []Response `bson:"data" json:"data"`
	TotalCount int64      `bson:"total_count" json:"total_count"`
}

func (r *AggregationResult) GetData() interface{} {
	return r.Data
}

func (r *AggregationResult) GetTotal() int64 {
	return r.TotalCount
}

type Template struct {
	ID          string       `bson:"_id" json:"_id"`
	Name        string       `bson:"name" json:"name"`
	Type        string       `bson:"type" json:"type"`
	Description string       `bson:"description" json:"description"`
	Permissions []Permission `bson:"permissions" json:"permissions"`
}

type TemplateResponse struct {
	Data []Template `json:"data"`
}
