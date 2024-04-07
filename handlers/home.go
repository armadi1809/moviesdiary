package handlers

import (
	"fmt"
	"net/http"

	"github.com/armadi1809/moviesdiary/views"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		fmt.Println("Heeeeereeee", user.Email)
		component := views.Hello(user.Email)
		err := component.Render(r.Context(), w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
