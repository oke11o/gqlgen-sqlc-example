package dataloaders

import (
	"context"
	"net/http"

	"github.com/oke11o/gqlgen-sqlc-example/pg"
)

// Middleware stores Loaders as a request-scoped context value.
func Middleware(repo pg.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(ctx, repo)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}
