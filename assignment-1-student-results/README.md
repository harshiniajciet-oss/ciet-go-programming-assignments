# ASSIGNMENT 1: Modular Go CLI Application for Student Result Processing with Version Control.

## Academic Credentials
- **Student Name:** Ezhilarasi K
- **Register Number:** 24BCS025
- **Institution:** Coimbatore Institute of Engineering and Technology (CIET)

---

## Project Overview
This project is a modular, production-ready Command-Line Interface (CLI) application developed in Go for academic result evaluation. The application handles student record entry, validates scores against academic standards, computes totals, averages, and letter grades, clearly flags Pass/Fail status, and outputs a formatted performance table along with a class summary report.

The system is constructed with standard Go conventions (`go.mod`, standard library only) and organized into cleanly separated packages adhering to SOLID principles.

---

## Features
1. **Interactive and Demonstration Modes:**
   - **Interactive Mode:** Prompts user for student records, number of subjects, and subject marks dynamically.
   - **Demo Mode (`--demo`):** Quickly loads a realistic student sample representing all letter grades and pass/fail conditions for fast grading and automated validation.
2. **Robust Input Validation:**
   - Enforces marks to be within the range `[0.0, 100.0]`.
   - Prevents empty student names and blank register numbers.
   - Ensures positive integer counts for students and subjects.
3. **Accurate Calculation & Grading System:**
   - Total Marks: Sum of marks across all enrolled subjects.
   - Average Marks: Rounded to two decimal places.
   - **Grading Rules:**
     - `A`: 90.00 – 100.00
     - `B`: 80.00 – 89.99
     - `C`: 70.00 – 79.99
     - `D`: 60.00 – 69.99
     - `E`: 50.00 – 59.99
     - `F`: Below 50.00 (or if any subject mark is below 50.00)
   - **Pass/Fail Criteria:** A student passes (`PASS`) only if each subject mark is $\ge 50$ and the overall average is $\ge 50$. If a student scores below 50 in any subject, they are marked `FAIL` with grade `F`.
4. **Readable Tabular Output & Class Analytics:**
   - Formatted ASCII table containing Register Number, Student Name, Subject Scores, Total, Average, Grade, and Status.
   - Class performance summary showing total students, passed/failed count, pass percentage, class average, highest scorer, and lowest scorer.
5. **Comprehensive Unit Testing:**
   - Table-driven unit tests across `calculator`, `validator`, and `display` packages covering edge cases, boundary marks, and validation errors.

---

## Architecture & Package Structure

```
assignment-1-student-results/
├── cmd/
│   └── main.go                  # CLI entry point, menu handling & interactive loop
├── internal/
│   ├── calculator/
│   │   ├── calculator.go        # Calculation of totals, averages, grades & class summary
│   │   └── calculator_test.go   # Table-driven unit tests for calculations & grading rules
│   ├── display/
│   │   ├── display.go           # Banner, ASCII table & summary report formatting
│   │   └── display_test.go      # Unit tests for output renderers
│   ├── model/
│   │   └── student.go           # Data structures: Student, Subject, ClassSummary
│   └── validator/
│       ├── validator.go         # Input validation rules (marks 0-100, non-empty fields)
│       └── validator_test.go    # Unit tests for all validation constraints
├── go.mod                       # Go module declaration
└── README.md                    # Project documentation
```

---

## Installation & Execution

### Prerequisites
- Go 1.22 or higher installed. Check with:
  ```bash
  go version
  ```

### Run the Application

Navigate to the project folder:
```bash
cd assignment-1-student-results
```

#### 1. Quick Demonstration Mode
To run with pre-populated student records across all grade bands:
```bash
go run ./cmd/main.go --demo
```

#### 2. Interactive Console Mode
To launch the interactive CLI where you can type student details:
```bash
go run ./cmd/main.go
```

#### 3. Build Executable Binary
```bash
go build -o student-results ./cmd/main.go
./student-results --demo
```
*(On Windows: `go build -o student-results.exe .\cmd\main.go` and run `.\student-results.exe --demo`)*

---

## Running Unit Tests

Run all unit tests with verbose output:
```bash
go test -v ./...
```

Run tests with test coverage analysis:
```bash
go test -v -cover ./...
```

---

## Sample Terminal Output

### Demo Mode Execution Output

```
================================================================================
   COIMBATORE INSTITUTE OF ENGINEERING AND TECHNOLOGY (CIET)
   ASSIGNMENT 1: Modular Go CLI Application for Student Result Processing
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
