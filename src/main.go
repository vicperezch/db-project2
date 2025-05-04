package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	r := chi.NewRouter()

	connectToDatabase()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}

func connectToDatabase() {
	dsn := "host=db user=postgres password=postgres dbname=editorial port=5432 sslmode=disable TimeZone=UTC"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Connection to database failed.")
	}
}
