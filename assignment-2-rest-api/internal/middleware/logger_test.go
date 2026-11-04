package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := RequestLogger(buf)

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("ok"))
	})

	wrapped := logger(dummyHandler)

	req := httptest.NewRequest(http.MethodPost, "/api/students", nil)
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("got status %d, want %d", rr.Code, http.StatusCreated)
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "POST") {
		t.Errorf("Log output missing HTTP method: %s", logOutput)
	}
	if !strings.Contains(logOutput, "/api/students") {
		t.Errorf("Log output missing URL path: %s", logOutput)
	}
	if !strings.Contains(logOutput, "201") {
		t.Errorf("Log output missing status code 201: %s", logOutput)
	}
}
