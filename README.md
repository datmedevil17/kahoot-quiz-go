# Kahoot Quiz Clone

A real-time multiplayer quiz application inspired by Kahoot. Built with a Go (Gin) backend and a Svelte (Vite) frontend.

## 🏗 Architecture

```mermaid
graph TD
    Client[Client (Svelte)] -->|HTTP/REST| API[API Gateway (Gin Router)]
    Client -->|WebSocket| WS[WebSocket Hub]
    
    subgraph Server
        API --> Auth[Auth Middleware]
        API --> UserSvc[User Service]
        API --> QuizSvc[Quiz Service]
        API --> GameSvc[Game Service]
        
        WS --> Room[Game Room]
        Room --> GameLogic[Game Logic]
        
        GameLogic --> DB[(PostgreSQL)]
        UserSvc --> DB
        QuizSvc --> DB
    end
```

## 🚀 Features

- **Real-time Gameplay**: WebSocket-based synchronized game state for all players.
- **Quiz Management**: Create, edit, and list quizzes.
- **User Authentication**: Secure signup and login with JWT.
- **Game Sessions**: Host games, multiple players join via PIN.
- **Live Leaderboard**: Real-time score updates.

## 🛠 Tech Stack

- **Backend**: Go (Golang), Gin Framework, Gorm, Gorilla WebSocket.
- **Frontend**: Svelte, Vite, TypeScript.
- **Database**: PostgreSQL.
- **Infrastructure**: Docker, Docker Compose.

## 📋 Prerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.25+](https://go.dev/) (for local dev)
- [Node.js 20+](https://nodejs.org/) (for local dev)

## ⚡ Quick Start (Docker)

The easiest way to run the application is using Docker Compose.

1.  **Clone the repository**:
    ```bash
    git clone <repository-url>
    cd kahoot-quiz
    ```

2.  **Start the services**:
    ```bash
    make docker-up
    # OR
    docker compose up -d
    ```

    This will start:
    - **Frontend**: http://localhost:5173
    - **Backend**: http://localhost:8080
    - **Database**: Port 5432

3.  **Stop the services**:
    ```bash
    make docker-down
    ```

## 🏃 Local Development

### Server

1.  Navigate to `server`:
    ```bash
    cd server
    ```
2.  Install dependencies:
    ```bash
    go mod download
    ```
3.  Set up environment:
    ```bash
    cp .env.example .env
    # Update .env with your local DB credentials if needed
    ```
4.  Run the server:
    ```bash
    go run ./cmd/api
    ```

### Client

1.  Navigate to `client`:
    ```bash
    cd client
    ```
2.  Install dependencies:
    ```bash
    npm install
    ```
3.  Run the dev server:
    ```bash
    npm run dev
    ```

## 📚 API Endpoints

| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| `POST` | `/auth/signup` | Register a new user | No |
| `POST` | `/auth/login` | Login user | No |
| `GET` | `/api/v1/users/me` | Get current user profile | Yes |
| `POST` | `/api/v1/quizzes` | Create a new quiz | Yes |
| `GET` | `/api/v1/quizzes` | List my quizzes | Yes |
| `POST` | `/api/v1/games` | Create a game session | Yes |
| `WS` | `/api/v1/ws` | WebSocket connection | Yes |

## 📁 Project Structure

```
.
├── client/                 # Frontend (Svelte)
├── server/                 # Backend (Go)
│   ├── cmd/api/            # Entry point
│   ├── internal/
│   │   ├── api/            # Router setup
│   │   ├── handlers/       # HTTP Handlers
│   │   ├── models/         # Database Models
│   │   ├── services/       # Business Logic
│   │   └── ws/             # WebSocket Logic
│   └── .env                # Env variables
├── docker-compose.yaml     # Docker orchestration
└── Makefile                # Development commands
```
