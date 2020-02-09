// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

type Agent struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Authors []Author `json:"authors"`
}

type AgentInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Author struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Website *string `json:"website"`
	Agent   *Agent  `json:"agent"`
	Books   []Book  `json:"books"`
}

type AuthorInput struct {
	Name    string  `json:"name"`
	Website *string `json:"website"`
	AgentID int64   `json:"agent_id"`
}

type Book struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	Authors     []Author `json:"authors"`
}

type BookInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cover       string  `json:"cover"`
	AuthorIDs   []int64 `json:"authorIDs"`
}