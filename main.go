package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nedpals/supabase-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sbClient := newSupabaseClient()
	db, err := openDb()
	if err != nil {
		log.Panicf("Unable to connect to the database. Shutting server down %v", err)
	}
	r := routes(sbClient, db)

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

func openDb() (*sql.DB, error) {
	connString := os.Getenv("DbConnString")
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
