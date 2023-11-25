"use client"
import { ArcElement, Legend, Tooltip } from "chart.js"
import { Chart as ChartJS } from "chart.js/auto"
import type { FC } from "react"
import { Line } from "react-chartjs-2"

import type { ChartDefaultPropertyTypes } from "../../types/chart.component"
import { formatDate } from "~/utils/date"

ChartJS.register(ArcElement, Tooltip, Legend)

const chartOptions = {
  plugins: {
    legend: {
      display: true,
    },
  },
  scales: {
    y: {
      display: true,
      title: {
        display: true,
        value: "Duration (minutes)",
      },
    },
  },
}

export type IChartCardProps = ChartDefaultPropertyTypes

export const BuildDurationChart: FC<IChartCardProps> = ({
  pipelineData,
  chartHeight,
}: IChartCardProps): JSX.Element => {
  const labels = pipelineData.map((p) => formatDate(p.weekly_pipelines.week))

  const data = pipelineData.map((p) =>
    p.weekly_pipelines.pipelines
      .map(({ duration }) => duration)
      .reduce((a, b) => a + b, 0),
  )

  const chartData = {
    datasets: [{
	data,
	label: "Duration",
      backgroundColor: "#9BD0F5",
      borderColor: "#9BD0F5",
	}],
    labels,
  }

  return (
    <section>
      <Line data={chartData} height={chartHeight} options={chartOptions} />
    </section>
  )
}
