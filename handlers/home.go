package handlers

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/views"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		component := views.LandingPage(user.Name != "")
		err := component.Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
