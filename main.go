package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/armadi1809/moviesdiary/sb"
	"github.com/armadi1809/moviesdiary/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sb.Init()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		component := views.Hello("Movie Diary")
		err := component.Render(r.Context(), w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	r.Get("/login/google", func(w http.ResponseWriter, r *http.Request) {
		resp, err := sb.Client.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
			Provider:   "google",
			RedirectTo: "http://localhost:3000/",
		})
		if err != nil {
			fmt.Println("Error ocurred")
		}
		http.Redirect(w, r, resp.URL, http.StatusSeeOther)
	})
	slog.Info("Server Starting on Port 3000...")
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	err = http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Printf("An error occurred %v", err)
	}

}
