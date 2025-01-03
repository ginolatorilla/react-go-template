# Gino's ReactJS+Gin Project Template

## UI

- Project is bootstrapped using **Vite** with the following options: *React* and *TypeScript*.
- **TailwindCSS** is the CSS framework (including Tailwind Forms).
- **Prettier** plugin for TailwindCSS is enabled.
- UI is **embedded** into the server.

The following subpaths in `ui/src/<path>` are mapped to `@<path>` in TypeScript

- `assets`
- `components`
- `context`
- `utils`

## Server

- **Gin** is the HTTP web framework with logging and CORS middlewares.
- **Viper** for configuration.
- **Zap** for server-side logging.
- **Air** runs the server in a development reload loop (install first with `go install github.com/cosmtrek/air@latest`)

The server defines these routes:

- `/` for the UI
- `/api/v1/hello` for the API

## VSCode integration

VSCode will run the following tasks in the background (see `.vscode/tasks.json`):

- `npm run dev`, which runs an auto-reloading development server for the UI at <http://localhost:5173>.
- `air`, which runs an auto-reloading development server for the _server_ at <http://localhost:8080>.

The UI connects to the server because it loads the server's address from `ui/.env.development`.
Air runs the server with CORS enabled, which allows the requests from the UI (see `.air.toml`).

## Customisation

1. `ui/package*.json`: change `{"name": "vite-project"}` to your project's name.
2. `go.mod`: change module path to your project's repo URL.
3. `Makefile`: change `APP` to your project's name and `GITHUB_OWNER` to your GitHub username.
4. `ui/vite.config.ts`: replace the `envPrefix` based on your app's name.
5. `ui/src/utils/server.ts`: change environment variable prefixes (ie REACT_GO_TEMPLATE_*).

## Configuration

All configuration is read from environment variables.

| Key                                | Scope  | Description                                                                                     |
| ---------------------------------- | ------ | ----------------------------------------------------------------------------------------------- |
| `REACT_GO_TEMPLATE_SERVER_URL`     | UI     | The URL of the backend server.                                                                  |
| `REACT_GO_TEMPLATE_LISTEN_ADDRESS` | Server | Binds the backend server to this address. Should be the same as `REACT_GO_TEMPLATE_SERVER_URL`. |
| `REACT_GO_TEMPLATE_ENABLE_CORS`    | Server | Enables CORS. Useful if the client and server are running on different processes.               |

## Building

Run `make` to build the UI and the server. It will create `bin/<app>`, which contains both the server and UI.

Run `make install` so you can run the server from anywhere.

See `make help` for more options.
