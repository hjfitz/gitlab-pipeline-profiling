package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/gitlab-pipeline-profiling/config"
	"github.com/hjfitz/gitlab-pipeline-profiling/services"
	"github.com/redis/go-redis/v9"
)

type Controller interface {
	ApplyRoutes(r *gin.Engine)
}

func BuildControllers() []Controller {
	httpClient := http.Client{}
	config := config.GetConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	gitlabService := &services.GitlabService{
		Config: &config,
		Client: &httpClient,
	}

	pipelineService := &services.PipelineService{
		Gitlab: gitlabService,
		Cache:  rdb,
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
