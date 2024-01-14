# Gino's ReactJS (frontend) with Go (backend) Project Template

This uses the following stacks:

- Project is bootstrapped using **Vite** with the following options: *React* and *TypeScript*.
- **TailwindCSS** is the CSS framework.
- **Prettier** plugin for TailwindCSS is enabled.

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
