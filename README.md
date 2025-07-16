# Allergy App

Allergy App is a microservices-based web application for tracking allergies, sharing experiences, and participating in a community forum. The project demonstrates modern backend and frontend development practices, including JWT/session authentication, RESTful APIs, and a user-friendly interface.

---

## Architecture

```
allergy-app/
  frontend/         # HTML, CSS, JavaScript (UI)
  user-service/     # Go (authentication, JWT/session, user management, Redis)
  forum-service/    # Go (forum posts, tags, search, Postgres)
```

- **Frontend:** Handles user login, registration, post creation, tag filtering, and forum browsing.
- **User Service:** Manages user registration, login, JWT and session management using Redis.
- **Forum Service:** Manages forum post creation, tag filtering, and search using PostgreSQL.

---

## Getting Started

### Prerequisites

- Go (for backend services)
- Docker and Docker Compose (recommended for setup)
- Redis (for user-service)
- PostgreSQL (for forum-service)

### Quick Start with Docker

1. Clone the repository:
   ```sh
   git clone https://github.com/YOUR_USERNAME/allergy-app.git
   cd allergy-app
   ```

2. Copy and configure environment variables:
   ```sh
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Start all services:
   ```sh
   docker-compose up --build
   ```

### Manual Start

1. Start Redis and Postgres locally.
2. Set environment variables for each service (see `.env.example`).
3. Run each Go service:
   ```sh
   cd user-service && go run main.go
   cd forum-service && go run main.go
   ```
4. Open `frontend/login.html` in your browser.

---

## Authentication Flow

- On login or registration, the backend sets a session cookie and returns a JWT.
- The frontend stores the JWT in `localStorage` and uses it for API requests.
- The session cookie is used for browser session management.

---

## Example API Usage

**Login**
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Create Post**
```http
POST /api/posts
Authorization: Bearer <JWT>
Content-Type: application/json

{
  "title": "My Allergy Story",
  "content": "Details...",
  "tags": "pollen,gluten"
}
```

## API Documentation

- See [API.md](API.md) for endpoint details.
- Swagger/OpenAPI documentation available at `/docs` if implemented.

---

## Security

- Secrets and credentials are managed via environment variables.
- Passwords are hashed using bcrypt.
- JWTs are signed with a secret key.
- Refresh tokens and sessions are stored in Redis.

---

## Features

- User registration and login
- JWT and session authentication
- Forum post creation with tags
- Tag filtering and search
- Responsive frontend
- Docker support (add Dockerfile/docker-compose.yml)
- API documentation (add API.md or Swagger)
- Unit and integration tests

---

## License

MIT

---

For questions or suggestions, johnaddokufuor@gmail.com
