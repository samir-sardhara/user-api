# User Management API

A RESTful User Management API built with Go, Fiber, MySQL, and SQLC.

This project follows a layered architecture with separate Handler, Service, and Repository layers for maintainability and scalability.

## Features

* CRUD operations for users
* MySQL database integration
* SQLC for type-safe database access
* Request validation using go-playground/validator
* Structured logging with Zap
* Pagination support
* Clean layered architecture
* Middleware for request logging and latency tracking

---

## Tech Stack

* Go
* Fiber
* MySQL
* SQLC
* Zap Logger
* Validator

---

## Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── db/
│   ├── migrations/
│   │   └── schema.sql
│   ├── queries.sql
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── logger/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── service/
├── go.mod
├── go.sum
└── sqlc.yaml
```

## Prerequisites

Make sure the following are installed:

* Go 1.21+
* MySQL
* SQLC

Install SQLC:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

---

## Database Setup

Login to MySQL:

```bash
mysql -u root -p
```

Create the database:

```sql
CREATE DATABASE IF NOT EXISTS userdb;

USE userdb;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);
```

---

## Generate SQLC Code

Run:

```bash
sqlc generate
```

This generates type-safe Go code inside:

```text
db/sqlc/
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run the Application

```bash
go run cmd/server/main.go
```

Server starts on:

```text
http://localhost:3000
```

---

## API Endpoints

### Create User

**POST** `/users`

Request:

```json
{
  "name": "Alice Smith",
  "dob": "1995-08-24"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice Smith",
  "dob": "1995-08-24"
}
```

---

### Get User By ID

**GET** `/users/:id`

Example:

```http
GET /users/1
```

---

### Get Users

**GET** `/users?page=1&limit=10`

Example Response:

```json
[
  {
    "id": 1,
    "name": "Alice Smith",
    "dob": "1995-08-24"
  }
]
```

---

### Update User

**PUT** `/users/:id`

Request:

```json
{
  "name": "Alice Johnson",
  "dob": "1995-08-24"
}
```

---

### Delete User

**DELETE** `/users/:id`

---

## Testing

Run all tests:

```bash
go test ./... -v
```

Run service tests only:

```bash
go test ./internal/service/... -v
```

---

## Logging

The application uses Zap Logger for structured logging.

Logged information includes:

* HTTP method
* Request path
* Response status
* Request latency
* Application errors

Example:

```json
{
  "level":"info",
  "msg":"HTTP Request",
  "method":"GET",
  "path":"/users",
  "status":200,
  "latency":"12ms"
}
```




```
```
