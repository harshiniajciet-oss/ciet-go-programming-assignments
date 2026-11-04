package middleware

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// responseWriterInterceptor wraps http.ResponseWriter to capture the HTTP status code.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterInterceptor) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// RequestLogger creates an HTTP middleware that logs the method, path, remote addr, status code, and latency.
func RequestLogger(out io.Writer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			interceptor := &responseWriterInterceptor{
				ResponseWriter: w,
				statusCode:     http.StatusOK, // Default to 200 unless WriteHeader is called
			}

			next.ServeHTTP(interceptor, r)

			duration := time.Since(startTime)

			fmt.Fprintf(out, "[REST-API] %s | %3d | %10v | %-7s %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				interceptor.statusCode,
				duration,
				r.Method,
				r.URL.Path,
			)
		})
	}
}
