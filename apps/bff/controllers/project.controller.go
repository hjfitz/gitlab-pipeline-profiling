package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/gitlab-pipeline-profiling/services"
)

type ProjectController struct {
	ProjectService *services.ProjectService
}

func (c *ProjectController) ApplyRoutes(r *gin.Engine) {
	r.GET("/projects", c.GetProjects)
}

func (c *ProjectController) GetProjects(ctx *gin.Context) {
	projects := c.ProjectService.GetProjects()
	ctx.JSON(http.StatusOK, projects)
}
