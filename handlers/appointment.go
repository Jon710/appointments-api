package handlers

import (
	"appointments-api/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateWeeklyAppointments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a types.Appointment
		json.NewDecoder(r.Body).Decode(&a)

		stmt := `WITH RECURSIVE appointment_cte AS (
			SELECT CAST('09:00' AS TIME) AS date, 'Available' AS status
			UNION ALL
			SELECT 
				date + INTERVAL '30 minutes',
				'Available' AS status
			FROM appointment_cte WHERE date < CAST('18:00' AS TIME)
		)
		
		INSERT INTO appointments (date, status, doctor_id)
		SELECT date, status, $1 FROM appointment_cte RETURNING id;`

		err := db.QueryRow(stmt, a.DoctorID).Scan(&a.ID)
		if err != nil {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(a)
	}
}

func ScheduleAppointment(db *sql.DB) http.HandlerFunc {
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
