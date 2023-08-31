package handlers

import (
	"appointments-api/types"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetPatients(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM patients")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		patients := []types.Patient{}
		for rows.Next() {
			var p types.Patient
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

func CreatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p types.Patient
		json.NewDecoder(r.Body).Decode(&p)

		err := db.QueryRow("INSERT INTO patients (name, email) VALUES ($1, $2) RETURNING id", p.Name, p.Email).Scan(&p.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(p)
	}
}
