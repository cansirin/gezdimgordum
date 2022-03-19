package auth

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/backend"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"github.com/cansirin/gezdimgordum/graphql/pkg/jwt"
	"github.com/gofrs/uuid"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(backend backend.Backender) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := models.User{Username: username}
			id, err := backend.GetUserID(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = uuid.FromStringOrNil(id)
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}
