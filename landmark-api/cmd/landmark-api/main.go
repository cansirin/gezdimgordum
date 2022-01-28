package main

import (
	"github.com/cansirin/gezdimgordum/landmark-api/internal/backend"
	"github.com/cansirin/gezdimgordum/landmark-api/internal/db"
	"github.com/cansirin/gezdimgordum/landmark-api/internal/models"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/cansirin/gezdimgordum/landmark-api/server"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
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

	s := server.NewLandmarkAPIServer(postgreSQLBackend)

	twirpHandler := api.NewLandmarkAPIServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		setupHeader(rw, r)
		rw.Write([]byte("OK"))
	}))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	http.ListenAndServe(":80", handler)
}

func setupHeader(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
