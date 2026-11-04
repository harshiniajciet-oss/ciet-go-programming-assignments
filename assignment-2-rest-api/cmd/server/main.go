package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"assignment2/internal/handler"
	"assignment2/internal/middleware"
	"assignment2/internal/store"
)

func main() {
	portFlag := flag.String("port", "8080", "Port number on which HTTP server listens")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = *portFlag
	}

	printBanner(port)

	// Initialize In-Memory Store
	studentStore := store.NewMemoryStore()

	// Initialize Handlers and ServeMux
	mux := http.NewServeMux()
	studentHandler := handler.NewStudentHandler(studentStore)
	studentHandler.RegisterRoutes(mux)

	// Apply Middleware Pipeline (Logging)
	loggedMux := middleware.RequestLogger(os.Stdout)(mux)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggedMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server runner channel
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("[SERVER] Listening on http://localhost:%s\n", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Graceful Shutdown Channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] Could not start server: %v", err)
		}
	case sig := <-shutdown:
		log.Printf("[SERVER] Caught signal %v. Initiating graceful shutdown...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("[ERROR] Graceful shutdown failed: %v", err)
			_ = server.Close()
		}
		log.Println("[SERVER] Server exited cleanly.")
	}
}

func printBanner(port string) {
	line := strings.Repeat("=", 80)
	fmt.Println(line)
	fmt.Println("   Coimbatore Institute of Engineering and Technology (CIET)")
	fmt.Println("   ASSIGNMENT 2: Design and Develop a RESTful Web Service using Go.")
	fmt.Println("   Student Name: Ezhilarasi K   |   Register Number: 24BCS025")
	fmt.Println(line)
	fmt.Printf("   -> REST API Base URL: http://localhost:%s\n", port)
	fmt.Println("   -> Endpoints Available:")
	fmt.Println("      GET    /api/students       - Retrieve all students")
	fmt.Println("      POST   /api/students       - Create a new student")
	fmt.Println("      GET    /api/students/{id}  - Retrieve student by ID")
	fmt.Println("      PUT    /api/students/{id}  - Update student by ID")
	fmt.Println("      DELETE /api/students/{id}  - Delete student by ID")
	fmt.Println("      GET    /health             - Health check")
	fmt.Println("   * Note: In-memory store active. Data resets when the server restarts.")
	fmt.Println(line)
}
