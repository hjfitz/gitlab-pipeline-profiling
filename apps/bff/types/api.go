package types

type ProjectResponseDTO struct {
	Name          string `json:"name"`
	Id            string `json:"id"`
	DefaultBranch string `json:"default_branch"`
	WebUrl        string `json:"web_url"`
}

type PipelineResponseDTO struct {
	WeeklyPipelnes struct {
		Week      string     `json:"week"`
		Pipelines []Pipeline `json:"pipelines"`
	} `json:"weekly_pipelines"`
}
