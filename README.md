# CIET Go Programming Assignments

## Academic Information
- **Student Name:** Ezhilarasi K
- **Register Number:** 24BCS025
- **Institution:** Coimbatore Institute of Engineering and Technology (CIET)

---

## Repository Structure & Overview

This GitHub-ready repository houses two independent, fully functional Go programming assignments developed as part of academic coursework at the **Coimbatore Institute of Engineering and Technology (CIET)**. Neither project relies on heavyweight third-party packages; both strictly use Go's standard library and adhere to idiomatic Go package conventions.

```
ciet-go-programming-assignments/
├── .github/
│   └── workflows/
│       └── test.yml                 # GitHub Actions CI workflow (tests both projects independently)
├── assignment-1-student-results/    # Modular CLI application for Student Results
│   ├── cmd/
│   │   └── main.go                  # CLI entry point, interactive input & demonstration mode
│   ├── internal/
│   │   ├── calculator/
│   │   │   ├── calculator.go        # Totals, averages, grades & class statistical analytics
│   │   │   └── calculator_test.go   # Table-driven unit tests for calculations
│   │   ├── display/
│   │   │   ├── display.go           # Banner, ASCII table & summary report formatting
│   │   │   └── display_test.go      # Output tests
│   │   ├── model/
│   │   │   └── student.go           # Subject, Student, ClassSummary structs
│   │   └── validator/
│   │       ├── validator.go         # Input validation (marks [0-100], non-empty fields)
│   │       └── validator_test.go    # Validation unit tests
│   ├── go.mod                       # Module definition: assignment1
│   └── README.md                    # Detailed documentation for Assignment 1
├── assignment-2-rest-api/           # RESTful Web Service using Go net/http
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # HTTP server initialization & graceful shutdown
│   ├── internal/
│   │   ├── handler/
│   │   │   ├── student_handler.go   # REST handlers (GET, POST, PUT, DELETE)
│   │   │   └── student_handler_test.go # Comprehensive httptest endpoint tests
│   │   ├── middleware/
│   │   │   ├── logger.go            # Request logging middleware (latency, status, method)
│   │   │   └── logger_test.go       # Middleware tests
│   │   ├── model/
│   │   │   └── student.go           # Student model, input validation & APIResponse schemas
│   │   └── store/
│   │       ├── memory_store.go      # Thread-safe in-memory store with sync.RWMutex
│   │       └── memory_store_test.go # Store CRUD and concurrency tests
│   ├── go.mod                       # Module definition: assignment2
│   └── README.md                    # Detailed documentation for Assignment 2
├── .gitignore                       # Ignored binaries, IDE files & temporary artifacts
├── LICENSE                          # MIT License
└── README.md                        # Master repository documentation
```

---

## Assignment Summaries

### ASSIGNMENT 1: Modular Go CLI Application for Student Result Processing with Version Control.
- **Goal:** Build a modular CLI application that accepts student details (name, register number, and multiple subject marks), validates all inputs against the range `[0, 100]`, calculates totals, averages, and assigns letter grades based on clear academic rules.
- **Grading Scale:**
  - `A`: 90 – 100
  - `B`: 80 – 89.99
  - `C`: 70 – 79.99
  - `D`: 60 – 69.99
  - `E`: 50 – 59.99
  - `F`: Below 50 (or any failed subject mark $< 50$)
- **Key Capabilities:**
  - Interactive console input mode.
  - Demonstration mode via `--demo` flag for instant evaluation.
  - Slices and custom structs for multiple student and subject aggregation.
  - Clean tabular output with column-aligned formatting and class statistical summary.
  - Complete unit test suite with table-driven tests.

### ASSIGNMENT 2: Design and Develop a RESTful Web Service using Go.
- **Goal:** Develop a production-grade RESTful web service for managing `Student` resources using Go's built-in `net/http` package without external framework dependencies.
- **Student Entity:**
  - `id` (string), `name` (string), `department` (string), `year` (int, 1–5), `email` (string, validated RFC format).
- **CRUD Endpoints:**
  - `POST /api/students` (201 Created)
  - `GET /api/students` (200 OK)
  - `GET /api/students/{id}` (200 OK / 404 Not Found)
  - `PUT /api/students/{id}` (200 OK / 404 Not Found)
  - `DELETE /api/students/{id}` (200 OK / 404 Not Found)
- **Key Capabilities:**
  - Request logging middleware measuring method, path, HTTP status, and request duration.
  - Graceful server shutdown upon OS interrupt signals (`SIGINT`, `SIGTERM`).
  - Thread-safe in-memory repository utilizing `sync.RWMutex`.
  - Comprehensive API endpoint testing using `net/http/httptest`.

---

## Quick Start & Execution

### Prerequisites
- Go 1.22 or higher installed.

### Running Assignment 1 (CLI Application)
```bash
cd assignment-1-student-results

# Run unit tests
go test -v ./...

# Run in demonstration mode
go run ./cmd/main.go --demo

# Run in interactive mode
go run ./cmd/main.go
```

### Running Assignment 2 (RESTful Web Service)
```bash
cd assignment-2-rest-api

# Run unit tests
go test -v ./...

# Start the web server (defaults to port 8080)
go run ./cmd/server/main.go

# Start on a custom port
go run ./cmd/server/main.go -port 9000
```

---

## Sample Execution Outputs

### Assignment 1: Student Result Report & Summary
```
================================================================================
   Coimbatore Institute of Engineering and Technology (CIET)
   ASSIGNMENT 1: Modular Go CLI Application for Student Result Processing with Version Control.
   Student Name: Ezhilarasi K   |   Register Number: 24BCS025
================================================================================

[*] Running in DEMO mode with pre-configured student dataset...

----------------------------------------------------------------------------------------------------
STUDENT ACADEMIC PERFORMANCE REPORT
----------------------------------------------------------------------------------------------------
Reg No       | Student Name         | Subjects (Marks)               | Total  | Average | Grade | Status
----------------------------------------------------------------------------------------------------
24BCS025     | Ezhilarasi K         | Data Structures:96, Operati... | 286.0  | 95.33   | A     | PASS  
24BCS001     | Aarav Sharma         | Data Structures:82, Operati... | 254.0  | 84.67   | B     | PASS  
24BCS015     | Bhavana R            | Data Structures:74, Operati... | 224.0  | 74.67   | C     | PASS  
24BCS020     | Dinesh Kumar         | Data Structures:62, Operati... | 195.0  | 65.00   | D     | PASS  
24BCS028     | Elango V             | Data Structures:52, Operati... | 157.0  | 52.33   | E     | PASS  
24BCS035     | Ganesh M             | Data Structures:78, Operati... | 205.0  | 68.33   | F     | FAIL  
----------------------------------------------------------------------------------------------------

--------------------------------------------------
           CLASS PERFORMANCE SUMMARY
--------------------------------------------------
  Total Students Evaluated  : 6
  Students Passed           : 5
  Students Failed           : 1
  Overall Pass Percentage   : 83.33%
  Class Average Mark        : 73.39
  Highest Scorer            : Ezhilarasi K (24BCS025) (95.33)
  Lowest Scorer             : Elango V (24BCS028) (52.33)
--------------------------------------------------
```

### Assignment 2: REST API Terminal Requests & Responses
```bash
# Health check
curl -s http://localhost:8080/health
# {"success":true,"message":"Service is healthy and operational"}

# Fetch student by ID
curl -s http://localhost:8080/api/students/24BCS025
# {"success":true,"data":{"id":"24BCS025","name":"Ezhilarasi K","department":"Computer Science and Engineering","year":2,"email":"ezhilarasi.24bcs025@ciet.edu.in"}}
```

---

## Continuous Integration (GitHub Actions)

A GitHub Actions workflow is configured at [`.github/workflows/test.yml`](.github/workflows/test.yml) that automatically executes the test suites of both projects in parallel on every push and pull request to `main`:

```yaml
name: CIET Go Programming Assignments CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  test-assignment-1:
    name: Test Assignment 1 - Student Results CLI
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - working-directory: assignment-1-student-results
        run: go test -v -race -cover ./...

  test-assignment-2:
    name: Test Assignment 2 - RESTful Web Service
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - working-directory: assignment-2-rest-api
        run: go test -v -race -cover ./...
```

---

## Git Workflow & Publishing Guide

### Meaningful Git Commit History Suggestions
When committing your work locally, use structured Conventional Commits:

```bash
git commit -m "feat(assignment-1): implement student results calculator and grading rules"
git commit -m "test(assignment-1): add table-driven unit tests for validation and calculations"
git commit -m "feat(assignment-2): create RESTful student CRUD API using net/http"
git commit -m "feat(assignment-2): add request logging middleware and in-memory store"
git commit -m "test(assignment-2): add httptest suite for all HTTP endpoints"
git commit -m "ci: configure GitHub Actions workflow for automated multi-project testing"
git commit -m "docs: add comprehensive README documentation and execution guides"
```

### Creating and Merging a Feature Branch
To follow best practices for version control:
```bash
# 1. Create and switch to a feature branch
git checkout -b feature/assignment-implementation

# 2. Add and commit your changes
git add .
git commit -m "feat: complete assignment 1 and assignment 2 implementations"

# 3. Switch back to the main branch
git checkout main

# 4. Merge the feature branch into main
git merge feature/assignment-implementation

# 5. Delete the feature branch after merging
git branch -d feature/assignment-implementation
```

### Publishing the Repository to Your GitHub Account
Follow these exact steps in your terminal to publish this repository to GitHub:

1. **Initialize Git inside the root directory:**
   ```bash
   cd ciet-go-programming-assignments
   git init
   git branch -M main
   ```

2. **Stage and commit all project files:**
   ```bash
   git add .
   git commit -m "feat: initial commit for CIET Go programming assignments"
   ```

3. **Create a new empty repository on GitHub:**
   - Log in to your GitHub account at [https://github.com/new](https://github.com/new).
   - Enter `ciet-go-programming-assignments` as the repository name.
   - Set visibility to **Public** (or **Private** as desired).
   - **Do NOT** check "Initialize this repository with a README, .gitignore, or license" (these already exist).
   - Click **Create repository**.

4. **Link your local repository to your remote GitHub repository:**
   *(Replace `<YOUR-GITHUB-USERNAME>` with your actual GitHub username)*
   ```bash
   git remote add origin https://github.com/<YOUR-GITHUB-USERNAME>/ciet-go-programming-assignments.git
   ```

5. **Push to GitHub:**
   ```bash
   git push -u origin main
   ```
