package display

import (
	"fmt"
	"io"
	"strings"

	"assignment1/internal/model"
)

// PrintBanner outputs the project and academic author details.
func PrintBanner(w io.Writer) {
	line := strings.Repeat("=", 80)
	fmt.Fprintln(w, line)
	fmt.Fprintln(w, "   Coimbatore Institute of Engineering and Technology (CIET)")
	fmt.Fprintln(w, "   ASSIGNMENT 1: Modular Go CLI Application for Student Result Processing with Version Control.")
	fmt.Fprintln(w, "   Student Name: Ezhilarasi K   |   Register Number: 24BCS025")
	fmt.Fprintln(w, line)
}

// PrintResultsTable prints a neatly formatted ASCII table showing each student's results.
func PrintResultsTable(w io.Writer, students []model.Student) {
	separator := strings.Repeat("-", 100)
	fmt.Fprintln(w, "\n"+separator)
	fmt.Fprintln(w, "STUDENT ACADEMIC PERFORMANCE REPORT")
	fmt.Fprintln(w, separator)
	fmt.Fprintf(w, "%-12s | %-20s | %-30s | %-6s | %-7s | %-5s | %-6s\n",
		"Reg No", "Student Name", "Subjects (Marks)", "Total", "Average", "Grade", "Status")
	fmt.Fprintln(w, separator)

	for _, s := range students {
		var subDetails []string
		for _, sub := range s.Subjects {
			subDetails = append(subDetails, fmt.Sprintf("%s:%.0f", sub.Name, sub.Marks))
		}
		subjectsStr := strings.Join(subDetails, ", ")
		if len(subjectsStr) > 30 {
			subjectsStr = subjectsStr[:27] + "..."
		}

		status := "PASS"
		if !s.IsPassed {
			status = "FAIL"
		}

		fmt.Fprintf(w, "%-12s | %-20s | %-30s | %-6.1f | %-7.2f | %-5s | %-6s\n",
			s.RegNo, s.Name, subjectsStr, s.TotalMarks, s.Average, s.Grade, status)
	}
	fmt.Fprintln(w, separator)
}

// PrintSummaryReport outputs the class performance statistics in a structured box.
func PrintSummaryReport(w io.Writer, summary model.ClassSummary) {
	line := strings.Repeat("-", 50)
	fmt.Fprintln(w, "\n"+line)
	fmt.Fprintln(w, "           CLASS PERFORMANCE SUMMARY")
	fmt.Fprintln(w, line)
	fmt.Fprintf(w, "  %-25s : %d\n", "Total Students Evaluated", summary.TotalStudents)
	fmt.Fprintf(w, "  %-25s : %d\n", "Students Passed", summary.PassedCount)
	fmt.Fprintf(w, "  %-25s : %d\n", "Students Failed", summary.FailedCount)
	fmt.Fprintf(w, "  %-25s : %.2f%%\n", "Overall Pass Percentage", summary.PassPercentage)
	fmt.Fprintf(w, "  %-25s : %.2f\n", "Class Average Mark", summary.ClassAverage)
	fmt.Fprintf(w, "  %-25s : %s (%.2f)\n", "Highest Scorer", summary.HighestScorer, summary.HighestScore)
	fmt.Fprintf(w, "  %-25s : %s (%.2f)\n", "Lowest Scorer", summary.LowestScorer, summary.LowestScore)
	fmt.Fprintln(w, line)
}
