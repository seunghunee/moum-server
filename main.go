package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seunghunee/moum-server/graph"
	"github.com/seunghunee/moum-server/graph/generated"
)

func main() {
	config := generated.Config{Resolvers: &graph.Resolver{}}
	schema := generated.NewExecutableSchema(config)
	server := handler.NewDefaultServer(schema)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", server)
	http.HandleFunc(pathPrefix, apiHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

const pathPrefix = "/api/v1/article/"
