"use client"
import { ArcElement, Legend, Tooltip } from "chart.js"
import { Chart as ChartJS } from "chart.js/auto"
import { useState, type FC, useEffect, useRef } from "react"
import { Bar } from "react-chartjs-2"

import type { ChartDefaultPropertyTypes } from "../../types/chart.component"
import { formatDate } from "~/utils/date"

ChartJS.register(ArcElement, Tooltip, Legend)

const defaultChartOptions = {
    plugins: {
    legend: {
      display: false
    },
  },
  scales: {
    x: {
      stacked: true,
    },
    y: {
      stacked: true,
    },
  },
}

export type IDurationChartProps = ChartDefaultPropertyTypes

export function unique<T>(arr: T[]): T[] {
  return Array.from(new Set(arr))
}

export const DurationChart: FC<IDurationChartProps> = ({
  pipelineData,
  chartHeight,
}): JSX.Element => {
  const chartRef = useRef(null)
  const [showLegend, setShowLegend] = useState(false)

  const [chartOptions, setChartOptions] = useState(defaultChartOptions)

  useEffect(() => {
    if (!chartRef.current) return
    // @ts-expect-error not today batman
		chartRef.current?.update()
  }, [pipelineData])


  useEffect(() => {
		setChartOptions({
				...chartOptions, 
				plugins: {
						legend: { 
								display: showLegend 
						} 
				} 
		})
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [showLegend])
  const labels = pipelineData.map((p) => formatDate(p.weekly_pipelines.week))

  const stages = pipelineData.map((p) =>
    p.weekly_pipelines.pipelines.flatMap((pipeline) => pipeline.stages),
  )

  const stageNames = unique(
    stages.flatMap((stage) => stage.map((s) => s.name)),
  )

  const datasets = stageNames.map((stageName) => ({
    data: stages.flatMap((stage) =>
      stage
        .filter((s) => s.name === stageName)
        .map((s) => s.duration)
        .map((durationMillis) => durationMillis / 1000 / 60),
    ),
    label: stageName,
  }))

  const chartData = {
    datasets,
    labels,
  }

  return (
    <section>
	  <p 
	    onClick={() => setShowLegend(!showLegend)}
		className="text-sm text-gray-400 cursor-pointer"
	  >
	    Toggle legend</p>
        <Bar
				data={chartData}
				height={chartHeight} 
				options={chartOptions}
				ref={chartRef}
		/>
    </section>
  )
}
