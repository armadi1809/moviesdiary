package handlers

import (
	"net/http"

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
