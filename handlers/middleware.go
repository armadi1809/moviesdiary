package handlers

import (
	"context"
	"net/http"

	"github.com/armadi1809/moviesdiary/db"
)

type userInfoKey string

func WithLocalUser(queries *db.Queries) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Single-user local mode (no authentication).
			const email = "local@localhost"
			const name = "Local"

			userDb, err := queries.GetUser(r.Context(), email)
			if err != nil {
				userDb, err = queries.CreateUser(r.Context(), db.CreateUserParams{Email: email, Name: name})
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Unable to setup local user"))
					return
				}
			}

			ctx := context.WithValue(r.Context(), userInfoKey("userInfo"), userDb)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func getUserFromRequest(r *http.Request) db.User {
	user, ok := r.Context().Value(userInfoKey("userInfo")).(db.User)
	if !ok {
		return db.User{}
	}
	return user
}
