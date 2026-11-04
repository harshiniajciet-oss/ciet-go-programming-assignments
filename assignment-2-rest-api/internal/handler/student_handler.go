package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"assignment2/internal/model"
	"assignment2/internal/store"
)

// StudentHandler handles HTTP requests for the Student resource.
type StudentHandler struct {
	store store.Store
}

// NewStudentHandler creates a new StudentHandler backed by the provided store.
func NewStudentHandler(s store.Store) *StudentHandler {
	return &StudentHandler{store: s}
}

// RegisterRoutes registers student REST endpoints onto the provided http.ServeMux.
// It supports Go 1.22+ pattern routing while remaining backward-compatible.
func (h *StudentHandler) RegisterRoutes(mux *http.ServeMux) {
	// Root and Health Endpoints
	mux.HandleFunc("GET /", h.RootHandler)
	mux.HandleFunc("GET /health", h.HealthCheck)

	// Collection Route: /api/students
	mux.HandleFunc("GET /api/students", h.GetAllStudents)
	mux.HandleFunc("POST /api/students", h.CreateStudent)

	// Item Route: /api/students/{id}
	mux.HandleFunc("GET /api/students/{id}", h.GetStudentByID)
	mux.HandleFunc("PUT /api/students/{id}", h.UpdateStudent)
	mux.HandleFunc("DELETE /api/students/{id}", h.DeleteStudent)

	// Also support non-prefixed routes for convenience: /students and /students/{id}
	mux.HandleFunc("GET /students", h.GetAllStudents)
	mux.HandleFunc("POST /students", h.CreateStudent)
	mux.HandleFunc("GET /students/{id}", h.GetStudentByID)
	mux.HandleFunc("PUT /students/{id}", h.UpdateStudent)
	mux.HandleFunc("DELETE /students/{id}", h.DeleteStudent)
}

// RootHandler provides API metadata and student creator details.
func (h *StudentHandler) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.writeError(w, http.StatusNotFound, "Endpoint not found")
		return
	}
	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Welcome to CIET Student REST API",
		Data: map[string]interface{}{
			"assignment":  "ASSIGNMENT 2: Design and Develop a RESTful Web Service using Go",
			"author":      "Ezhilarasi K",
			"reg_no":      "24BCS025",
			"institution": "Coimbatore Institute of Engineering and Technology (CIET)",
			"endpoints": []string{
				"GET    /health",
				"GET    /api/students",
				"POST   /api/students",
				"GET    /api/students/{id}",
				"PUT    /api/students/{id}",
				"DELETE /api/students/{id}",
			},
		},
	})
}

// HealthCheck returns a 200 OK status verifying server health.
func (h *StudentHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Service is healthy and operational",
	})
}

// GetAllStudents handles GET /api/students and returns all student records.
func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := h.store.GetAll()
	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Retrieved %d student record(s)", len(students)),
		Data:    students,
	})
}

// GetStudentByID handles GET /api/students/{id} and retrieves a student by identifier.
func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := h.extractID(r)
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "Missing student ID in request URL")
		return
	}

	student, found := h.store.GetByID(id)
	if !found {
		h.writeError(w, http.StatusNotFound, fmt.Sprintf("Student with ID '%s' not found", id))
		return
	}

	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    student,
	})
}

// CreateStudent handles POST /api/students and adds a new student.
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID string `json:"id"`
		model.StudentInput
	}

	if err := h.decodeJSON(r.Body, &payload); err != nil {
		h.writeError(w, http.StatusBadRequest, fmt.Sprintf("Malformed or invalid JSON payload: %v", err))
		return
	}

	if err := payload.StudentInput.Validate(); err != nil {
		h.writeError(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err))
		return
	}

	student := model.Student{
		ID:         strings.TrimSpace(payload.ID),
		Name:       payload.Name,
		Department: payload.Department,
		Year:       payload.Year,
		Email:      payload.Email,
	}

	created, err := h.store.Create(student)
	if err != nil {
		if errors.Is(err, store.ErrDuplicateID) {
			h.writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		h.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create student: %v", err))
		return
	}

	h.writeJSON(w, http.StatusCreated, model.APIResponse{
		Success: true,
		Message: "Student created successfully",
		Data:    created,
	})
}

// UpdateStudent handles PUT /api/students/{id} and modifies an existing student.
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := h.extractID(r)
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "Missing student ID in request URL")
		return
	}

	var input model.StudentInput
	if err := h.decodeJSON(r.Body, &input); err != nil {
		h.writeError(w, http.StatusBadRequest, fmt.Sprintf("Malformed or invalid JSON payload: %v", err))
		return
	}

	if err := input.Validate(); err != nil {
		h.writeError(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err))
		return
	}

	updated, found := h.store.Update(id, input)
	if !found {
		h.writeError(w, http.StatusNotFound, fmt.Sprintf("Cannot update: Student with ID '%s' not found", id))
		return
	}

	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Student updated successfully",
		Data:    updated,
	})
}

// DeleteStudent handles DELETE /api/students/{id} and removes a student record.
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := h.extractID(r)
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "Missing student ID in request URL")
		return
	}

	deleted := h.store.Delete(id)
	if !deleted {
		h.writeError(w, http.StatusNotFound, fmt.Sprintf("Cannot delete: Student with ID '%s' not found", id))
		return
	}

	h.writeJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Student with ID '%s' deleted successfully", id),
	})
}

// extractID extracts the path parameter 'id' from the request using Go 1.22 PathValue or URL parsing.
func (h *StudentHandler) extractID(r *http.Request) string {
	if val := r.PathValue("id"); val != "" {
		return val
	}

	// Fallback extraction
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

// decodeJSON safely decodes request body JSON and checks for EOF / empty body errors.
func (h *StudentHandler) decodeJSON(r io.Reader, v interface{}) error {
	if r == nil {
		return errors.New("request body is nil")
	}
	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}

// writeJSON serializes the payload to JSON with appropriate Content-Type header.
func (h *StudentHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// writeError responds with a structured error JSON object.
func (h *StudentHandler) writeError(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, model.APIResponse{
		Success: false,
		Error:   message,
	})
}
