
export const Json = ({json}: {json: any}) => {
  return (
    <div>
      <pre>{JSON.stringify(json, null, 2)}</pre>
    </div>
  )
}
