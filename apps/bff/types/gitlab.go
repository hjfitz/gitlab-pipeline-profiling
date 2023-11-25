package types

import (
	"time"
)

type PipelineResponse struct {
	Id        int64  `json:"id"`
	Iid       int64  `json:"iid"`
	ProjectId int64  `json:"project_id"`
	Status    string `json:"status"`
	Source    string `json:"source"`
	Ref       string `json:"ref"`
	Sha       string `json:"sha"`
	Name      string `json:"name"`
	WebUrl    string `json:"web_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type JobResponse struct {
	Commit struct {
		AuthorEmail string    `json:"author_email"`
		AuthorName  string    `json:"author_name"`
		CreatedAt   time.Time `json:"created_at"`
		ID          string    `json:"id"`
		Message     string    `json:"message"`
		ShortID     string    `json:"short_id"`
		Title       string    `json:"title"`
	} `json:"commit"`
	Coverage          any       `json:"coverage"`
	AllowFailure      bool      `json:"allow_failure"`
	CreatedAt         time.Time `json:"created_at"`
	StartedAt         time.Time `json:"started_at"`
	FinishedAt        time.Time `json:"finished_at"`
	ErasedAt          any       `json:"erased_at"`
	Duration          float64   `json:"duration"`
	QueuedDuration    float64   `json:"queued_duration"`
	ArtifactsExpireAt time.Time `json:"artifacts_expire_at"`
	TagList           []string  `json:"tag_list"`
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Pipeline          struct {
		ID        int    `json:"id"`
		ProjectID int    `json:"project_id"`
		Ref       string `json:"ref"`
		Sha       string `json:"sha"`
		Status    string `json:"status"`
	} `json:"pipeline"`
	Ref           string `json:"ref"`
	Artifacts     []any  `json:"artifacts"`
	Runner        any    `json:"runner"`
	Stage         string `json:"stage"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason"`
	Tag           bool   `json:"tag"`
	WebURL        string `json:"web_url"`
	Project       struct {
		CiJobTokenScopeEnabled bool `json:"ci_job_token_scope_enabled"`
	} `json:"project"`
	User struct {
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		Username     string    `json:"username"`
		State        string    `json:"state"`
		AvatarURL    string    `json:"avatar_url"`
		WebURL       string    `json:"web_url"`
		CreatedAt    time.Time `json:"created_at"`
		Bio          any       `json:"bio"`
		Location     any       `json:"location"`
		PublicEmail  string    `json:"public_email"`
		Skype        string    `json:"skype"`
		Linkedin     string    `json:"linkedin"`
		Twitter      string    `json:"twitter"`
		WebsiteURL   string    `json:"website_url"`
		Organization string    `json:"organization"`
	} `json:"user"`
}

type ProjectResponse struct {
	ID                   int       `json:"id"`
	Description          string    `json:"description"`
	DefaultBranch        string    `json:"default_branch"`
	TagList              []any     `json:"tag_list"`
	Topics               []any     `json:"topics"`
	Archived             bool      `json:"archived"`
	Visibility           string    `json:"visibility"`
	SSHURLToRepo         string    `json:"ssh_url_to_repo"`
	HTTPURLToRepo        string    `json:"http_url_to_repo"`
	WebURL               string    `json:"web_url"`
	Name                 string    `json:"name"`
	NameWithNamespace    string    `json:"name_with_namespace"`
	Path                 string    `json:"path"`
	PathWithNamespace    string    `json:"path_with_namespace"`
	IssuesEnabled        bool      `json:"issues_enabled"`
	MergeRequestsEnabled bool      `json:"merge_requests_enabled"`
	WikiEnabled          bool      `json:"wiki_enabled"`
	JobsEnabled          bool      `json:"jobs_enabled"`
	SnippetsEnabled      bool      `json:"snippets_enabled"`
	CreatedAt            time.Time `json:"created_at"`
	LastActivityAt       time.Time `json:"last_activity_at"`
	SharedRunnersEnabled bool      `json:"shared_runners_enabled"`
	CreatorID            int       `json:"creator_id"`
	Namespace            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Path string `json:"path"`
		Kind string `json:"kind"`
	} `json:"namespace"`
	AvatarURL            any   `json:"avatar_url"`
	StarCount            int   `json:"star_count"`
	ForksCount           int   `json:"forks_count"`
	OpenIssuesCount      int   `json:"open_issues_count"`
	PublicJobs           bool  `json:"public_jobs"`
	SharedWithGroups     []any `json:"shared_with_groups"`
	RequestAccessEnabled bool  `json:"request_access_enabled"`
}
