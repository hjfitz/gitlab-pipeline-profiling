import { FC } from "react"

interface ChartContainerProps {
  title: string
  chartHeight: number
  children: React.ReactNode
}

const ChartContainer: FC<ChartContainerProps> = ({chartHeight, title, children}) => {
  return (
    <section className="bg-white rounded-lg shadow-lg h-80 p-4">
	  <header className="text-slate-500">
			  {title}
	  </header>
	  <div>
			  {children}
	  </div>
    </section>
  )
}

export { ChartContainer }
