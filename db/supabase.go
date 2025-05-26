package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	supabase "github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func InitSupabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	Client, err = supabase.NewClient(supabaseUrl, supabaseKey, nil)
	if err != nil {
		log.Fatalf("Cannot initialize supabase client: %v", err)
	}

	// Test connection
	_, err = Client.From("users").Select("*", "1", false).Execute()
	if err != nil {
		log.Fatalf("Supabase connection test failed: %v", err)
	}

	log.Println("Connected to Supabase successfully!")
}