package services

import (
	"fmt"
	"sync"

	. "github.com/hjfitz/gitlab-pipeline-profiling/types"
)

type ProjectService struct {
	Gitlab *GitlabService
	Config *Config
}

func (s *ProjectService) GetProjects() []ProjectResponseDTO {
	projects := []ProjectResponseDTO{}
	var wg sync.WaitGroup
	for _, groupId := range s.Config.GroupIds {
		wg.Add(1)
		go func(groupId string) {
			rawGroupProjects := s.Gitlab.GetProjectsForGroup(groupId)
			groupProjects := parseProjects(rawGroupProjects)
			projects = append(projects, groupProjects...)
			wg.Done()
		}(groupId)
	}

	wg.Wait()

	return projects

}

func parseProjects(rawGroupProjects []ProjectResponse) []ProjectResponseDTO {
	projects := []ProjectResponseDTO{}
	for _, project := range rawGroupProjects {
		projects = append(projects, ProjectResponseDTO{
			Name:          project.Name,
			Id:            fmt.Sprint(project.ID),
			DefaultBranch: project.DefaultBranch,
			WebUrl:        project.WebURL,
		})
	}

	return projects
}
