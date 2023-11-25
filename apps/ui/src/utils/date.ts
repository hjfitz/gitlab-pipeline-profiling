
export function formatDate(date: Date | string): string {
  date = new Date(date)
  return date.toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  })
}
