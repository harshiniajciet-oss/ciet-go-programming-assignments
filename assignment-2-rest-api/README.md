# ASSIGNMENT 2: Design and Develop a RESTful Web Service using Go.

## Academic Credentials
- **Student Name:** Ezhilarasi K
- **Register Number:** 24BCS025
- **Institution:** Coimbatore Institute of Engineering and Technology (CIET)

---

## Project Overview
This project provides a complete, robust, production-style RESTful Web Service built exclusively with Go's standard library (`net/http`). The service manages academic **Student** resources and adheres strictly to REST principles, JSON data interchange, standard HTTP response status codes, structured error handling, request logging middleware, and in-memory thread-safe state management.

---

## Student Resource Schema

| Field | Type | Description | Validation Rule |
| :--- | :--- | :--- | :--- |
| `id` | `string` | Unique Student Identifier (e.g., `24BCS025`) | Auto-generated if omitted (`CIET####`) |
| `name` | `string` | Full Name of the student | Required, non-empty string |
| `department` | `string` | Academic department/branch | Required, non-empty string |
| `year` | `integer` | Current academic year | Required, must be integer between 1 and 5 |
| `email` | `string` | Official institutional email | Required, valid RFC email format |

> [!NOTE]
> **In-Memory Data Store:**
> All student records are maintained in RAM using a thread-safe repository backed by `sync.RWMutex`. The store is seeded with sample records upon startup. As with all in-memory databases, records reset to their initial state whenever the server process restarts.

---

## REST Endpoints Specification

| Method | Endpoint | Description | Success Status | Error Statuses |
| :--- | :--- | :--- | :--- | :--- |
| `GET` | `/` | API Metadata & Author Info | `200 OK` | `404 Not Found` |
| `GET` | `/health` | Server Health Status | `200 OK` | `500 Internal Server Error` |
| `GET` | `/api/students` | Retrieve All Students | `200 OK` | `500 Internal Server Error` |
| `POST` | `/api/students` | Create New Student Record | `201 Created` | `400 Bad Request` |
| `GET` | `/api/students/{id}` | Retrieve Student by ID | `200 OK` | `400 Bad Request`, `404 Not Found` |
| `PUT` | `/api/students/{id}` | Update Student Record | `200 OK` | `400 Bad Request`, `404 Not Found` |
| `DELETE` | `/api/students/{id}`| Delete Student Record | `200 OK` | `400 Bad Request`, `404 Not Found` |

---

## Architecture & Directory Structure

```
assignment-2-rest-api/
├── cmd/
│   └── server/
│       └── main.go                  # Server startup, banner, graceful shutdown, routing
├── internal/
│   ├── handler/
│   │   ├── student_handler.go       # REST route handlers & JSON input/output helpers
│   │   └── student_handler_test.go  # Comprehensive API unit tests with httptest
│   ├── middleware/
│   │   ├── logger.go                # HTTP request logging middleware (status, latency, path)
│   │   └── logger_test.go           # Middleware unit test
│   ├── model/
│   │   └── student.go               # Student model, input validation & APIResponse schemas
│   └── store/
│       ├── memory_store.go          # Thread-safe in-memory store with sync.RWMutex
│       └── memory_store_test.go     # CRUD and concurrency safety tests
├── go.mod                           # Go module declaration
└── README.md                        # Project documentation
```

---

## Installation & Running the Server

### Prerequisites
- Go 1.22+ installed.

### 1. Start the API Server
Navigate to the assignment directory and run:
```bash
cd assignment-2-rest-api
go run ./cmd/server/main.go
```
By default, the server starts on port `8080`. To customize the listening port:
```bash
go run ./cmd/server/main.go -port 9000
# or set environment variable
PORT=9000 go run ./cmd/server/main.go
```

---

## Testing API Endpoints with `curl`

### 1. Health Check
```bash
curl -i http://localhost:8080/health
```
**Response:**
```json
{
  "success": true,
  "message": "Service is healthy and operational"
}
```

### 2. Retrieve All Students
```bash
curl -i http://localhost:8080/api/students
```
**Response (200 OK):**
```json
{
  "success": true,
  "message": "Retrieved 3 student record(s)",
  "data": [
    {
      "id": "23BIT012",
      "name": "Kavitha M",
      "department": "Information Technology",
      "year": 3,
      "email": "kavitha.m@ciet.edu.in"
    },
    {
      "id": "24BCS001",
      "name": "Aarav Sharma",
      "department": "Computer Science and Engineering",
      "year": 2,
      "email": "aarav.sharma@ciet.edu.in"
    },
    {
      "id": "24BCS025",
      "name": "Ezhilarasi K",
      "department": "Computer Science and Engineering",
      "year": 2,
      "email": "ezhilarasi.24bcs025@ciet.edu.in"
    }
  ]
}
```

### 3. Retrieve Student by ID
```bash
curl -i http://localhost:8080/api/students/24BCS025
```
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "24BCS025",
    "name": "Ezhilarasi K",
    "department": "Computer Science and Engineering",
    "year": 2,
    "email": "ezhilarasi.24bcs025@ciet.edu.in"
  }
}
```

### 4. Create a New Student
```bash
curl -i -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "id": "24BCS050",
    "name": "Rohit Verma",
    "department": "Computer Science and Engineering",
    "year": 1,
    "email": "rohit.verma@ciet.edu.in"
  }'
```
**Response (201 Created):**
```json
{
  "success": true,
  "message": "Student created successfully",
  "data": {
    "id": "24BCS050",
    "name": "Rohit Verma",
    "department": "Computer Science and Engineering",
    "year": 1,
    "email": "rohit.verma@ciet.edu.in"
  }
}
```

### 5. Update an Existing Student
```bash
curl -i -X PUT http://localhost:8080/api/students/24BCS050 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Rohit Verma",
    "department": "Computer Science and Engineering",
    "year": 2,
    "email": "rohit.v2@ciet.edu.in"
  }'
```
**Response (200 OK):**
```json
{
  "success": true,
  "message": "Student updated successfully",
  "data": {
    "id": "24BCS050",
    "name": "Rohit Verma",
    "department": "Computer Science and Engineering",
    "year": 2,
    "email": "rohit.v2@ciet.edu.in"
  }
}
```

### 6. Delete a Student
```bash
curl -i -X DELETE http://localhost:8080/api/students/24BCS050
```
**Response (200 OK):**
```json
{
  "success": true,
  "message": "Student with ID '24BCS050' deleted successfully"
}
```

### 7. Validation Error Demonstration (Bad Request)
```bash
curl -i -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pooja S",
    "department": "ECE",
    "year": 8,
    "email": "invalid-email-string"
  }'
```
**Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "Validation error: academic year must be between 1 and 5: received 8"
}
```

### 8. Resource Not Found Demonstration
```bash
curl -i http://localhost:8080/api/students/UNKNOWN_ID
```
**Response (404 Not Found):**
```json
{
  "success": false,
  "error": "Student with ID 'UNKNOWN_ID' not found"
}
```

---

## Running Automated Tests

Run the full test suite including endpoint simulation with `httptest`:
```bash
go test -v ./...
```

Run test suite with test coverage metrics:
```bash
go test -v -cover ./...
```
