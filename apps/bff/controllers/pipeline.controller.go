package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/gitlab-pipeline-profiling/services"
)

type PipelineController struct {
	PipelineService *services.PipelineService
}

func (c *PipelineController) ApplyRoutes(r *gin.Engine) {
	r.GET("/pipelines", c.GetPipelines)
}

func (c *PipelineController) GetPipelines(ctx *gin.Context) {

	branch, branchExists := ctx.GetQuery("branch")
	projectId, projectExists := ctx.GetQuery("project_id")

	if !branchExists || !projectExists {
		ctx.String(http.StatusBadRequest, "Missing branch or project_id parameter")
		return
	}

	pipelines := c.PipelineService.GetPipelinesAndJobs(projectId, branch)
	ctx.JSON(http.StatusOK, pipelines)
}
