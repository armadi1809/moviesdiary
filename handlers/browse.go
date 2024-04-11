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
