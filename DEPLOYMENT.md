# Deployment and Development Guide

## Running with Docker (Recommended for Full Stack)

To run the entire application stack (Postgres, Go Server, Frontend Client) using Docker Compose:

1.  **Build and Start**:
    ```bash
    make docker-up
    # OR
    docker compose up -d
    ```

2.  **View Logs**:
    ```bash
    make docker-logs
    # OR
    docker compose logs -f
    ```

3.  **Stop**:
    ```bash
    make docker-down
    # OR
    docker compose down
    ```

## Local Development (Hybrid)

To run the Go server locally while using Docker for the database:

### 1. Start Postgres Database
Start only the database service from the docker-compose setup:

```bash
docker compose up -d db
```

Wait until the database is ready. You can check status with:
```bash
docker compose ps
```

### 2. Configure Environment
Ensure your `.env` file in `server/.env` (or environment variables) points to the local Docker database.
Default `docker-compose.yaml` maps Postgres port `5432` to host `5432`.

**Example .env for local server:**
```env
PORT=8080
DATABASE_URL=host=localhost user=postgres password=postgres dbname=kahoot_quiz port=5432 sslmode=disable
JWT_SECRET=your_local_secret
CLIENT_URL=http://localhost:5173
```

### 3. Run the Go Server
Navigate to the server directory and run:

```bash
cd server
go run ./cmd/api
```
Or use the Makefile from root:
```bash
make run-server
```

### 4. Run the Client (Optional)
If you need the frontend as well:
```bash
cd client
npm run dev
```
Or:
```bash
make run-client
```

## Health Checks

- **API Health**: `GET http://localhost:8080/health`
- **WebSocket Health**: `GET ws://localhost:8080/health/ws`
