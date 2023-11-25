package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/gitlab-pipeline-profiling/config"
	"github.com/hjfitz/gitlab-pipeline-profiling/services"
)

type Controller interface {
	ApplyRoutes(r *gin.Engine)
}

func BuildControllers() []Controller {
	httpClient := http.Client{}
	config := config.GetConfig()

	gitlabService := &services.GitlabService{
		Config: &config,
		Client: &httpClient,
	}

	pipelineService := &services.PipelineService{
		Gitlab: gitlabService,
	}

	projectService := &services.ProjectService{
		Gitlab: gitlabService,
		Config: &config,
	}

	return []Controller{
		&PipelineController{
			PipelineService: pipelineService,
		},
		&ProjectController{
			ProjectService: projectService,
		},
	}
}
