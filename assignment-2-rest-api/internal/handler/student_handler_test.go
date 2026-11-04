package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"assignment2/internal/model"
	"assignment2/internal/store"
)

func setupTestServer() (*http.ServeMux, *store.MemoryStore) {
	st := store.NewMemoryStore()
	st.Reset()
	h := NewStudentHandler(st)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	return mux, st
}

func TestRootAndHealth(t *testing.T) {
	mux, _ := setupTestServer()

	// Root
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GET / returned status %d; want 200", rr.Code)
	}

	// Health
	reqHealth := httptest.NewRequest(http.MethodGet, "/health", nil)
	rrHealth := httptest.NewRecorder()
	mux.ServeHTTP(rrHealth, reqHealth)

	if rrHealth.Code != http.StatusOK {
		t.Errorf("GET /health returned status %d; want 200", rrHealth.Code)
	}
}

func TestGetAllStudents(t *testing.T) {
	mux, st := setupTestServer()

	// Insert one student
	_, _ = st.Create(model.Student{
		ID:         "24BCS025",
		Name:       "Ezhilarasi K",
		Department: "CSE",
		Year:       2,
		Email:      "ezhilarasi@ciet.edu.in",
	})

	req := httptest.NewRequest(http.MethodGet, "/api/students", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /api/students status = %d; want 200", rr.Code)
	}

	var resp model.APIResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	if !resp.Success {
		t.Errorf("Expected success to be true")
	}
}

func TestCreateStudent(t *testing.T) {
	mux, _ := setupTestServer()

	t.Run("Valid student creation", func(t *testing.T) {
		body := `{
			"id": "24BCS025",
			"name": "Ezhilarasi K",
			"department": "Computer Science and Engineering",
			"year": 2,
			"email": "ezhilarasi.24bcs025@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf("POST /api/students status = %d; want 201. Body: %s", rr.Code, rr.Body.String())
		}
	})

	t.Run("Duplicate ID returns 400", func(t *testing.T) {
		body := `{
			"id": "24BCS025",
			"name": "Duplicate Student",
			"department": "CSE",
			"year": 2,
			"email": "duplicate@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Duplicate ID status = %d; want 400", rr.Code)
		}
	})

	t.Run("Invalid year returns 400", func(t *testing.T) {
		body := `{
			"name": "Aarav Sharma",
			"department": "CSE",
			"year": 6,
			"email": "aarav@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Invalid year status = %d; want 400", rr.Code)
		}
	})

	t.Run("Invalid email returns 400", func(t *testing.T) {
		body := `{
			"name": "Aarav Sharma",
			"department": "CSE",
			"year": 2,
			"email": "invalid-email"
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Invalid email status = %d; want 400", rr.Code)
		}
	})

	t.Run("Missing name returns 400", func(t *testing.T) {
		body := `{
			"department": "CSE",
			"year": 2,
			"email": "test@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Missing name status = %d; want 400", rr.Code)
		}
	})

	t.Run("Malformed JSON returns 400", func(t *testing.T) {
		body := `{"name": "Bad JSON, missing close brace"`

		req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Malformed JSON status = %d; want 400", rr.Code)
		}
	})
}

func TestGetStudentByID(t *testing.T) {
	mux, st := setupTestServer()

	_, _ = st.Create(model.Student{
		ID:         "24BCS025",
		Name:       "Ezhilarasi K",
		Department: "CSE",
		Year:       2,
		Email:      "ezhilarasi@ciet.edu.in",
	})

	t.Run("Found student", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/students/24BCS025", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("GET /api/students/24BCS025 status = %d; want 200", rr.Code)
		}
	})

	t.Run("Not found student returns 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/students/NONEXISTENT", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("GET /api/students/NONEXISTENT status = %d; want 404", rr.Code)
		}
	})
}

func TestUpdateStudent(t *testing.T) {
	mux, st := setupTestServer()

	_, _ = st.Create(model.Student{
		ID:         "24BCS025",
		Name:       "Ezhilarasi K",
		Department: "CSE",
		Year:       2,
		Email:      "ezhilarasi@ciet.edu.in",
	})

	t.Run("Successful update", func(t *testing.T) {
		body := `{
			"name": "Ezhilarasi K (Updated)",
			"department": "Computer Science and Engineering",
			"year": 3,
			"email": "ezhilarasi.updated@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPut, "/api/students/24BCS025", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("PUT /api/students/24BCS025 status = %d; want 200", rr.Code)
		}
	})

	t.Run("Update not found student returns 404", func(t *testing.T) {
		body := `{
			"name": "Ghost",
			"department": "CSE",
			"year": 1,
			"email": "ghost@ciet.edu.in"
		}`

		req := httptest.NewRequest(http.MethodPut, "/api/students/UNKNOWN_ID", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("PUT /api/students/UNKNOWN_ID status = %d; want 404", rr.Code)
		}
	})

	t.Run("Update with invalid body returns 400", func(t *testing.T) {
		body := `{"name": ""}`

		req := httptest.NewRequest(http.MethodPut, "/api/students/24BCS025", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("PUT invalid body status = %d; want 400", rr.Code)
		}
	})
}

func TestDeleteStudent(t *testing.T) {
	mux, st := setupTestServer()

	_, _ = st.Create(model.Student{
		ID:         "24BCS025",
		Name:       "Ezhilarasi K",
		Department: "CSE",
		Year:       2,
		Email:      "ezhilarasi@ciet.edu.in",
	})

	t.Run("Successful delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/students/24BCS025", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("DELETE /api/students/24BCS025 status = %d; want 200", rr.Code)
		}

		// Ensure it is deleted
		reqGet := httptest.NewRequest(http.MethodGet, "/api/students/24BCS025", nil)
		rrGet := httptest.NewRecorder()
		mux.ServeHTTP(rrGet, reqGet)
		if rrGet.Code != http.StatusNotFound {
			t.Errorf("Deleted student still returned 200 OK")
		}
	})

	t.Run("Delete non-existent student returns 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/students/UNKNOWN_ID", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("DELETE non-existent status = %d; want 404", rr.Code)
		}
	})
}
