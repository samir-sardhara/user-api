readme_content = """# User Management API

A high-performance, production-ready RESTful API built in Go using the **Fiber** web framework, a local **MySQL** database instance for persistence, and **SQLC** for generating type-safe, compile-time checked database routines.

This project implements a clean, decoupled Layered Architecture focusing entirely on local native execution.

---

## 🛠 Features

* **Layered Architecture:** Strict isolation of concerns across Routing, Handlers, Services, and Repositories.
* **Compile-Time Type-Safe SQL:** SQL queries compiled directly into native Go routines using SQLC.
* **Pagination Engine:** Scalable page-and-limit offset retrieval strategies on collection pathways.
* **Structured Telemetry Middleware:** Embedded unique context request IDs mapped directly with execution latency metrics via Zap logging.
* **Input Data Validation:** Strict body payload structural schema checking using `go-playground/validator`.

---

## 📂 Project Structure

Code output
File successfully created.

```text
├── cmd/
│   └── server/
│       └── main.go           # Application entrypoint & dependency injection matrix
├── config/
│   └── config.go         # Configuration loading layer (DSN management)
├── db/
│   ├── migrations/
│   │   └── schema.sql    # Database structural DDL migrations
│   ├── queries.sql       # Raw parameterized SQL operation queries
│   └── sqlc/             # Auto-generated type-safe code output
├── internal/
│   ├── handler/          # HTTP request binders & status code controllers
│   ├── logger/           # Uber Zap structured log core engine
│   ├── middleware/       # Telemetry logging & Request ID headers
│   ├── models/           # DTO request-response schemas
│   ├── repository/       # Direct data-store interaction methods
│   ├── routes/           # REST URI pathway endpoint definitions
│   └── service/          # Core business validation domain operations
└── sqlc.yaml             # Code generation target parameters configuration
🚀 Native Setup & Execution
Follow these step-by-step instructions to prepare your environment and run the application natively on your machine.

📋 Prerequisites
Ensure you have the following tools installed locally:

Go (v1.21+)

MySQL Server (running locally on port 3306)

SQLC CLI (brew install sqlc or go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest)

Step 1: Initialize the Database & Schema
Log into your local MySQL CLI tool via your terminal:

Bash
mysql -u root -p
(Enter your local password when prompted).

Create the target database and execute the table schema definition:

SQL
-- Create the database
CREATE DATABASE IF NOT EXISTS userdb;

-- Select the target database context
USE userdb;

-- Create the users structural database table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);
Type exit to leave the MySQL prompt.

Step 2: Generate Type-Safe Go Code
Navigate to the project root directory (user-api/) and compile your SQL definitions into native Go database code layer files using SQLC:

Bash
sqlc generate
This maps your schema.sql and queries.sql into programmatic functions inside the /db/sqlc/ folder automatically.

Step 3: Start the Server Application
Fetch and synchronize your package tracking modules, then execute the runtime server entrypoint framework:

Bash
# Clean and pull runtime vendor dependencies
go mod tidy

# Launch the server instance process loop
go run cmd/server/main.go
The server will boot up instantly and listen on http://localhost:3000.

Step 4: Run Unit Tests
To verify internal business logic validation processing loops:

Bash
go test ./internal/service/... -v
🛰 API Endpoints & Testing Specification
All endpoints are accessible via http://localhost:3000.

1. Create a User
Method: POST

Path: /users

Headers: Content-Type: application/json

Payload Request Body Data:

JSON
{
  "name": "Alice Smith",
  "dob": "1995-08-24"
}
2. List Users (With Pagination Parameters)
Method: GET

Path: /users?page=1&limit=10

Response Payload Collection Output:

JSON
[
  {
    "id": 1,
    "name": "Alice Smith",
    "dob": "1995-08-24",
    "age": 31
  }
]
3. Get User By Primary Key ID
Method: GET

Path: /users/:id

4. Update Existing User Profile
Method: PUT

Path: /users/:id

Payload Request Body Data:

JSON
{
  "name": "Alice Thomson",
  "dob": "1995-08-24"
}
5. Remove User From Registry
Method: DELETE

Path: /users/:id

Expected Termination Status: 204 No Content
