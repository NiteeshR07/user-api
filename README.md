# User API – Go RESTful Service
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![GoFiber](https://img.shields.io/badge/Framework-GoFiber-00ACD7)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?logo=postgresql&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Active-success)


A RESTful API built with Go to manage users with name and date of birth, featuring dynamic age calculation.

---

## Tech Stack

- GoFiber
- PostgreSQL
- Uber Zap (Logging)
- go-playground/validator (Validation)

---

## Prerequisites

- Go 1.21+
- PostgreSQL 15+

---

## Project Structure

```text
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
```

---

## Installation

### Clone the repository
```bash
git clone https://github.com/NiteeshR07/user-api.git
cd user-api
```

### Install dependencies
```bash
go mod download
```

### Set up PostgreSQL database
```bash
createdb userdb
psql -U postgres -d userdb -f db/migrations/001_users.sql
```

### Configure environment variables (optional)
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=userdb
```

### Run the server
```bash
go run cmd/server/main.go
```

Server starts at  
http://localhost:3000

---

## CRUD Operations (cURL Examples)

### Create User
```bash
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice Johnson",
  "dob": "1990-05-10"
}'
```

### Get All Users
```bash
curl http://localhost:3000/users
```

### Get User by ID
```bash
curl http://localhost:3000/users/1
```

### Update User
```bash
curl -X PUT http://localhost:3000/users/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}'
```

### Delete User
```bash
curl -X DELETE http://localhost:3000/users/1
```

### Health Check
```bash
curl http://localhost:3000/health
```

---

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## Error Responses

- 400 – Bad Request
- 404 – Not Found
- 500 – Internal Server Error

All errors return JSON with appropriate status codes.

---

## Notes

- Age is calculated dynamically on every request
- Date format must be YYYY-MM-DD
- All timestamps handled in UTC
- Database credentials should be set via environment variables

---

## Author

Niteesh

---

## License

MIT License
