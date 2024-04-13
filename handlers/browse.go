package handlers

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/tmdb"
	"github.com/armadi1809/moviesdiary/views"
)

func BrowseHandler(tmdbClient *tmdb.TmdbClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		movies, err := tmdbClient.GetNowPlayingMovies()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		component := views.BrowsePage(user.Name != "", movies)
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}

func AddMovieModalHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posterUrl := r.URL.Query().Get("posterUrl")
		movieName := r.URL.Query().Get("movieName")
		description := r.URL.Query().Get("description")
		component := views.AddModalMovie(movieName, posterUrl, description)
		err := component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func SearchForMovieHandler(tmdbClient *tmdb.TmdbClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		searchQuery := r.Form.Get("query")
		var movies []tmdb.TmdbMovie
		if searchQuery == "" {
			movies, err = tmdbClient.GetNowPlayingMovies()
		} else {
			movies, err = tmdbClient.GetMovies(searchQuery)
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		component := views.MoviesListContainer(movies)
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
