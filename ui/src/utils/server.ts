const apiURL = (path: string) =>
  `${import.meta.env.REACT_GO_TEMPLATE_SERVER_URL ?? ""}${path}`;

export async function hello() {
  const response = await fetch(apiURL("/api/v1/hello"));
  return await response.text();
}
