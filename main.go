package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seunghunee/moum-server/accessor"
	"github.com/seunghunee/moum-server/graph"
	"github.com/seunghunee/moum-server/graph/generated"
)

func main() {
	config := generated.Config{Resolvers: &graph.Resolver{Accessor: m}}
	schema := generated.NewExecutableSchema(config)
	server := handler.NewDefaultServer(schema)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", server)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// FIXME: m is NOT thread-safe
var m = accessor.NewInMemoryAccessor()
