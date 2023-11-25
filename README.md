# GitLab Pipeline Profiler

## Overview

The GitLab Pipeline Profiler is a powerful tool designed to analyze the runtime and performance of GitLab pipelines. It provides insights that can help optimize your CI/CD processes, making them more efficient and reliable.

## Features

- Detailed analysis of GitLab pipeline performance
- Easy to install and use
- Docker-based deployment

## Installation

### Prerequisites

- Docker and Docker Compose installed on your machine
- GitLab access token

### Steps

1. Clone the repository to your local machine using the following command:

```bash
$ git clone git@github.com:hjfitz/gitlab-pipeline-profiling.git
```

2. Create a file named env.production.local in the project root directory and add the following fields:

```ini
GROUP_IDS=<Your Group IDs>
GITLAB_TOKEN=<Your GitLab Token>
```

3. Run the application using Docker Compose:

```bash
$ docker-compose up
```

After the setup completes, the GitLab Pipeline Profiler should be up and running.

## Usage

Detailed usage instructions will be added soon.

## Contributing

Fork it, pull request it.

## Contact

For any queries or further information, please reach out to Harry Fitzgerald at harry@hjf.io.
