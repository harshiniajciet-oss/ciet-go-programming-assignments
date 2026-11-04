package model

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

var (
	ErrEmptyID         = errors.New("student id cannot be empty")
	ErrEmptyName       = errors.New("student name is required and cannot be blank")
	ErrEmptyDepartment = errors.New("department is required and cannot be blank")
	ErrInvalidYear     = errors.New("academic year must be between 1 and 5")
	ErrInvalidEmail    = errors.New("invalid email address format")
)

// Student represents an academic student entity in the system.
type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Year       int    `json:"year"`
	Email      string `json:"email"`
}

// StudentInput represents the payload submitted for creating or updating a student.
type StudentInput struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Year       int    `json:"year"`
	Email      string `json:"email"`
}

// Validate ensures all required student fields conform to academic rules.
func (s *StudentInput) Validate() error {
	s.Name = strings.TrimSpace(s.Name)
	if s.Name == "" {
		return ErrEmptyName
	}

	s.Department = strings.TrimSpace(s.Department)
	if s.Department == "" {
		return ErrEmptyDepartment
	}

	if s.Year < 1 || s.Year > 5 {
		return fmt.Errorf("%w: received %d", ErrInvalidYear, s.Year)
	}

	s.Email = strings.TrimSpace(s.Email)
	if s.Email == "" {
		return ErrInvalidEmail
	}

	addr, err := mail.ParseAddress(s.Email)
	if err != nil || !strings.Contains(addr.Address, ".") {
		return fmt.Errorf("%w: '%s'", ErrInvalidEmail, s.Email)
	}

	return nil
}

// APIResponse represents a standard JSON API response structure.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
