// This package is a demonstration how to build and use a GraphQL server in Go
package main

import (
	"grapql/gopher"
	"grapql/job"
	"grapql/schemas"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Create the Gopher Repository
	gopherService := gopher.NewService(gopher.NewMemoryRepository(), job.NewMemoryRepository())
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
