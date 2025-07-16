# Allergy App

**Allergy App** is a microservices-based web application for allergy tracking, personal health journaling, and forum-based discussion. This project demonstrates backend system design with secure authentication, PostgreSQL data persistence, and microservice separation. It emphasizes backend architecture and API structure over frontend framework usage.

## Project Structure

```
allergy-app/
├── frontend/              # Static UI (HTML, CSS, JS)
├── user-service/          # Go microservice for auth (JWT + Redis)
└── forum-service/         # Go microservice for posts/comments (PostgreSQL)
```

### Microservices

- **Frontend**: Login, registration, post creation, and tag filtering. Uses raw JavaScript with fetch-based API interaction.
- **User Service**: Handles user registration, login, logout, JWT issuance, and Redis-based session tracking.
- **Forum Service**: Handles post creation, comment submission, tag filtering, and JWT-authenticated user validation.

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL (forum-service)
- Redis (user-service)
- Docker (optional)

### Run Locally

```bash
# Start Redis and PostgreSQL manually

cd user-service
go run main.go

cd ../forum-service
go run main.go

# Then open frontend/login.html in your browser
```

## Authentication Flow

- After login or registration, the user-service returns a JWT and a refresh token.
- The JWT is stored in `localStorage` and sent via `Authorization` headers.
- Redis manages session state and allows logout and token revocation.
- The forum-service extracts `user_id` from the verified JWT for secure post creation.

## API Usage

Refer to [API.md](API.md) for full endpoint documentation.

Example:

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

## Security Overview

- Passwords hashed with bcrypt.
- JWTs signed with HMAC secret.
- Redis stores refresh tokens for session control.
- No client-supplied user IDs are trusted.

## Features

- User registration and login
- Stateless authentication with Redis-backed session
- PostgreSQL storage for forum posts/comments
- Secure API interaction
- Separation of concerns via microservices
- HTML/CSS/JS frontend without frameworks

## License

MIT


For questions or suggestions, johnaddokufuor@gmail.com
