Allergy App
Allergy App is a microservices-based web application for tracking allergies, sharing experiences, and participating in a community forum. It was built not just to showcase backend skills, but to think through how infrastructure — in code — shapes access, interaction, and trust.

At its core, the app is about structure. Who gets to log in? Who gets to post? How is identity verified? These questions are answered not by the UI, but by the backend — through token checks, session management, and the layout of the data models themselves. This project is a technical response to those questions.

Architecture
pgsql
Copy
Edit
allergy-app/
  frontend/         # HTML, CSS, JavaScript (UI)
  user-service/     # Go (authentication, JWT/session, user management, Redis)
  forum-service/    # Go (forum posts, tags, search, Postgres)
Frontend: Handles user login, registration, post creation, tag filtering, and forum browsing. It’s minimalist — built without frameworks — so interaction depends directly on the API’s structure.

User Service: Controls registration, login, and authentication. Uses Redis to track sessions and rotate refresh tokens. Passwords are hashed. JWTs are signed.

Forum Service: Lets users create posts, comment, and filter by tags. All actions are tied to verified user identities, checked via the token passed in each request.

Getting Started
Prerequisites
Go (for backend services)

Docker and Docker Compose (recommended for setup)

Redis (for user-service)

PostgreSQL (for forum-service)

Manual Start
Start Redis and Postgres locally.

Set environment variables for each service (see .env.example).

Run each Go service:

sh
Copy
Edit
cd user-service && go run main.go
cd forum-service && go run main.go
Open frontend/login.html in your browser.

Authentication Flow
On login or registration, the backend returns a JWT and sets a session cookie.

The frontend stores the JWT in localStorage and attaches it to protected API requests.

Redis tracks sessions and allows token revocation on logout.

The backend never trusts data from the frontend — it validates every action.

Example API Usage
Login

http
Copy
Edit
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
Create Post

http
Copy
Edit
POST /api/posts
Authorization: Bearer <JWT>
Content-Type: application/json

{
  "title": "My Allergy Story",
  "content": "Details...",
  "tags": "pollen,gluten"
}
API Documentation
See API.md for endpoint details.

Swagger/OpenAPI docs available at /docs if implemented.

Security
Passwords are hashed using bcrypt.

JWTs are signed with a secret key and verified per request.

Refresh tokens are tracked in Redis and rotated automatically.

No user ID is accepted from the frontend — it’s derived server-side from the token.

Features
User registration and login

JWT + session-based authentication

Forum post creation with tag filtering

Responsive static frontend

Docker-ready deployment

RESTful API design

Clean separation between services

License
MIT



For questions or suggestions, johnaddokufuor@gmail.com
