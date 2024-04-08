package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/armadi1809/moviesdiary/db"
	"github.com/nedpals/supabase-go"
)

type userInfoKey string

func WithAuth(sbClient *supabase.Client, queries *db.Queries) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("at")
			if err != nil {
				fmt.Printf("An error occurred getting the cookie in with auth %v", err)
				http.Redirect(w, r, "/login/google", http.StatusTemporaryRedirect)
				return
			}
			accessToken := cookie.Value
			user, err := sbClient.Auth.User(r.Context(), accessToken)
			if err != nil {
				fmt.Printf("Could not authenticate user with google %v", err)
				http.Redirect(w, r, "/login/google", http.StatusTemporaryRedirect)
				return
			}
			userDb, err := queries.GetUser(r.Context(), user.Email)
			if err != nil {
				userEmailSplit := strings.Split(user.Email, "@")
				userDb, err = queries.CreateUser(r.Context(), db.CreateUserParams{Name: userEmailSplit[0], Email: user.Email})
				if err != nil {
					fmt.Printf("An Error occurred creating user %v\n", err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Unabele to find or setup account for user"))
					return
				}
			}
			ctx := context.WithValue(r.Context(), userInfoKey("userInfo"), userDb)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}

}
