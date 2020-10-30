package auth

import (
	"context"
	"graphServer/jwt"
	"graphServer/users"
	"net/http"
	"strconv"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

//Middleware is ..
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			//Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}
			user := users.User{Username: username}
			id, err := users.GetUserIDByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)

			//Put in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//ForContext is ..
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
