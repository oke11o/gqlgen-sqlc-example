package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/oke11o/gqlgen-sqlc-example/gqlgen"
	"github.com/oke11o/gqlgen-sqlc-example/pg"
)

type Config struct {
	Postgres struct {
		Password   string `default:"postgres"`
		User       string `default:"postgres"`
		Db         string `default:"gqlgen-example"`
		PortExpose int    `envconfig:"port_expose"`
	}
}

func main() {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		panic(fmt.Sprintf("init config: %w", err))
	}
	// initialize the db
	db, err := pg.Open(
		fmt.Sprintf("host=localhost port=%d dbname=%s user=%s password=%s  sslmode=disable",
			config.Postgres.PortExpose, config.Postgres.Db, config.Postgres.User, config.Postgres.Password))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
