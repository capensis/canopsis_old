package contextgraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/common"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const (
	dirPerm     = 0o770
	filePerm    = 0o640
	filePattern = "import_%s.json"
)

type api struct {
	reporter    StatusReporter
	dir         string
	filePattern string
	logger      zerolog.Logger
}

func NewApi(
	conf config.CanopsisConf,
	reporter StatusReporter,
	logger zerolog.Logger,
) API {
	a := &api{
		dir:         filepath.Join(conf.File.Dir, canopsis.SubDirImport),
		filePattern: filePattern,
		reporter:    reporter,
		logger:      logger,
	}

	return a
}

// ImportAll
// @Param body body []importcontextgraph.EntityConfiguration true "body"
// @Success 200 {object} contextgraph.ImportResponse
func (a *api) ImportAll(c *gin.Context) {
	query := ImportQuery{}
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, query))
		return
	}

	job := ImportJob{
		Creation: time.Now(),
		Status:   StatusPending,
		Source:   query.Source,
	}

	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err))
		return
	}

	jobID, err := a.createImportJob(c, job, raw)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ImportResponse{ID: jobID})
}

// ImportPartial
// @Param body body []importcontextgraph.EntityConfiguration true "body"
// @Success 200 {object} contextgraph.ImportResponse
func (a *api) ImportPartial(c *gin.Context) {
	query := ImportQuery{}
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, query))
		return
	}

	job := ImportJob{
		Creation:  time.Now(),
		Status:    StatusPending,
		Source:    query.Source,
		IsPartial: true,
	}

	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err))
		return
	}

	jobID, err := a.createImportJob(c, job, raw)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ImportResponse{ID: jobID})
}

func (a *api) createImportJob(ctx context.Context, job ImportJob, raw []byte) (string, error) {
	err := a.reporter.ReportCreate(ctx, &job)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(a.dir, os.ModeDir|dirPerm)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(filepath.Join(a.dir, fmt.Sprintf(a.filePattern, job.ID)), raw, filePerm)
	if err != nil {
		return "", err
	}

	return job.ID, nil
}

// Status
// @Success 200 {object} contextgraph.ImportJob
func (a *api) Status(c *gin.Context) {
	status, err := a.reporter.GetStatus(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundResponse)
			return
		}

		panic(err)
	}

	c.JSON(http.StatusOK, status)
}
