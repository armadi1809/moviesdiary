package handlers

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/tmdb"
	"github.com/armadi1809/moviesdiary/views"
)

const timeLayout = "2006-01-02"

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

func AddMovieHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		if user.ID == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// dateWatched, err := time.Parse(timeLayout, r.Form.Get("dateWatched"))
		// if err != nil {
		// 	w.Header().Set("HX-Reswap", "none")
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("Watched Date is invalid"))
		// 	return
		// }
		params := db.CreateMovieParams{
			UserID:          user.ID,
			PosterUrl:       r.URL.Query().Get("posterUrl"),
			Name:            r.URL.Query().Get("movieName"),
			Description:     r.URL.Query().Get("description"),
			Diary:           r.Form.Get("diary"),
			LocationWatched: r.Form.Get("locationWatched"),
			//WatchedDate:     sql.NullTime{},
		}

		movie, err := queries.CreateMovie(r.Context(), params)

		if err != nil {
			w.Header().Set("HX-Reswap", "none")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An Error ocurred when adding movie to db"))
			return
		}

		component := views.SuccessfullAdditionMessage(movie.Name.String)
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
