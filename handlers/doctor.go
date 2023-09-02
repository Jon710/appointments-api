package handlers

import (
	"appointments-api/types"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetDoctors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM doctors")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		doctors := []types.Doctor{}
		for rows.Next() {
			var d types.Doctor
			if err := rows.Scan(&d.ID, &d.Name, &d.Specialty); err != nil {
				log.Fatal(err)
			}
			doctors = append(doctors, d)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(doctors)
	}
}

func CreateDoctor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d types.Doctor
		json.NewDecoder(r.Body).Decode(&d)

		err := db.QueryRow("INSERT INTO doctors (name, specialty) VALUES ($1, $2) RETURNING id", d.Name, d.Specialty).Scan(&d.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(d)
	}
}
