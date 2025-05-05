package main

import (
	"editorial-backend/database"
	"editorial-backend/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://localhost:5173", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	database.ConnectToDatabase()

	r.Get("/api/clients/", handler.GetClients)
	r.Get("/api/clients/csv", handler.ExportClientsToCSV)
	r.Get("/api/authors/", handler.GetAuthors)
	r.Get("/api/authors/csv", handler.ExportAuthorsToCSV)
	r.Get("/api/employees/", handler.GetEmployees)
	r.Get("/api/employees/csv", handler.ExportEmployeesToCSV)
	r.Get("/api/loans/", handler.GetLoans)
	r.Get("/api/loans/csv", handler.ExportLoansToCSV)
	r.Get("/api/fines/", handler.GetFines)
	r.Get("/api/fines/csv", handler.ExportFinesToCSV)

	log.Println("API is ready")

	http.ListenAndServe(":8080", r)
}
