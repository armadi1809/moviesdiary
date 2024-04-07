package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nedpals/supabase-go"
)

type userInfoKey string

func WithAuth(sbClient *supabase.Client) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("at")
			fmt.Println("Heeeeeeereeee", cookie)

			if err != nil {
				http.Redirect(w, r, "/login/google", http.StatusPermanentRedirect)
				return
			}
			accessToken := cookie.Value
			user, err := sbClient.Auth.User(r.Context(), accessToken)
			if err != nil {
				http.Redirect(w, r, "/login/google", http.StatusPermanentRedirect)
				return
			}
			ctx := context.WithValue(r.Context(), userInfoKey("userInfo"), *user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}

}
