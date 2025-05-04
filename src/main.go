package main

import (
	"editorial-backend/database"
	"editorial-backend/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	database.ConnectToDatabase()

	r.Get("/api/clients/", handler.GetClients)
	r.Get("/api/authors/", handler.GetAuthors)

	http.ListenAndServe(":8080", r)

	log.Println("API is ready")
}
