package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"assignment1/internal/calculator"
	"assignment1/internal/display"
	"assignment1/internal/model"
	"assignment1/internal/validator"
)

func main() {
	demoMode := flag.Bool("demo", false, "Run with pre-configured sample student data")
	flag.Parse()

	display.PrintBanner(os.Stdout)

	var students []model.Student

	if *demoMode {
		fmt.Println("\n[*] Running in DEMO mode with pre-configured student dataset...")
		students = getSampleStudents()
	} else {
		fmt.Println("\nChoose input mode:")
		fmt.Println("  1. Interactive Console Entry")
		fmt.Println("  2. Load Sample Demonstration Data")
		fmt.Print("Enter choice (1 or 2): ")

		reader := bufio.NewReader(os.Stdin)
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		if choiceStr == "2" {
			students = getSampleStudents()
		} else {
			students = collectInteractiveStudents(reader)
		}
	}

	if len(students) == 0 {
		fmt.Println("\nNo student records to process. Exiting program.")
		return
	}

	// Process all student calculations
	for i := range students {
		if err := calculator.ProcessStudent(&students[i]); err != nil {
			fmt.Printf("Error processing student %s: %v\n", students[i].Name, err)
			return
		}
	}

	// Display Tabular Results
	display.PrintResultsTable(os.Stdout, students)

	// Display Class Summary
	summary := calculator.CalculateClassSummary(students)
	display.PrintSummaryReport(os.Stdout, summary)
}

// collectInteractiveStudents prompts the user for student details and subject marks with robust validation.
func collectInteractiveStudents(reader *bufio.Reader) []model.Student {
	var students []model.Student

	studentCount := readPositiveInt(reader, "Enter total number of students to evaluate: ")

	for sIdx := 1; sIdx <= studentCount; sIdx++ {
		fmt.Printf("\n--- Entering Details for Student #%d ---\n", sIdx)

		name := readNonEmptyString(reader, "Enter Student Name: ")
		regNo := readNonEmptyString(reader, "Enter Register Number: ")

		subjectCount := readPositiveInt(reader, fmt.Sprintf("Enter number of subjects for %s: ", name))
		var subjects []model.Subject

		for subIdx := 1; subIdx <= subjectCount; subIdx++ {
			subName := readNonEmptyString(reader, fmt.Sprintf("  Enter name of subject #%d: ", subIdx))
			marks := readValidMark(reader, fmt.Sprintf("  Enter marks for %s (0 - 100): ", subName))

			subjects = append(subjects, model.Subject{
				Name:  subName,
				Marks: marks,
			})
		}

		student := model.Student{
			Name:     name,
			RegNo:    regNo,
			Subjects: subjects,
		}

		if err := validator.ValidateStudent(student); err != nil {
			fmt.Printf("[Error] Invalid student data: %v. Please re-enter.\n", err)
			sIdx--
			continue
		}

		students = append(students, student)
	}

	return students
}

// readNonEmptyString prompts until a non-empty string is provided.
func readNonEmptyString(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			return line
		}
		fmt.Println(" [!] Input cannot be blank. Please try again.")
	}
}

// readPositiveInt prompts until a positive integer (> 0) is entered.
func readPositiveInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		val, err := strconv.Atoi(line)
		if err == nil && val > 0 {
			return val
		}
		fmt.Println(" [!] Invalid number. Please enter an integer greater than 0.")
	}
}

// readValidMark prompts until a valid mark between 0 and 100 is entered.
func readValidMark(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		mark, err := strconv.ParseFloat(line, 64)
		if err == nil {
			if valErr := validator.ValidateMark(mark); valErr == nil {
				return mark
			}
		}
		fmt.Println(" [!] Invalid mark. Marks must be a numerical value between 0.00 and 100.00.")
	}
}

// getSampleStudents returns a realistic dataset showcasing all grades and pass/fail conditions.
func getSampleStudents() []model.Student {
	return []model.Student{
		{
			Name:  "Ezhilarasi K",
			RegNo: "24BCS025",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 96},
				{Name: "Operating Systems", Marks: 92},
				{Name: "Go Programming", Marks: 98},
			},
		},
		{
			Name:  "Aarav Sharma",
			RegNo: "24BCS001",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 82},
				{Name: "Operating Systems", Marks: 84},
				{Name: "Go Programming", Marks: 88},
			},
		},
		{
			Name:  "Bhavana R",
			RegNo: "24BCS015",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 74},
				{Name: "Operating Systems", Marks: 78},
				{Name: "Go Programming", Marks: 72},
			},
		},
		{
			Name:  "Dinesh Kumar",
			RegNo: "24BCS020",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 62},
				{Name: "Operating Systems", Marks: 65},
				{Name: "Go Programming", Marks: 68},
			},
		},
		{
			Name:  "Elango V",
			RegNo: "24BCS028",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 52},
				{Name: "Operating Systems", Marks: 55},
				{Name: "Go Programming", Marks: 50},
			},
		},
		{
			Name:  "Ganesh M",
			RegNo: "24BCS035",
			Subjects: []model.Subject{
				{Name: "Data Structures", Marks: 78},
				{Name: "Operating Systems", Marks: 42}, // Failed single subject -> Grade F
				{Name: "Go Programming", Marks: 85},
			},
		},
	}
}
