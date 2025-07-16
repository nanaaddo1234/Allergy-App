# API Documentation

This document describes the main API endpoints for the Allergy App microservices.

---

## User Service

### Register

**POST** `/auth/register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
- `201 Created`  
  ```json
  {
    "message": "User registered successfully"
  }
  ```
- `400 Bad Request`  
  ```json
  {
    "error": "Email already exists"
  }
  ```

---

### Login

**POST** `/auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
- `200 OK`  
  ```json
  {
    "token": "<JWT>",
    "message": "Login successful"
  }
  ```
- `401 Unauthorized`  
  ```json
  {
    "error": "Invalid credentials"
  }
  ```

---

### Logout

**POST** `/auth/logout`

**Headers:**  
`Authorization: Bearer <JWT>`

**Response:**
- `200 OK`  
  ```json
  {
    "message": "Logged out successfully"
  }
  ```

---

## Forum Service

### Create Post

**POST** `/api/posts/`

**Headers:**  
`Authorization: Bearer <JWT>`  
`Content-Type: application/json`

**Request Body:**
```json
{
  "title": "My Allergy Story",
  "content": "Details about my experience...",
  "tags": "pollen,gluten"
}
```

**Response:**
- `201 Created`  
  ```json
  {
    "message": "Post created successfully",
    "post_id": 123
  }
  ```
- `400 Bad Request`  
  ```json
  {
    "error": "Missing required fields"
  }
  ```
- `401 Unauthorized`  
  ```json
  {
    "error": "Invalid or missing token"
  }
  ```

---

### Get Posts

**GET** `/api/posts/`

**Query Parameters (optional):**
- `tag`: Filter posts by tag (e.g., `/api/posts/?tag=pollen`)
- `search`: Search posts by keyword

**Response:**
- `200 OK`  
  ```json
  [
    {
      "id": 123,
      "title": "My Allergy Story",
      "content": "Details...",
      "tags": ["pollen", "gluten"],
      "author": "user@example.com",
      "created_at": "2025-07-15T12:34:56Z"
    },
    ...
  ]
  ```

---

### Get Single Post

**GET** `/api/posts/{id}`

**Response:**
- `200 OK`  
  ```json
  {
    "id": 123,
    "title": "My Allergy Story",
    "content": "Details...",
    "tags": ["pollen", "gluten"],
    "author": "user@example.com",
    "created_at": "2025-07-15T12:34:56Z"
  }
  ```
- `404 Not Found`  
  ```json
  {
    "error": "Post not found"
  }
  ```

---

## Authentication

- Most endpoints (except register/login) require the `Authorization: Bearer <JWT>` header.
- JWT is obtained from the login endpoint and should be stored securely on the client.

---

## Error Format

All error responses follow this format:
```json
{
  "error": "Description of the error"
}
```

---

_This document should be updated as new endpoints or features are added._
