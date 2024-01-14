# Gino's ReactJS (frontend) with Go (backend) Project Template

This uses the following stacks:

- Project is bootstrapped using **Vite** with the following options: *React* and *TypeScript*.
- **TailwindCSS** is the CSS framework.
- **Prettier** plugin for TailwindCSS is enabled.
- **Air** runs the server in a development reload loop (install first with `go install github.com/cosmtrek/air@latest`)
- **Gin** is the HTTP web framework.
- **Viper** for reading configuration from environment variables.
- **Zap** for server-side logging.

Additional features:

VSCode will run the following tasks in the background:

- `npm run dev`
- `npm run dev-tw`

The following subaths in `client/src/<path>` are mapped to `@<path>` in TypeScript

- `assets`
- `components`
- `context`
- `utils`

## Customisation

1. `client/package*.json`: change `{"name": "vite-project"}` to your project's name.
2. `server/go.mod`: change module path to your project's repo URL.
3. `server/Makefile`: change `PROJ` to your project's name and `ORGPATH` to your project's base repo URL.
4. `.vscode/tasks.json`: change the environment variables (e.g. REACT_GO_TEMPLATE_LISTEN_ADDRESS) based on your app's name.
5. `client/vite.config.ts`: replace the `envPrefix` based on your app's name.
6. `client/src/utils/server.ts`: change environment variable prefixes (ie REACT_GO_TEMPLATE_*).

## Configuration

All configuration is read from environment variables.

| Key                                | Scope  | Description                                                                                     |
| ---------------------------------- | ------ | ----------------------------------------------------------------------------------------------- |
| `REACT_GO_TEMPLATE_SERVER_URL`     | Client | The URL of the backend server.                                                                  |
| `REACT_GO_TEMPLATE_LISTEN_ADDRESS` | Server | Binds the backend server to this address. Should be the same as `REACT_GO_TEMPLATE_SERVER_URL`. |
| `REACT_GO_TEMPLATE_ENABLE_CORS`    | Server | Enables CORS. Useful if the client and server are running on different processes.               |
