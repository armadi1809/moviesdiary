package handlers

import (
	"fmt"
	"net/http"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/armadi1809/moviesdiary/views"
	"github.com/nedpals/supabase-go"
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
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	fmt.Println("access_token", accessToken)
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/myMovies", http.StatusSeeOther)
}

func GoogleLoginHandler(sbClient *supabase.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := sbClient.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
			Provider:   "google",
			RedirectTo: "http://localhost:3000/auth/callback",
		})
		if err != nil {
			fmt.Println("Error ocurred")
		}
		http.Redirect(w, r, resp.URL, http.StatusSeeOther)
	}
}

func getUserFromRequest(r *http.Request) db.User {
	user, ok := r.Context().Value(userInfoKey("userInfo")).(db.User)
	if !ok {
		return db.User{}
	}
	return user
}
