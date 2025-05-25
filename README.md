# Go Events REST API

A REST API built with Go that provides event management functionality with user authentication. Users can create, update, delete events and register/cancel registration for events.

## Features

- User authentication with JWT
- Event CRUD operations
- Event registration management
- SQLite database storage
- RESTful API endpoints

## Technologies Used

- Go 1.24.2
- Gin Web Framework (github.com/gin-gonic/gin)
- SQLite3 (github.com/mattn/go-sqlite3)
- JWT Authentication (github.com/golang-jwt/jwt/v5)
- Bcrypt for password hashing (golang.org/x/crypto/bcrypt)

## Installation

1. Clone the repository
2. Install Go 1.24.2 or later
3. Install dependencies: `go mod tidy`
4. Start server: `go run main.go`
5. Server runs on http://localhost:8080

## API Documentation

### Authentication Endpoints

POST /signup

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

POST /login

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

Response includes JWT token for authentication

### Event Endpoints

GET /events

- No payload required
- Returns all events

GET /events/:id

- No payload required
- Returns single event

POST /events (Authenticated)

```json
{
  "name": "Event Name",
  "description": "Event Description",
  "location": "Event Location",
  "datetime": "2025-05-25T19:01:43Z"
}
```

PUT /events/:id (Authenticated)

```json
{
  "name": "Updated Event Name",
  "description": "Updated Description",
  "location": "Updated Location",
  "datetime": "2025-05-25T19:01:43Z"
}
```

DELETE /events/:id (Authenticated)

- No payload required

### Event Registration Endpoints

POST /events/:id/register (Authenticated)

- No payload required
- Registers current user for event

POST /events/:id/cancel-register (Authenticated)

- No payload required
- Cancels user's registration for event

## Authentication

- All authenticated routes require a JWT token in the Authorization header
- Token is obtained from the login endpoint
- Token format: Bearer <token>

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
)
```

### Events Table

```sql
CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
)
```

### Registrations Table

```sql
CREATE TABLE registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    event_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
    FOREIGN KEY(event_id) REFERENCES events(id)
)
```
