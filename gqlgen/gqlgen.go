package gqlgen

import (
	"github.com/oke11o/gqlgen-sqlc-example/pg"
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(repo pg.Repository) http.Handler {
	return handler.GraphQL(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository: repo,
		},
	}))
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.Handler {
	return handler.Playground("GraphQL Playground", endpoint)
}