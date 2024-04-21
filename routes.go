package main

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/handlers"
	"github.com/armadi1809/moviesdiary/tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nedpals/supabase-go"
)

func routes(sbClient *supabase.Client, db *db.Queries, tmdbClient *tmdb.TmdbClient) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.WithUser(sbClient, db))
	r.Group(func(authenticated chi.Router) {
		authenticated.Use(handlers.WithAuth())
		authenticated.Get("/browse", handlers.BrowseHandler(tmdbClient))
		authenticated.Get("/addMovieModal", handlers.AddMovieModalHandler())
		authenticated.Get("/editMovieModal", handlers.EditMovieModalHandler())
		authenticated.Post("/searchMovie", handlers.SearchForMovieHandler(tmdbClient))
		authenticated.Post("/addMovie", handlers.AddMovieHandler(db))
		authenticated.Post("/editMovie", handlers.EditMovieHandler(db))
		authenticated.Post("/searchMyMovies", handlers.SearchMyMovies(db))
		authenticated.Get("/myMovies", handlers.MyMoviesHandler(db))
		authenticated.Get("/login", handlers.LoginPageHandler())

	})
	r.Get("/", handlers.HomeHandler())
	r.Get("/logout", handlers.HandleLogout)
	r.Get("/login", handlers.LoginPageHandler())
	r.Get("/login/google", handlers.GoogleLoginHandler(sbClient))
	r.Get("/auth/callback", handlers.HandleAuthCallback)
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	return r
}
