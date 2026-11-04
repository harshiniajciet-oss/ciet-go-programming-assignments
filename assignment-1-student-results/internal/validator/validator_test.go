package validator

import (
	"errors"
	"testing"

	"assignment1/internal/model"
)

func TestValidateMark(t *testing.T) {
	tests := []struct {
		name    string
		mark    float64
		wantErr bool
	}{
		{"Valid zero mark", 0.0, false},
		{"Valid pass mark", 50.0, false},
		{"Valid full mark", 100.0, false},
		{"Valid fractional mark", 87.5, false},
		{"Invalid negative mark", -1.0, true},
		{"Invalid above 100 mark", 100.5, true},
		{"Invalid far above 100", 250.0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateMark(tc.mark)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateMark(%.2f) error = %v, wantErr = %v", tc.mark, err, tc.wantErr)
			}
			if tc.wantErr && !errors.Is(err, ErrMarksOutOfRange) {
				t.Errorf("expected ErrMarksOutOfRange, got %v", err)
			}
		})
	}
}

func TestValidateSubject(t *testing.T) {
	tests := []struct {
		name    string
		subject model.Subject
		wantErr bool
	}{
		{"Valid Subject", model.Subject{Name: "Go Programming", Marks: 92}, false},
		{"Empty Subject Name", model.Subject{Name: "   ", Marks: 75}, true},
		{"Invalid Marks in Subject", model.Subject{Name: "Mathematics", Marks: 105}, true},
		{"Negative Marks in Subject", model.Subject{Name: "Physics", Marks: -10}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateSubject(tc.subject)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateSubject() error = %v, wantErr = %v", err, tc.wantErr)
			}
		})
	}
}

func TestValidateStudent(t *testing.T) {
	tests := []struct {
		name    string
		student model.Student
		wantErr bool
	}{
		{
			name: "Valid Student",
			student: model.Student{
				Name:  "Ezhilarasi K",
				RegNo: "24BCS025",
				Subjects: []model.Subject{
					{Name: "Data Structures", Marks: 90},
					{Name: "Go Programming", Marks: 95},
				},
			},
			wantErr: false,
		},
		{
			name: "Empty Name",
			student: model.Student{
				Name:  "",
				RegNo: "24BCS025",
				Subjects: []model.Subject{
					{Name: "Maths", Marks: 85},
				},
			},
			wantErr: true,
		},
		{
			name: "Empty RegNo",
			student: model.Student{
				Name:  "Ezhilarasi K",
				RegNo: "",
				Subjects: []model.Subject{
					{Name: "Maths", Marks: 85},
				},
			},
			wantErr: true,
		},
		{
			name: "No Subjects",
			student: model.Student{
				Name:     "Ezhilarasi K",
				RegNo:    "24BCS025",
				Subjects: []model.Subject{},
			},
			wantErr: true,
		},
		{
			name: "Invalid Mark in Subject list",
			student: model.Student{
				Name:  "Ezhilarasi K",
				RegNo: "24BCS025",
				Subjects: []model.Subject{
					{Name: "Maths", Marks: 110},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateStudent(tc.student)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateStudent() error = %v, wantErr = %v", err, tc.wantErr)
			}
		})
	}
}
