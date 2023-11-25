import Head from 'next/head'
import { GitlabService } from '../services/gitlab.service'
import { PipelineResponseDTO, ProjectResponseDTO } from '~/types/api'
import { useEffect, useRef, useState } from 'react'
import { ChartContainer } from '~/components/chart-container.component'
import { BuildDurationChart } from '~/components/charts/build-duration.component'
import { FailuresChart } from '~/components/charts/failure-vs-success.component'
import { FailurePercentageChart } from '~/components/charts/failure-perc.component'
import { DurationChart } from '~/components/charts/duration.component'
import { getConfig } from '~/utils/config'


interface IHomepageProps {
  projects: ProjectResponseDTO[]
}

function buildGetServerSideProps(gitlabService: GitlabService) {
  return async function() {
    const projects = await gitlabService.getProjects()
	return {
	  props: {
	    projects,
      },
    }
  }
}

const getServerSideProps = buildGetServerSideProps(new GitlabService(getConfig().internalApiBase))

function toFriendlyName(name: string) {
		return name.split('-').map((word) => word[0].toUpperCase() + word.slice(1)).join(' ')
}

const chartHeight = 200

const Home = (props: IHomepageProps) => {
  const [selectedProject, setSelectedProject] = useState<ProjectResponseDTO>(props.projects[0])
  const [projects, setProjects] = useState<ProjectResponseDTO[]>(props.projects)
  const [pipelineData, setPipelineData] = useState<PipelineResponseDTO[]>([])
  const gitlabService = useRef(new GitlabService())
  useEffect(() => {
	gitlabService.current.getProjects().then(resp => setProjects(resp))
  }, [])

  useEffect(() => {
		if (!selectedProject) {
				return
		}

		gitlabService.current
				.getPipelinesForProject(selectedProject.id, selectedProject.default_branch)
				.then((pipelineData) => {
						setPipelineData(pipelineData)
				})
  }, [selectedProject])
  return (
    <>
      <Head>
        <title>Gitlab Pipeline Profiler</title>
        <meta name="description" content="Gitlab Pipeline Profiler" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="h-full grid grid-cols-12">
	    <aside className="col-span-3 bg-indigo-800">
				<header>
						<h1 className="text-2xl text-white text-center py-4 mb-1 bg-indigo-900">
								Gitlab Pipeline Profiler
						</h1>
				</header>
				<section className="text-sm">
						{projects.map((project) => (
								<button
										key={project.id}
										onClick={() => setSelectedProject(project)}
										className={`w-full text-left text-white hover:bg-indigo-600 px-4 py-2 ${selectedProject.id === project.id ? 'selectedProject' : ''}`}
								>
										{toFriendlyName(project.name)}
								</button>
						))}
				</section>
		</aside>
		<section className="bg-gradient-to-b from-slate-100 to-white col-span-9">
				<header className="bg-white ">
						<h1 className="text-2xl text-center py-4">
								{toFriendlyName(selectedProject.name)}
						</h1>
				</header>
				<div className="grid grid-cols-12 m-8 gap-4">

						<div className="col-span-6">
								<ChartContainer title="Pipeline runs" chartHeight={200}>
										<FailuresChart
												pipelineData={pipelineData}
												chartHeight={chartHeight}
										/>
								</ChartContainer>
						</div>


						<div className="col-span-6">
								<ChartContainer title="Pipeline Error Rate" chartHeight={200}>
										<FailurePercentageChart
												pipelineData={pipelineData}
												chartHeight={chartHeight}
										/>
								</ChartContainer>
						</div>


						<div className="col-span-6">
								<ChartContainer title="Pipeline Length (seconds)" chartHeight={200}>
										<BuildDurationChart
												pipelineData={pipelineData}
												chartHeight={chartHeight}
										/>
								</ChartContainer>
						</div>


						<div className="col-span-6">
								<ChartContainer title="Time spent per stage (%)" chartHeight={200}>
										<DurationChart
												pipelineData={pipelineData}
												chartHeight={chartHeight}
										/>
								</ChartContainer>
						</div>

				</div>
		</section>
      </main>
    </>
  )
}

export { getServerSideProps }

export default Home
