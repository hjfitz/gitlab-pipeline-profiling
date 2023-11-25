import type { FC } from "react"
import { Bar } from "react-chartjs-2"

import type { ChartDefaultPropertyTypes } from "../../types/chart.component"
import { formatDate } from "~/utils/date"

const chartOptions = {
  responsive: true,
  scales: {
    x: {
      stacked: true,
    },
    y: {
      stacked: true,
    },
  },
}

export type IFailureChartProps = ChartDefaultPropertyTypes

export const FailuresChart: FC<IFailureChartProps> = ({
  pipelineData,
  chartHeight,
}) => {
  const labels = pipelineData.map((p) => formatDate(p.weekly_pipelines.week))

  const failed = pipelineData.map(
    ({ weekly_pipelines: weekly }) =>
      weekly.pipelines.filter((p) => p.failed).length,
  )

  const passed = pipelineData.map(
    ({ weekly_pipelines: weekly }) =>
      weekly.pipelines.filter((p) => !p.failed).length,
  )

  const datasets = [
    {
      data: passed,
      fill: true,
      label: "Passed",
    },
    {
      data: failed,
      fill: true,
      label: "Failed",
    },
  ]

  const chartData = {
    datasets,
    labels,
  }

  return (
    <section>
      <Bar data={chartData} height={chartHeight} options={chartOptions} />
    </section>
  )
}
