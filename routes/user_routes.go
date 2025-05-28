package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go-supabase-api/supabase"
)

type apiResponse struct {
	Data   json.RawMessage `json:"data"`
	Status string          `json:"status"`
}

func SetupUserRoutes(client *supabase.SupabaseClient) http.Handler {
	mux := http.NewServeMux()

	// Get all users
	mux.HandleFunc("/get-users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query := "users?select*"
		resp, err := client.Request("GET", query, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
			return
		}

		response := apiResponse{
			Status: "success",
			Data:   resp,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error formatting response: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	// Add a new user
	mux.HandleFunc("/add-user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := apiResponse{
			Status: "success",
			Data:   json.RawMessage(`{"message": "User added successfully"}`),
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error formatting response: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	return mux
}