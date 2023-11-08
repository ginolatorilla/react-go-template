const apiURL = (path: string) => `${import.meta.env.REACT_GO_TEMPLATE_SERVER_URL ?? ""}${path}`

export async function hello() {
  const response = await fetch(apiURL("/"));
  return await response.text();
}
