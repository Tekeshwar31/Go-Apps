package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SupabaseClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

type apiResponse struct {
	Data   json.RawMessage `json:"data"`
	Status string          `json:"status"`
}

func NewSupabaseClient(baseURL, apiKey string) *SupabaseClient {
	return &SupabaseClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func (c *SupabaseClient) Request(method, path string, body interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("supabase error: %s", string(respBody))
	}

	return respBody, nil
}

// CORS middleware wrapper
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // or "*" for all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize Supabase client
	client := NewSupabaseClient(
		"https://bkuaxwxcwolujjwzhfvb.supabase.co/rest/v1/",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJrdWF4d3hjd29sdWpqd3poZnZiIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDgzMjY5MDksImV4cCI6MjA2MzkwMjkwOX0.WisnF6wJQBLvEaZvRKtZqtxIoOhZk60eO_-3QB2Iw3c",
	)
	// Wrap default mux with CORS middleware
	handlerWithCORS := corsMiddleware(http.DefaultServeMux)

	// Create HTTP server
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// query := "users?email=eq.west@example.com"
		query := "users?select*"
		// Get users from Supabase
		resp, err := client.Request("GET", query, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
			return
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Create a response struct
		response := apiResponse{
			Status: "success",
			Data:   resp,
		}

		// Marshal the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error formatting response: %v", err), http.StatusInternalServerError)
			return
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Send the formatted response
		w.Write(jsonResponse)
	})

	// Start the server
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
