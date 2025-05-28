package main

import (
	"fmt"
	"net/http"
	"go-supabase-api/routes"
	"go-supabase-api/supabase"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
		// Initialize Supabase client
	client := supabase.NewSupabaseClient(
		"https://bkuaxwxcwolujjwzhfvb.supabase.co/rest/v1/",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJrdWF4d3hjd29sdWpqd3poZnZiIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDgzMjY5MDksImV4cCI6MjA2MzkwMjkwOX0.WisnF6wJQBLvEaZvRKtZqtxIoOhZk60eO_-3QB2Iw3c",
	)

	// Setup routes
	handler := routes.SetupUserRoutes(client)

	// Apply CORS middleware
	handler = corsMiddleware(handler)

	// Start the server
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}