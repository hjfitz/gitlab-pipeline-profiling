config:

- GITLAB_TOKEN: token to access gitlab. permissions required tbd
- GROUP_IDS: a comma-delimited list of group IDS to pull repo metadata from

endpoints:
- GET /pipelines. query params: project_id (string) and branch (string)
- GET /projects. no params
