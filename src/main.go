package main

import (
	"editorial-backend/database"
	"editorial-backend/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	database.ConnectToDatabase()

	r.Get("/api/clients/", handler.GetClients)

	http.ListenAndServe(":8080", r)
}
