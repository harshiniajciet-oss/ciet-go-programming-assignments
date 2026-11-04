package calculator

import (
	"math"
	"testing"

	"assignment1/internal/model"
)

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name     string
		subjects []model.Subject
		want     float64
	}{
		{
			name:     "Empty subjects",
			subjects: []model.Subject{},
			want:     0.0,
		},
		{
			name: "Single subject",
			subjects: []model.Subject{
				{Name: "Go", Marks: 85.5},
			},
			want: 85.5,
		},
		{
			name: "Multiple subjects",
			subjects: []model.Subject{
				{Name: "Math", Marks: 90.0},
				{Name: "Physics", Marks: 80.5},
				{Name: "Chemistry", Marks: 70.0},
			},
			want: 240.5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateTotal(tc.subjects)
			if math.Abs(got-tc.want) > 0.001 {
				t.Errorf("CalculateTotal() = %.2f; want %.2f", got, tc.want)
			}
		})
	}
}

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name    string
		total   float64
		count   int
		want    float64
		wantErr bool
	}{
		{"Normal average", 270.0, 3, 90.0, false},
		{"Fractional average", 265.0, 3, 88.33, false},
		{"Zero count", 100.0, 0, 0.0, true},
		{"Negative count", 100.0, -1, 0.0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateAverage(tc.total, tc.count)
			if (err != nil) != tc.wantErr {
				t.Fatalf("CalculateAverage() error = %v; wantErr = %v", err, tc.wantErr)
			}
			if !tc.wantErr && math.Abs(got-tc.want) > 0.01 {
				t.Errorf("CalculateAverage() = %.2f; want %.2f", got, tc.want)
			}
		})
	}
}

func TestDetermineGrade(t *testing.T) {
	tests := []struct {
		name             string
		average          float64
		hasFailedSubject bool
		want             string
	}{
		{"Top mark 100", 100.0, false, "A"},
		{"Grade A lower boundary", 90.0, false, "A"},
		{"Grade B upper boundary", 89.9, false, "B"},
		{"Grade B lower boundary", 80.0, false, "B"},
		{"Grade C upper boundary", 79.9, false, "C"},
		{"Grade C lower boundary", 70.0, false, "C"},
		{"Grade D upper boundary", 69.9, false, "D"},
		{"Grade D lower boundary", 60.0, false, "D"},
		{"Grade E upper boundary", 59.9, false, "E"},
		{"Grade E lower boundary", 50.0, false, "E"},
		{"Grade F below 50", 49.9, false, "F"},
		{"Grade F zero", 0.0, false, "F"},
		{"High average but failed subject", 92.0, true, "F"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DetermineGrade(tc.average, tc.hasFailedSubject)
			if got != tc.want {
				t.Errorf("DetermineGrade(%.1f, %v) = %s; want %s", tc.average, tc.hasFailedSubject, got, tc.want)
			}
		})
	}
}

func TestHasFailedSubject(t *testing.T) {
	tests := []struct {
		name     string
		subjects []model.Subject
		want     bool
	}{
		{
			name: "All passed",
			subjects: []model.Subject{
				{Name: "Go", Marks: 50.0},
				{Name: "C", Marks: 75.0},
			},
			want: false,
		},
		{
			name: "One subject failed",
			subjects: []model.Subject{
				{Name: "Go", Marks: 90.0},
				{Name: "Maths", Marks: 49.5},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := HasFailedSubject(tc.subjects)
			if got != tc.want {
				t.Errorf("HasFailedSubject() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestProcessStudent(t *testing.T) {
	t.Run("Pass student", func(t *testing.T) {
		s := model.Student{
			Name:  "Ezhilarasi K",
			RegNo: "24BCS025",
			Subjects: []model.Subject{
				{Name: "Algorithms", Marks: 95},
				{Name: "Cloud Computing", Marks: 92},
				{Name: "Go Programming", Marks: 98},
			},
		}

		err := ProcessStudent(&s)
		if err != nil {
			t.Fatalf("ProcessStudent failed: %v", err)
		}

		if s.TotalMarks != 285 {
			t.Errorf("got total %.2f; want 285", s.TotalMarks)
		}
		if s.Average != 95 {
			t.Errorf("got average %.2f; want 95", s.Average)
		}
		if s.Grade != "A" {
			t.Errorf("got grade %s; want A", s.Grade)
		}
		if !s.IsPassed {
			t.Errorf("got IsPassed false; want true")
		}
	})

	t.Run("Fail student due to subject", func(t *testing.T) {
		s := model.Student{
			Name:  "Student Two",
			RegNo: "24BCS099",
			Subjects: []model.Subject{
				{Name: "Algorithms", Marks: 95},
				{Name: "Go Programming", Marks: 45}, // failed subject
			},
		}

		err := ProcessStudent(&s)
		if err != nil {
			t.Fatalf("ProcessStudent failed: %v", err)
		}

		if s.Grade != "F" {
			t.Errorf("got grade %s; want F", s.Grade)
		}
		if s.IsPassed {
			t.Errorf("got IsPassed true; want false")
		}
	})

	t.Run("Student with no subjects error", func(t *testing.T) {
		s := model.Student{
			Name:     "Student Three",
			RegNo:    "24BCS100",
			Subjects: []model.Subject{},
		}

		err := ProcessStudent(&s)
		if err == nil {
			t.Errorf("expected error for empty subjects, got nil")
		}
	})
}

func TestCalculateClassSummary(t *testing.T) {
	students := []model.Student{
		{
			Name:       "Ezhilarasi K",
			RegNo:      "24BCS025",
			Average:    95.0,
			Grade:      "A",
			IsPassed:   true,
			TotalMarks: 285.0,
		},
		{
			Name:       "Aravind R",
			RegNo:      "24BCS003",
			Average:    82.0,
			Grade:      "B",
			IsPassed:   true,
			TotalMarks: 246.0,
		},
		{
			Name:       "Deepa M",
			RegNo:      "24BCS012",
			Average:    48.0,
			Grade:      "F",
			IsPassed:   false,
			TotalMarks: 144.0,
		},
	}

	summary := CalculateClassSummary(students)

	if summary.TotalStudents != 3 {
		t.Errorf("got TotalStudents %d, want 3", summary.TotalStudents)
	}
	if summary.PassedCount != 2 {
		t.Errorf("got PassedCount %d, want 2", summary.PassedCount)
	}
	if summary.FailedCount != 1 {
		t.Errorf("got FailedCount %d, want 1", summary.FailedCount)
	}
	if math.Abs(summary.PassPercentage-66.67) > 0.01 {
		t.Errorf("got PassPercentage %.2f, want 66.67", summary.PassPercentage)
	}
	if math.Abs(summary.ClassAverage-75.0) > 0.01 {
		t.Errorf("got ClassAverage %.2f, want 75.0", summary.ClassAverage)
	}
	if summary.HighestScore != 95.0 {
		t.Errorf("got HighestScore %.2f, want 95.0", summary.HighestScore)
	}
	if summary.LowestScore != 48.0 {
		t.Errorf("got LowestScore %.2f, want 48.0", summary.LowestScore)
	}
}
