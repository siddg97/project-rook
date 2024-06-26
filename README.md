## Project Name

Welcome to the project repository! This is a monorepo that contains two main components: a React-based frontend and a Go Gin-based backend.

## Directory Structure

The repository is organized as follows:

```
├── LICENSE
├── README.md
├── server -----------> Directory containing server source code (go)
└── ui ---------------> Directory containing frontend code (react/ts)
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
3. Create a `.env` file with the following
   ```bash
   LOG_LEVEL=debug
   PORT=3000
   ENV=local
   GEMINI_KEY=<Google Gemini API Key>
   GOOGLE_APPLICATION_CREDENTIALS=/absolute/path/to/repo/project-rook/server/gcp-sa.json # Update this path accordingly
   ```
4. Create a `gcp-sa.json` file with a [private key generated from gcloud console](https://console.cloud.google.com/iam-admin/serviceaccounts/details/102126832070335213872/keys?hl=en&project=western-voyage-419302) (click `ADD KEY` and download the json file)
5. Run the Go server:
   ```bash
   go run .
   ```

> If you wish to run an executable/binary of the go server, run the following commands
> ```shell
> go build -o ./bin/server ./cmd # compile go binary
> ./bin/server
> ```

> If you wish to run a docker container for the server, do the setup for local as described above. Then install docker, and run the following commands
> ```shell
> cd server
> docker build -t server .
> docker run -p 3000:3000 -e GEMINI_KEY=<Google Gemini API Key> -e GOOGLE_APPLICATION_CREDENTIALS=/gcp-sa.json -v ./gcp-sa.json:/gcp-sa.json server
> ```

- To check if server is running execute
    ```shell
    curl http:/localhost:3000/checks/ping
    # {"ping":"pong"}
    ```

### React Frontend

1. Navigate to the `ui` directory:
   ```bash
   cd ui
   ```
2. Install the npm dependencies:
   ```bash
   npm install
   ```
3. Create a `.env` file with the following
   ```bash
   VITE_FIREBASE_API_KEY=AIzaSyBlYjcyuV4Z4sWzmVXC_Nu9MLCDQB1utX4
   VITE_FIREBASE_AUTH_DOMAIN=western-voyage-419302.firebaseapp.com
   VITE_FIREBASE_PROJECT_ID=western-voyage-419302
   VITE_FIREBASE_STORAGE_BUCKET=western-voyage-419302.appspot.com
   VITE_FIREBASE_MESSAGING_SENDER_ID=178510623950
   VITE_FIREBASE_APP_ID=1:178510623950:web:b141ea76815e4e290c8c9f
   VITE_FIREBASE_MEASUREMENT_ID=G-4CGW4ELJ13
   ```
4. Start the React development server:
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
