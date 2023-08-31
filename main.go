package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Patient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := sql.Open("postgres", "user=joaoluismoraes dbname=appointments password=999selva sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// MUX Router
	router := mux.NewRouter()
	router.HandleFunc("/patients", getPatients(db)).Methods("GET")
	// router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/patients", createPatient(db)).Methods("POST")
	// router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getPatients(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM patients")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		patients := []Patient{}
		for rows.Next() {
			var p Patient
			if err := rows.Scan(&p.ID, &p.Name, &p.Email); err != nil {
				log.Fatal(err)
			}
			patients = append(patients, p)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(patients)
	}
}

func createPatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Patient
		json.NewDecoder(r.Body).Decode(&p)

		err := db.QueryRow("INSERT INTO patients (name, email) VALUES ($1, $2) RETURNING id", p.Name, p.Email).Scan(&p.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(p)
	}
}
