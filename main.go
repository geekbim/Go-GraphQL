package main

import (
	"log"
	"net/http"

	"github.com/geekbim/Go-GraphQL/gopher"
	"github.com/geekbim/Go-GraphQL/job"
	"github.com/geekbim/Go-GraphQL/schemas"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	gopherService := gopher.NewService(
		gopher.NewMemoryRepository(),
		job.NewMemoryRepository(),
	)

	schema, err := schemas.GenerateSchema(&gopherService)
	if err != nil {
		panic(err)
	}

	StartServer(schema)
}

// StartServer will trigger the server with a Playground
func StartServer(schema *graphql.Schema) {
	// Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema: schema,
		// Pretty print JSON response
		Pretty: true,
		// Host a GraphiQL Playground to use for testing Queries
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
