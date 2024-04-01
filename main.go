package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sbClient := newSupabaseClient()
	r := routes(sbClient)

	slog.Info("Server Starting on Port 3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("An error occurred %v", err)
	}
}

func newSupabaseClient() *supabase.Client {
	sbHost := os.Getenv("SUPABASE_URL")
	sbSecret := os.Getenv("SUPABASE_SECRET")
	return supabase.CreateClient(sbHost, sbSecret)
}
