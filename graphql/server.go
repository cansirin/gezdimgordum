package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cansirin/gezdimgordum/graphql/graph"
	"github.com/cansirin/gezdimgordum/graphql/graph/generated"
	"github.com/cansirin/gezdimgordum/graphql/internal/backend"
	"github.com/cansirin/gezdimgordum/graphql/internal/db"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbClient, err := db.NewPostgreSQLConnection(db.PostgreSQLConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	})

	if err != nil {
		log.Fatal("error while creating a db connection pool", err)
	}

	err = models.AutoMigrate(dbClient)
	if err != nil {
		return
	}

	postgreSQLBackend := backend.NewPostgreSQLBackend(dbClient)

	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler

	router.Use(c)

	if err != nil {
		log.Fatalf("initializing clients: %s", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Backend: postgreSQLBackend}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
