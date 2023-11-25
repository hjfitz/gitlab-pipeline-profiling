package services

import (
	"sync"
	"time"

	. "github.com/hjfitz/gitlab-pipeline-profiling/types"
)

type PipelineService struct {
	Gitlab *GitlabService
}

// TODO: we should probably store the response from gitlab
// if there is a request, fetch them all from the store
// defer an update using `since` in the API
func (s *PipelineService) GetPipelinesAndJobs(projectId, branch string) []PipelineResponseDTO {
	pipelines := s.Gitlab.GetPipelines(projectId, branch)

	pipelines = s.HydratePipelineJobs(projectId, pipelines)

	sortPipelineByDate(pipelines)

	groupedPipelines := groupPipelinesByWeek(pipelines)

	return groupedPipelines
}

func (s *PipelineService) HydratePipelineJobs(projectId string, pipelines []Pipeline) []Pipeline {
	hydratedPipelines := []Pipeline{}
	var wg sync.WaitGroup
	for _, pipeline := range pipelines {
		wg.Add(1)

		go func(pipeline Pipeline) {
			hydratedPipeline := s.HydratePipelineWithJobs(pipeline, projectId)
			hydratedPipelines = append(hydratedPipelines, hydratedPipeline)
			wg.Done()
		}(pipeline)

	}
	wg.Wait()

	return hydratedPipelines
}

func (s *PipelineService) HydratePipelineWithJobs(pipeline Pipeline, projectId string) Pipeline {
	jobs := s.Gitlab.GetPipelineJobs(pipeline.Id, projectId)
	stages := getPipelineStages(jobs)
	pipelineDuration := getPipelineDuration(jobs)

	finishedAt := pipeline.FinishedAt
	if finishedAt == "" {
		started, _ := time.Parse(time.RFC3339, pipeline.StartedAt)
		finishedAt = started.Add(time.Duration(pipelineDuration) * time.Millisecond).Format(time.RFC3339)
	}

	return Pipeline{
		Id:               pipeline.Id,
		DurationInMillis: pipelineDuration,
		StartedAt:        pipeline.StartedAt,
		FinishedAt:       finishedAt,
		Name:             pipeline.Name,
		Jobs:             jobs,
		Stages:           stages,
		Failed:           pipeline.Failed,
	}
}

func sortPipelineByDate(pipelines []Pipeline) {
	for i := 0; i < len(pipelines); i++ {
		for j := i + 1; j < len(pipelines); j++ {
			a, _ := time.Parse(time.RFC3339, pipelines[i].StartedAt)
			b, _ := time.Parse(time.RFC3339, pipelines[j].StartedAt)
			if a.Unix() > b.Unix() {
				pipelines[i], pipelines[j] = pipelines[j], pipelines[i]
			}
		}
	}
}

func getPipelineStages(jobs []Job) []Stage {
	stages := map[string]Stage{}

	for _, job := range jobs {
		stage := stages[job.Stage]

		stage.Name = job.Stage
		stage.DurationInMillis += job.DurationInMillis
		stage.JobsRun++
		stages[job.Stage] = stage
	}

	stagesArray := []Stage{}

	for _, stage := range stages {
		stagesArray = append(stagesArray, stage)
	}

	return stagesArray
}

func getPipelineDuration(job []Job) int64 {
	var duration int64 = 0
	for _, job := range job {
		duration += job.DurationInMillis
	}
	return duration
}

func groupPipelinesByWeek(sortedPipelines []Pipeline) []PipelineResponseDTO {
	if len(sortedPipelines) == 0 {
		return []PipelineResponseDTO{}
	}
	// assume that pipelines are sorted by date
	// get all possible weeks from start to end
	// for each week, get all pipelines that started in that week
	start, _ := time.Parse(time.RFC3339, sortedPipelines[0].StartedAt)
	// set start hours and minutes to 0
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())

	// create start of week
	start = start.AddDate(0, 0, -int(start.Weekday()))

	end, _ := time.Parse(time.RFC3339, sortedPipelines[len(sortedPipelines)-1].StartedAt)
	end = end.AddDate(0, 0, -int(end.Weekday()))

	weekMap := map[time.Time][]Pipeline{}
	for i := start; i.Before(end); i = i.AddDate(0, 0, 7) {
		weekMap[i] = []Pipeline{}
	}

	for _, pipeline := range sortedPipelines {
		started, _ := time.Parse(time.RFC3339, pipeline.StartedAt)
		// round started to beginning of the week
		pipelineWeek := time.Date(started.Year(), started.Month(), started.Day(), 0, 0, 0, 0, started.Location())
		pipelineWeek = pipelineWeek.AddDate(0, 0, -int(pipelineWeek.Weekday()))
		weekMap[pipelineWeek] = append(weekMap[pipelineWeek], pipeline)
	}

	weekArray := []PipelineResponseDTO{}

	for week, pipelines := range weekMap {
		resp := PipelineResponseDTO{}
		resp.WeeklyPipelnes.Week = week.Format(time.RFC3339)
		resp.WeeklyPipelnes.Pipelines = pipelines
		weekArray = append(weekArray, resp)
	}

	return weekArray
}
