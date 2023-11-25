import type { PipelineResponseDTO, ProjectResponseDTO } from "../types/api"
import { getConfig } from "../utils/config"

export class GitlabService {
  constructor(private readonly apiBase = getConfig().apiBase) {}

  public async getPipelinesForProject(
    projectId: string,
    branch: string,
  ): Promise<PipelineResponseDTO[]> {
    const endpoint = new URL("/pipelines", this.apiBase)

    const parameters = new URLSearchParams({
      branch,
      // eslint-disable-next-line camelcase
      project_id: projectId,
    })

    endpoint.search = parameters.toString()

    const resp = await fetch(endpoint.toString())

	try {
			const data = (await resp.json()) as PipelineResponseDTO[]

			return this.sortPipelines(data)
	} catch (err) {
		console.error(err)
		return []
	}
  }

  public async getProjects(): Promise<ProjectResponseDTO[]> {
    const endpoint = new URL("/projects", this.apiBase)

    const resp = await fetch(endpoint.toString())

    const data = (await resp.json()) as ProjectResponseDTO[]

    return data
  }

  private sortPipelines(
    pipelines: PipelineResponseDTO[],
  ): PipelineResponseDTO[] {
    return pipelines.sort((a, b) => {
      const aDate = new Date(a.weekly_pipelines.week)
      const bDate = new Date(b.weekly_pipelines.week)
      return aDate.getTime() - bDate.getTime()
    })
  }
}
