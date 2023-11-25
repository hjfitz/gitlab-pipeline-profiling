interface ProjectResponseDTO {
  name: string
  id: string
  web_url: string
  default_branch: string
}

interface PipelineResponseDTO {
  weekly_pipelines: {
    week: string
    pipelines: PipelineDTO[]
  }
}

interface PipelineDTO {
  id: number
  duration: number
  started_at: string
  finished_at: string
  name: string
  jobs: JobDTO[]
  stages: StageDTO[]
  failed: boolean
}

interface StageDTO {
  name: string
  duration: number
  number_of_jobs: number
}

interface JobDTO {
  name: string
  status: string
  duration: number
  started_at: string
  finished_at: string
  stage: string
}

export {
  JobDTO,
  PipelineDTO,
  PipelineResponseDTO,
  ProjectResponseDTO,
  StageDTO,
}
