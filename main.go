package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/armadi1809/moviesdiary/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		component := views.Hello("Movie Diary")
		err := component.Render(r.Context(), w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	slog.Info("Server Starting on Port 3000...")
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Printf("An error occurred %v", err)
	}

}
