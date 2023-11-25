package types

type Job struct {
	Name             string `json:"name"`
	Status           string `json:"status"`
	DurationInMillis int64  `json:"duration"`
	StartedAt        string `json:"started_at"`
	FinishedAt       string `json:"finished_at"`
	Stage            string `json:"stage"`
}

type Stage struct {
	Name             string `json:"name"`
	DurationInMillis int64  `json:"duration"`
	JobsRun          int    `json:"number_of_jobs_run"`
}

type Pipeline struct {
	Id               int64   `json:"id"`
	DurationInMillis int64   `json:"duration"`
	StartedAt        string  `json:"started_at"`
	FinishedAt       string  `json:"finished_at"`
	Name             string  `json:"name"`
	Jobs             []Job   `json:"jobs"`
	Stages           []Stage `json:"stages"`
	Failed           bool    `json:"failed"`
}

type Config struct {
	GitlabToken string
	GroupIds    []string
}
