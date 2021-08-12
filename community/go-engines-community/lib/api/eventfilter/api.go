package eventfilter

import (
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/logger"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/eventfilter"
	"net/http"

	mongodriver "go.mongodb.org/mongo-driver/mongo"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/pagination"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/common"
	"github.com/gin-gonic/gin"
)

type api struct {
	store        Store
	actionLogger logger.ActionLogger
}

// Create eventfilter
// @Summary Create eventfilter
// @Description Create eventfilter
// @Tags eventfilters
// @ID eventfilters-create
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BasicAuth
// @Param body body eventfilter.Rule true "body"
// @Success 201 {object} eventfilter.Rule
// @Failure 400 {object} common.ErrorResponse
// @Router /eventfilter/rules [post]
func (a api) Create(c *gin.Context) {
	request := eventfilter.Rule{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))
		return
	}

	err := a.store.Insert(c.Request.Context(), &request)
	if err != nil {
		panic(err)
	}

	err = a.actionLogger.Action(c, logger.LogEntry{
		Action:    logger.ActionCreate,
		ValueType: logger.ValueTypeEventFilter,
		ValueID:   request.ID,
	})
	if err != nil {
		a.actionLogger.Err(err, "failed to log action")
	}

	c.JSON(http.StatusCreated, request)
}

// Find all eventfilter
// @Summary Find all eventfilter
// @Description Get paginated list of eventfilter
// @Tags eventfilters
// @ID eventfilters-find-all
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BasicAuth
// @Param page query integer true "current page"
// @Param limit query integer true "items per page"
// @Param search query string false "search query"
// @Success 200 {object} common.PaginatedListResponse{data=[]eventfilter.Rule}
// @Failure 400 {object} common.ErrorResponse
// @Router /eventfilter/rules [get]
func (a api) List(c *gin.Context) {
	var query FilteredQuery
	query.Query = pagination.GetDefaultQuery()

	if err := c.ShouldBind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, query))
		return
	}

	aggregationResult, err := a.store.Find(c.Request.Context(), query)
	if err != nil {
		panic(err)
	}

	res, err := common.NewPaginatedResponse(query.Query, aggregationResult)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get eventfilter by id
// @Summary Get eventfilter by id
// @Description Get eventfilter by id
// @Tags eventfilters
// @ID eventfilters-get-by-id
// @Produce json
// @Security ApiKeyAuth
// @Security BasicAuth
// @Param id path string true "eventfilter id"
// @Success 200 {object} eventfilter.Rule
// @Failure 404 {object} common.ErrorResponse
// @Router /eventfilter/rules/{id} [get]
func (a api) Get(c *gin.Context) {
	evf, err := a.store.GetById(c.Request.Context(), c.Param("id"))

	if err == mongodriver.ErrNoDocuments || evf == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, evf)
}

// Update eventfilter by id
// @Summary Update eventfilter by id
// @Description Update eventfilter by id
// @Tags eventfilters
// @ID eventfilters-update-by-id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BasicAuth
// @Param id path string true "eventfilter id"
// @Param body body eventfilter.Rule true "body"
// @Success 200 {object} eventfilter.Rule
// @Failure 400 {object} common.ValidationErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Router /eventfilter/rules/{id} [put]
func (a api) Update(c *gin.Context) {
	var request eventfilter.Rule
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))

		return
	}

	request.ID = c.Param("id")
	ok, _ := a.store.Update(c.Request.Context(), &request)

	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	err := a.actionLogger.Action(c, logger.LogEntry{
		Action:    logger.ActionUpdate,
		ValueType: logger.ValueTypeEventFilter,
		ValueID:   request.ID,
	})
	if err != nil {
		a.actionLogger.Err(err, "failed to log action")
	}

	c.JSON(http.StatusOK, request)
}

// Delete eventfilter by id
// @Summary Delete eventfilter by id
// @Description Delete eventfilter by id
// @Tags eventfilters
// @ID eventfilters-delete-by-id
// @Security ApiKeyAuth
// @Security BasicAuth
// @Param id path string true "eventfilter id"
// @Success 204
// @Failure 404 {object} common.ErrorResponse
// @Router /eventfilter/rules/{id} [delete]
func (a api) Delete(c *gin.Context) {
	ok, err := a.store.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		panic(err)
	}

	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
		return
	}

	err = a.actionLogger.Action(c, logger.LogEntry{
		Action:    logger.ActionDelete,
		ValueType: logger.ValueTypeEventFilter,
		ValueID:   c.Param("id"),
	})
	if err != nil {
		a.actionLogger.Err(err, "failed to log action")
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewApi(
	store Store,
	actionLogger logger.ActionLogger,
) common.CrudAPI {
	return &api{
		store:        store,
		actionLogger: actionLogger,
	}
}
