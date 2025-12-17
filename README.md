User API - Go RESTful Service
A RESTful API built with Go to manage users with name and date of birth, featuring dynamic age calculation.
Tech Stack

GoFiber
PostgreSQL
Uber Zap (Logging)
go-playground/validator (Validation)

Project Structure
user-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── user_handler.go
│   ├── service/
│   │   └── user_service.go
│   ├── repository/
│   │   ├── db.go
│   │   └── user_repository.go
│   ├── models/
│   │   ├── user.go
│   │   ├── user_request.go
│   │   └── user_response.go
│   └── logger/
│       └── logger.go
├── db/
│   ├── migrations/
│   │   └── 001_users.sql
│   ├── query/
│   │   └── users.sql
│   └── sqlc/
├── sqlc.yaml
├── go.mod
├── go.sum
└── README.md
Setup Instructions
Prerequisites

Go 1.21+
PostgreSQL 15+

Installation

Clone the repository
Install dependencies

bashgo mod download

Set up PostgreSQL database

bashcreatedb userdb
psql -U postgres -d userdb -f db/migrations/001_users.sql

Configure environment variables (optional)

bashexport DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=userdb

Run the server

bashgo run cmd/server/main.go
Server will start on http://localhost:3000
API Endpoints
Create User
POST /users
Request:
json{
"name": "Alice Johnson",
"dob": "1990-05-10"
}
Response (201):
json{
"id": 1,
"name": "Alice Johnson",
"dob": "1990-05-10"
}
Get User by ID
GET /users/:id
Response (200):
json{
"id": 1,
"name": "Alice Johnson",
"dob": "1990-05-10",
"age": 34
}
List All Users
GET /users
Response (200):
json[
{
"id": 1,
"name": "Alice Johnson",
"dob": "1990-05-10",
"age": 34
}
]
Update User
PUT /users/:id
Request:
json{
"name": "Alice Updated",
"dob": "1991-03-15"
}
Response (200):
json{
"id": 1,
"name": "Alice Updated",
"dob": "1991-03-15"
}
Delete User
DELETE /users/:id
Response: 204 No Content
Health Check
GET /health
Response (200):
json{
"status": "ok"
}
Testing with cURL
Create user:
bashcurl -X POST http://localhost:3000/users -H "Content-Type: application/json" -d '{"name":"Alice Johnson","dob":"1990-05-10"}'
Get all users:
bashcurl http://localhost:3000/users
Get specific user:
bashcurl http://localhost:3000/users/1
Update user:
bashcurl -X PUT http://localhost:3000/users/1 -H "Content-Type: application/json" -d '{"name":"Alice Updated","dob":"1991-03-15"}'
Delete user:
bashcurl -X DELETE http://localhost:3000/users/1
Features

RESTful API design with proper HTTP methods
Dynamic age calculation from date of birth
Input validation using go-playground/validator
Structured logging with Uber Zap
Clean layered architecture
Environment-based configuration
Graceful server shutdown
Comprehensive error handling
Health check endpoint

Database Schema
sqlCREATE TABLE users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
dob DATE NOT NULL
);
Dependencies
bashgo get github.com/gofiber/fiber/v2
go get github.com/lib/pq
go get go.uber.org/zap
go get github.com/go-playground/validator/v10
Error Responses
All errors return JSON format with appropriate status codes:

400 - Bad Request (invalid input)
404 - Not Found (user doesn't exist)
500 - Internal Server Error

Notes

Age is calculated dynamically on every request
Date format must be YYYY-MM-DD
All timestamps handled in UTC
Database credentials should be set via environment variables

Author
Niteesh
License
This project is open source and available under the MIT License.