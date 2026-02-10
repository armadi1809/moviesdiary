package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/tmdb"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := openDb()
	if err != nil {
		log.Panicf("Unable to connect to the database. Shutting server down %v", err)
	}

	tmdbCleint := tmdb.NewTmdbClient(os.Getenv("TMDB_API_KEY"))
	r := routes(db, tmdbCleint)

	slog.Info("Server Starting on Port 3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("An error occurred %v", err)
	}
}

func openDb() (*db.Queries, error) {
	dsn := os.Getenv("DbConnString")
	if dsn == "" {
		// Local default; foreign keys enabled.
		dsn = "file:moviesdiary.db?_foreign_keys=on"
	}
	slog.Info("Using SQLite", "dsn", dsn)
	dbConn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	if _, err := dbConn.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return nil, err
	}
	if err := applySchema(dbConn); err != nil {
		return nil, err
	}
	if err = dbConn.Ping(); err != nil {
		return nil, err
	}
	return db.New(dbConn), nil
}

func applySchema(dbConn *sql.DB) error {
	f, err := os.Open("sqlc/schema.sql")
	if err != nil {
		return err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	statements := strings.Split(string(content), ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := dbConn.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}
