package handlers

import (
	"net/http"
	"strings"

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
