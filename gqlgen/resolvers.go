package gqlgen

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*Agent, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAgent(ctx context.Context, id int64, data AgentInput) (*Agent, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAgent(ctx context.Context, id int64) (*Agent, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (*Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (*Book, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Agent(ctx context.Context, id int64) (*Agent, error) {
	panic("not implemented")
}
func (r *queryResolver) Agents(ctx context.Context) ([]Agent, error) {
	panic("not implemented")
}
func (r *queryResolver) Author(ctx context.Context, id int64) (*Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Authors(ctx context.Context) ([]Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Book(ctx context.Context, id int64) (*Book, error) {
	panic("not implemented")
}
func (r *queryResolver) Books(ctx context.Context) ([]Book, error) {
	panic("not implemented")
}
