package main

import (
	"appointments-api/handlers"
	"appointments-api/middlewares"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", "user="+dbUser+" password="+dbPassword+" dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/patients", handlers.GetPatients(db)).Methods("GET")
	router.HandleFunc("/patients", handlers.CreatePatient(db)).Methods("POST")
	router.HandleFunc("/doctors", handlers.GetDoctors(db)).Methods("GET")
	router.HandleFunc("/doctors", handlers.CreateDoctor(db)).Methods("POST")
	router.HandleFunc("/appointments", handlers.CreateWeeklyAppointments(db)).Methods("POST")
	router.HandleFunc("/appointments/{appointmentId}", handlers.ScheduleAppointment(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", middlewares.JsonContentTypeMiddleware(router)))
}
