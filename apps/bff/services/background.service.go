package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/hjfitz/gitlab-pipeline-profiling/types"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type BackgroundService struct {
	Cache           *redis.Client
	PipelineService *PipelineService
	ProjectService  *ProjectService
}

func (s *BackgroundService) Schedule(freq time.Duration) {
	ticker := time.NewTicker(freq)
	log.Info().Msgf("Scheduling background service to run every %s", freq)

	go func() {
		for range ticker.C {
			log.Info().Msg("Running background service")
			projects := s.ProjectService.GetProjects()
			for _, project := range projects {
				go func(project ProjectResponseDTO) {
					pipelines := s.PipelineService.GetPipelinesAndJobs(project.Id, project.DefaultBranch)
					pipelinesStr, _ := json.Marshal(pipelines)
					projectKey := fmt.Sprintf("pipelines:%s:%s", project.Id, project.DefaultBranch)
					ctx := context.Background()
					s.Cache.Set(ctx, projectKey, pipelinesStr, (6 * time.Hour))
				}(project)
			}
			pipelineKey := "pipelines"
			ctx := context.Background()
			pipelinesStr, _ := json.Marshal(projects)
			s.Cache.Set(ctx, pipelineKey, pipelinesStr, (6 * time.Hour))
		}
	}()
}
