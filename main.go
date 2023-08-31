package main

import (
	"appointments-api/handlers"
	"appointments-api/middlewares"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=joaoluismoraes dbname=appointments password=999selva sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// MUX Router
	router := mux.NewRouter()
	router.HandleFunc("/patients", handlers.GetPatients(db)).Methods("GET")
	// router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/patients", handlers.CreatePatient(db)).Methods("POST")
	// router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", middlewares.JsonContentTypeMiddleware(router)))
}
