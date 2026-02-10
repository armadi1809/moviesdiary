package main

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/handlers"
	"github.com/armadi1809/moviesdiary/tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(db *db.Queries, tmdbClient *tmdb.TmdbClient) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.WithLocalUser(db))
	r.Get("/browse", handlers.BrowseHandler(tmdbClient))
	r.Get("/addMovieModal", handlers.AddMovieModalHandler())
	r.Get("/editMovieModal", handlers.EditMovieModalHandler())
	r.Post("/searchMovie", handlers.SearchForMovieHandler(tmdbClient))
	r.Post("/addMovie", handlers.AddMovieHandler(db))
	r.Post("/editMovie", handlers.EditMovieHandler(db))
	r.Post("/searchMyMovies", handlers.SearchMyMovies(db))
	r.Get("/myMovies", handlers.MyMoviesHandler(db))
	r.Post("/deleteMovie", handlers.DeleteMovieHandler(db))
	r.Get("/deleteMovieModal", handlers.DeleteMovieModalHandler())
	r.Get("/", handlers.HomeHandler())
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	return r
}
