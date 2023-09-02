package handlers_test

import (
	"appointments-api/handlers"
	"appointments-api/types"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetDoctors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	req, err := http.NewRequest("GET", "/doctors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/doctors", handlers.GetDoctors(db)).Methods("GET")

	rows := sqlmock.NewRows([]string{"id", "name", "specialty"}).
		AddRow(1, "Alice", "cardio").
		AddRow(2, "Bob", "teste")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM doctors")).WillReturnRows(rows)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var doctors []types.Doctor
	err = json.NewDecoder(rr.Body).Decode(&doctors)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []types.Doctor{
		{ID: 1, Name: "Alice", Specialty: "cardio"},
		{ID: 2, Name: "Bob", Specialty: "teste"},
	}, doctors)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO doctors (name, specialty) VALUES ($1, $2) RETURNING id")).
		WithArgs("Sarah", "cardio").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	reqBody := types.Doctor{Name: "Sarah", Specialty: "cardio"}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/doctors", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/doctors", handlers.CreateDoctor(db)).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []byte(`{"id":1,"name":"Sarah","specialty":"cardio"}` + "\n")
	if !bytes.Equal(rr.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
