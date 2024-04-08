## Project Name

Welcome to the project repository! This is a monorepo that contains two main components: a React-based frontend and a Go Gin-based backend.

## Directory Structure

The repository is organized as follows:

```
.
├── LICENSE
├── README.md
├── server
│   ├── cmd
│   │   ├── config.go
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   └── pkg
│       ├── api.go
│       ├── models
│       │   └── responses.go
│       ├── routes
│       │   └── health.go
│       └── utils
│           └── env.go
└── ui
    ├── README.md
    ├── index.html
    ├── package-lock.json
    ├── package.json
    ├── public
    │   └── vite.svg
    ├── src
    │   ├── App.tsx
    │   ├── index.css
    │   ├── main.tsx
    │   └── vite-env.d.ts
    ├── tsconfig.json
    ├── tsconfig.node.json
    └── vite.config.ts
```

## Prerequisites

### Go

To set up the Go backend, you'll need to install [Go](https://golang.org/doc/install) version 1.22 or later.

### React and TypeScript

For the frontend, you'll need to update your Node.js installation to use NVM and install the latest LTS version of Node.js:
1. Install NVM (Node Version Manager) using the curl command:
   ```bash
   # This will download and install the latest version of NVM.
   curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
   ```
2. Close and reopen your terminal for the changes to take effect.
3. Install the latest LTS (Long-Term Support) version of Node.js using NVM:
   ```bash
   # This will install the latest LTS version of Node.js.
   nvm install --lts
   ```
4. Set the installed LTS version as the default:
   ```bash
   # This will make the latest LTS version of Node.js the default version used in your environment.
   nvm alias default node
   ```
5. Verify the installed version:
   ```bash
   # This should output the version of the installed LTS Node.js, for example, `v20.11.0`.
   node -v
   ```

## Installation

### Go Backend

1. Navigate to the `server` directory:
   ```bash
   cd server
   ```
2. Install the Go dependencies:
   ```bash
   go mod download
   ```
3. Run the Go server:
   ```bash
   go run ./cmd
   ```

> If you wish to run an executable/binary of the go server, run the following commands
> ```shell
> go build -o ./bin/server ./cmd # compile go binary
> ./bin/server
> ```

> To check if server is running execute
> ```shell
> curl http:/localhost:3000/checks/ping
> # {"ping":"pong"}
> ```

### React Frontend

1. Navigate to the `ui` directory:
   ```bash
   cd ui
   ```
2. Install the npm dependencies:
   ```bash
   npm install
   ```
3. Start the React development server:
   ```bash
   npm run dev
   ```

## Useful Commands

### Go

- `go build`: Compile the Go code.
- `go test`: Run the Go tests.
- `go fmt`: Format the Go code.
- `go vet`: Check the Go code for potential issues.

### React

- `npm run dev`: Start the React development server.
- `npm test`: Run the React tests.
- `npm run build`: Build the React application for production.
- `npm run lint`: Check the React code for linting issues.
- `npm run format`: Format the React code.
