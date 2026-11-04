package display

import (
	"bytes"
	"strings"
	"testing"

	"assignment1/internal/model"
)

func TestPrintBanner(t *testing.T) {
	buf := new(bytes.Buffer)
	PrintBanner(buf)
	out := buf.String()

	if !strings.Contains(out, "Ezhilarasi K") {
		t.Errorf("Banner missing student name")
	}
	if !strings.Contains(out, "24BCS025") {
		t.Errorf("Banner missing register number")
	}
	if !strings.Contains(out, "Coimbatore Institute of Engineering and Technology (CIET)") {
		t.Errorf("Banner missing institution name")
	}
}

func TestPrintResultsTable(t *testing.T) {
	students := []model.Student{
		{
			Name:  "Ezhilarasi K",
			RegNo: "24BCS025",
			Subjects: []model.Subject{
				{Name: "Go", Marks: 95},
				{Name: "DS", Marks: 90},
			},
			TotalMarks: 185.0,
			Average:    92.5,
			Grade:      "A",
			IsPassed:   true,
		},
	}

	buf := new(bytes.Buffer)
	PrintResultsTable(buf, students)
	out := buf.String()

	if !strings.Contains(out, "24BCS025") || !strings.Contains(out, "Ezhilarasi K") {
		t.Errorf("Results table missing student data")
	}
	if !strings.Contains(out, "PASS") {
		t.Errorf("Results table missing PASS status")
	}
}

func TestPrintSummaryReport(t *testing.T) {
	summary := model.ClassSummary{
		TotalStudents:  2,
		PassedCount:    2,
		FailedCount:    0,
		PassPercentage: 100.0,
		ClassAverage:   92.5,
		HighestScorer:  "Ezhilarasi K (24BCS025)",
		HighestScore:   92.5,
		LowestScorer:   "Ezhilarasi K (24BCS025)",
		LowestScore:    92.5,
	}

	buf := new(bytes.Buffer)
	PrintSummaryReport(buf, summary)
	out := buf.String()

	if !strings.Contains(out, "CLASS PERFORMANCE SUMMARY") {
		t.Errorf("Summary report missing header")
	}
	if !strings.Contains(out, "100.00%") {
		t.Errorf("Summary report missing pass percentage")
	}
}
