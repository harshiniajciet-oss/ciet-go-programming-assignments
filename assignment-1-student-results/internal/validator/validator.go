package validator

import (
	"errors"
	"fmt"
	"strings"

	"assignment1/internal/model"
)

var (
	ErrEmptyName       = errors.New("student name cannot be empty")
	ErrEmptyRegNo      = errors.New("register number cannot be empty")
	ErrNoSubjects      = errors.New("student must have at least one subject")
	ErrEmptySubName    = errors.New("subject name cannot be empty")
	ErrMarksOutOfRange = errors.New("marks must be between 0 and 100 inclusive")
)

// ValidateMark checks if a mark falls within the valid range [0, 100].
func ValidateMark(mark float64) error {
	if mark < 0 || mark > 100 {
		return fmt.Errorf("%w: received %.2f", ErrMarksOutOfRange, mark)
	}
	return nil
}

// ValidateSubject validates individual subject details (name and marks).
func ValidateSubject(sub model.Subject) error {
	if strings.TrimSpace(sub.Name) == "" {
		return ErrEmptySubName
	}
	if err := ValidateMark(sub.Marks); err != nil {
		return fmt.Errorf("subject '%s': %w", sub.Name, err)
	}
	return nil
}

// ValidateStudent performs complete validation on a student's input fields.
func ValidateStudent(s model.Student) error {
	if strings.TrimSpace(s.Name) == "" {
		return ErrEmptyName
	}
	if strings.TrimSpace(s.RegNo) == "" {
		return ErrEmptyRegNo
	}
	if len(s.Subjects) == 0 {
		return ErrNoSubjects
	}
	for i, sub := range s.Subjects {
		if err := ValidateSubject(sub); err != nil {
			return fmt.Errorf("subject #%d invalid: %w", i+1, err)
		}
	}
	return nil
}
