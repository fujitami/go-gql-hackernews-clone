package main

import (
	"go-gql-hackernews-clone/graph"
	"go-gql-hackernews-clone/graph/generated"
	database "go-gql-hackernews-clone/internal/pkg/db/mysql"
	"go-gql-hackernews-clone/internal/auth"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "5002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
