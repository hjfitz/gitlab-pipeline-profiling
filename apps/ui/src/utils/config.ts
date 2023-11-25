interface Config {
  apiBase: string
  internalApiBase: string
}

function getConfig(): Config {
  return {
    apiBase: process.env.NEXT_PUBLIC_API_BASE!,
    internalApiBase: process.env.API_BASE!,
  }
}

export { getConfig }
