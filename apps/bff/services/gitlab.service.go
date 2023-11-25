package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/hjfitz/gitlab-pipeline-profiling/types"
)

// TODO: do with config
const NUM_PIPELINES = 100

type GitlabService struct {
	Config *Config
	Client *http.Client
}

func (s *GitlabService) GetPipelines(projectId, branch string) []Pipeline {
	pipelines := []Pipeline{}

	pipelineUrl := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/pipelines?ref=%s&per_page=%d", projectId, branch, NUM_PIPELINES)

	req, _ := http.NewRequest("GET", pipelineUrl, nil)

	req.Header.Add("PRIVATE-TOKEN", s.Config.GitlabToken)

	resp, _ := s.Client.Do(req)
	defer resp.Body.Close()

	pipelineResponse := []PipelineResponse{}

	json.NewDecoder(resp.Body).Decode(&pipelineResponse)

	for _, pipeline := range pipelineResponse {
		pipelineData := Pipeline{
			Id:        pipeline.Id,
			Failed:    pipeline.Status == "failed",
			Name:      pipeline.Name,
			StartedAt: pipeline.CreatedAt,
		}
		pipelines = append(pipelines, pipelineData)
	}

	return pipelines
}

func (s *GitlabService) GetPipelineJobs(pipelineId int64, projectId string) []Job {

	jobUrl := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/pipelines/%d/jobs?per_page=100", projectId, pipelineId)
	jobUrl = fmt.Sprintf("%s&scope[]=success&scope[]=failed", jobUrl)

	req, _ := http.NewRequest("GET", jobUrl, nil)

	req.Header.Add("PRIVATE-TOKEN", s.Config.GitlabToken)

	resp, err := s.Client.Do(req)

	if err != nil {
		// todo: log
		return []Job{}
	}

	defer resp.Body.Close()

	jobResponse := []JobResponse{}

	json.NewDecoder(resp.Body).Decode(&jobResponse)

	jobs := []Job{}

	for _, job := range jobResponse {
		jobData := Job{
			Name:             job.Name,
			Status:           job.Status,
			Stage:            job.Stage,
			DurationInMillis: int64(job.Duration * 1000),
			StartedAt:        job.StartedAt.Format(time.RFC3339),
			FinishedAt:       job.FinishedAt.Format(time.RFC3339),
		}
		jobs = append(jobs, jobData)
	}

	return jobs
}

func (s *GitlabService) GetProjectsForGroup(groupId string) []ProjectResponse {

	url := fmt.Sprintf("https://gitlab.com/api/v4/groups/%s/projects?per_page=100", groupId)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("PRIVATE-TOKEN", s.Config.GitlabToken)

	resp, err := s.Client.Do(req)

	if err != nil {
		fmt.Println(err)
		return []ProjectResponse{}
	}

	defer resp.Body.Close()

	var apiResponse []ProjectResponse
	json.NewDecoder(resp.Body).Decode(&apiResponse)

	return apiResponse

}
