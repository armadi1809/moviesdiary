package handlers

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/views"
)

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		views.CallbackScript().Render(r.Context(), w)
		return
	}
	cookie := http.Cookie{
		Name:     "at",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
