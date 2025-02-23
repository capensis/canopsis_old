package user

import (
	"errors"
	"net/http"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/auth"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/bulk"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/common"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/pagination"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type API interface {
	common.BulkCrudAPI
	Patch(c *gin.Context)
	BulkPatch(c *gin.Context)
}

type api struct {
	store  Store
	logger zerolog.Logger

	metricMetaUpdater metrics.MetaUpdater
}

func NewApi(
	store Store,
	logger zerolog.Logger,
	metricMetaUpdater metrics.MetaUpdater,
) API {
	return &api{
		store:  store,
		logger: logger,

		metricMetaUpdater: metricMetaUpdater,
	}
}

// List
// @Success 200 {object} common.PaginatedListResponse{data=[]User}
func (a *api) List(c *gin.Context) {
	var query ListRequest
	query.Query = pagination.GetDefaultQuery()

	if err := c.ShouldBind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, query))
		return
	}

	users, err := a.store.Find(c, query, c.MustGet(auth.UserKey).(string))
	if err != nil {
		panic(err)
	}

	res, err := common.NewPaginatedResponse(query.Query, users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get
// @Success 200 {object} User
func (a *api) Get(c *gin.Context) {
	user, err := a.store.GetOneBy(c, c.Param("id"))
	if err != nil {
		panic(err)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create
// @Param body body CreateRequest true "body"
// @Success 201 {object} User
func (a *api) Create(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))
		return
	}

	user, err := a.store.Insert(c, request)
	if err != nil {
		panic(err)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	a.metricMetaUpdater.UpdateById(c, user.ID)

	c.JSON(http.StatusCreated, user)
}

// Update
// @Param body body UpdateRequest true "body"
// @Success 200 {object} User
func (a *api) Update(c *gin.Context) {
	request := UpdateRequest{
		ID: c.Param("id"),
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))
		return
	}

	user, err := a.store.Update(c, request, c.MustGet(auth.UserKey).(string))
	if err != nil {
		valErr := common.ValidationError{}
		if errors.As(err, &valErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, valErr.ValidationErrorResponse())

			return
		}

		panic(err)
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	a.metricMetaUpdater.UpdateById(c, user.ID)

	c.JSON(http.StatusOK, user)
}

// Patch
// @Param body body PatchRequest true "body"
// @Success 200 {object} User
func (a *api) Patch(c *gin.Context) {
	request := PatchRequest{
		ID: c.Param("id"),
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))
		return
	}

	user, err := a.store.Patch(c, request, c.MustGet(auth.UserKey).(string))
	if err != nil {
		valErr := common.ValidationError{}
		if errors.As(err, &valErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, valErr.ValidationErrorResponse())

			return
		}

		panic(err)
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	a.metricMetaUpdater.UpdateById(c, user.ID)

	c.JSON(http.StatusOK, user)
}

func (a *api) Delete(c *gin.Context) {
	id := c.Param("id")
	ok, err := a.store.Delete(c, id, c.MustGet(auth.UserKey).(string))
	if err != nil {
		valErr := common.ValidationError{}
		if errors.As(err, &valErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(err))

			return
		}

		panic(err)
	}

	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	a.metricMetaUpdater.DeleteById(c, id)

	c.Status(http.StatusNoContent)
}

// BulkCreate
// @Param body body []CreateRequest true "body"
func (a *api) BulkCreate(c *gin.Context) {
	userIDs := make([]string, 0)
	bulk.Handler(c, func(request CreateRequest) (string, error) {
		user, err := a.store.Insert(c, request)
		if err != nil {
			return "", err
		}

		userIDs = append(userIDs, user.ID)
		return user.ID, nil
	}, a.logger)
	a.metricMetaUpdater.UpdateById(c, userIDs...)
}

// BulkUpdate
// @Param body body []BulkUpdateRequestItem true "body"
func (a *api) BulkUpdate(c *gin.Context) {
	userID := c.MustGet(auth.UserKey).(string)
	userIDs := make([]string, 0)
	bulk.Handler(c, func(request BulkUpdateRequestItem) (string, error) {
		user, err := a.store.Update(c, UpdateRequest(request), userID)
		if err != nil || user == nil {
			return "", err
		}

		userIDs = append(userIDs, user.ID)

		return user.ID, nil
	}, a.logger)
	a.metricMetaUpdater.UpdateById(c, userIDs...)
}

// BulkDelete
// @Param body body []BulkDeleteRequestItem true "body"
func (a *api) BulkDelete(c *gin.Context) {
	userID := c.MustGet(auth.UserKey).(string)
	userIDs := make([]string, 0)
	bulk.Handler(c, func(request BulkDeleteRequestItem) (string, error) {
		ok, err := a.store.Delete(c, request.ID, userID)
		if err != nil || !ok {
			return "", err
		}

		userIDs = append(userIDs, request.ID)
		return request.ID, nil
	}, a.logger)

	a.metricMetaUpdater.DeleteById(c, userIDs...)
}

// BulkPatch
// @Param body body []BulkPatchRequestItem true "body"
func (a *api) BulkPatch(c *gin.Context) {
	userID := c.MustGet(auth.UserKey).(string)
	userIDs := make([]string, 0)
	bulk.Handler(c, func(request BulkPatchRequestItem) (string, error) {
		user, err := a.store.Patch(c, PatchRequest(request), userID)
		if err != nil || user == nil {
			return "", err
		}

		userIDs = append(userIDs, user.ID)

		return user.ID, nil
	}, a.logger)
	a.metricMetaUpdater.UpdateById(c, userIDs...)
}
