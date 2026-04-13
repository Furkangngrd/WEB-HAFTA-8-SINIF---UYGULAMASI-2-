# GoLearn Education API

GoLearn Remote Education Platform API, developed with Go, Gin, SQLite, GORM, and JWT authentication. Clean architecture with RBAC (Role-Based Access Control) support, Quiz systems, Progress tracking, and WebSockets functionality.

## Features
- **User Authentication**: JWT-based login/register (Student/Teacher roles)
- **Role-Based Access Control**: Middleware to authorize users by roles.
- **Courses & Lessons**: CRUD operations for interactive courses.
- **Quizzes**: Real-time quiz creation and evaluation logic.
- **WebSockets**: Real-time chat & notifications module.
- **Dockerized**: Easy-to-deploy multi-container docker build.
- **Rate-Limiting**: IP-based rate limiting per user connection.
- **Swagger Documentation**: Accessible over (`/swagger/index.html`).

## Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose (Optional)

### Run with Docker (Recommended)
```bash
docker-compose up --build
```
The server will be running at `http://localhost:8080`.

### Run Locally (Without Docker)
1. Install dependencies:
```bash
go mod tidy
```
2. Build and run:
```bash
go run main.go
```

## API Endpoints (Quick Overview)

### Auth
- `POST /api/v1/auth/register` - Create account
- `POST /api/v1/auth/login` - Authenticate & get JWT

### Courses
- `GET /api/v1/courses` - List courses
- `POST /api/v1/courses` - Create course (Teacher only)
- `POST /api/v1/courses/:id/enroll` - Enroll in a course (Student only)

### Lessons
- `GET /api/v1/courses/:id/lessons` - List lessons
- `POST /api/v1/courses/:id/lessons` - Add a lesson (Teacher only)

### Quizzes
- `POST /api/v1/courses/:id/quizzes` - Create a quiz
- `POST /api/v1/courses/:id/quizzes/:quizId/submit` - Submit quiz answers

### WebSockets
- `GET /api/v1/ws/:roomId` - Connect to chat room (Needs Auth header)

## Project Structure
- `main.go`: API Entry point
- `config/`: Configuration parsing & env variable bindings
- `database/`: Database connectivity
- `models/`: GORM objects & Data Transfer Objects (DTO)
- `handlers/`: Route controllers & logic mapping
- `middleware/`: HTTP security chains
- `websocket/`: Pub/Sub socket handling
