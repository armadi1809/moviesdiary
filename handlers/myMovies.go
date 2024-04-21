package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/views"
)

func MyMoviesHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		if user.ID == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		myMovies, err := queries.GetMoviesForUser(r.Context(), user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		component := views.MyMoviesPage(user.Name != "", myMovies)
		err = component.Render(r.Context(), w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func SearchMyMovies(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		if user.ID == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		myMovies, err := queries.GetMoviesForUser(r.Context(), user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		searchString := r.Form.Get("query")
		filteredMovies := []db.Movie{}
		if searchString == "" {
			filteredMovies = myMovies
		} else {
			for _, movie := range myMovies {
				if !strings.Contains(strings.ToLower(movie.Name), strings.ToLower(searchString)) {
					continue
				}
				filteredMovies = append(filteredMovies, movie)
			}
		}

		component := views.MyMoviesListContainer(filteredMovies)
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func EditMovieHandler(queries *db.Queries) http.HandlerFunc {
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
		dateWatched, err := time.Parse(timeLayout, r.Form.Get("dateWatched"))
		if err != nil {
			w.Header().Set("HX-Reswap", "none")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Watched Date is invalid"))
			return
		}
		movieId, err := strconv.Atoi(r.URL.Query().Get("movieId"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Invalid Movie Id"))
			return
		}
		params := db.EditMovieParams{
			ID:              int64(movieId),
			Diary:           r.Form.Get("diary"),
			LocationWatched: r.Form.Get("locationWatched"),
			WatchedDate:     dateWatched,
		}

		movie, err := queries.EditMovie(r.Context(), params)

		if err != nil {
			w.Header().Set("HX-Reswap", "none")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An Error ocurred when adding movie to db"))
			return
		}
		w.Header().Set("HX-Trigger", "editMovie")
		component := views.SuccessfullEditMessage(movie.Name)
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func EditMovieModalHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posterUrl := r.URL.Query().Get("posterUrl")
		movieName := r.URL.Query().Get("movieName")
		description := r.URL.Query().Get("description")
		locationWatched := r.URL.Query().Get("locationWatched")
		dateWatched := r.URL.Query().Get("dateWatched")
		diary := r.URL.Query().Get("diary")
		movieId := r.URL.Query().Get("movieId")
		component := views.EditModalMovie(movieName, posterUrl, description, locationWatched, dateWatched, diary, movieId)
		err := component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
func DeleteMovieHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieId, err := strconv.Atoi(r.URL.Query().Get("movieId"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Invalid Movie Id"))
			return
		}
		err = queries.DeleteMovie(r.Context(), int64(movieId))
		if err != nil {
			w.Header().Set("HX-Reswap", "none")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An Error ocurred when deleting movie from db"))
			return
		}

		w.Header().Set("HX-Trigger", "editMovie")
		component := views.SuccessfullDeleteMessage()
		err = component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func DeleteMovieModalHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieName := r.URL.Query().Get("movieName")
		movieId := r.URL.Query().Get("movieId")
		component := views.DeleteMovieModal(movieName, movieId)
		err := component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
