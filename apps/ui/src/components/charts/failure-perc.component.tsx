import type { FC } from "react"
import { Line } from "react-chartjs-2"

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

export type IChartCardProps = ChartDefaultPropertyTypes

export const FailurePercentageChart: FC<IChartCardProps> = ({
  pipelineData,
  chartHeight,
}) => {
  const labels = pipelineData.map((p) => formatDate(p.weekly_pipelines.week))

  const errorPerc = pipelineData.map((p) => {
    const failed = p.weekly_pipelines.pipelines.filter(
      (pline) => pline.failed,
    ).length

    const total = p.weekly_pipelines.pipelines.length
    return (failed / total) * 100
  })

  const datasets = [
    {
      backgroundColor: "#ffa5b4",
      borderColor: "#ffa5b4",
      data: errorPerc,
      label: "Failed",
    },
  ]

  const chartData = {
    datasets,
    labels,
  }

  return (
    <section>
      <Line data={chartData} height={chartHeight} options={chartOptions} />
    </section>
  )
}
