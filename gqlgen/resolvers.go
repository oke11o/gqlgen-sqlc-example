package gqlgen

import (
	"context"
	"fmt"
	"github.com/oke11o/gqlgen-sqlc-example/pg"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Repository pg.Repository
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*Agent, error) {
	agent, err := r.Repository.CreateAgent(ctx, pg.CreateAgentParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &Agent{
		ID:    agent.ID,
		Name:  agent.Name,
		Email: agent.Email,
	}, nil
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
	a, err := r.Repository.ListAgents(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't find agents; err: %w", err)
	}
	res := make([]Agent, len(a))
	for i, agent := range a {
		res[i] = Agent{
			ID:    agent.ID,
			Name:  agent.Name,
			Email: agent.Email,
		}
	}
	return res, err
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
