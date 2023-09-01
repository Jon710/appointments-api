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

func TestGetPatients(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	req, err := http.NewRequest("GET", "/patients", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/patients", handlers.GetPatients(db)).Methods("GET")

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "Alice", "alice@example.com").
		AddRow(2, "Bob", "bob@example.com")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM patients")).WillReturnRows(rows)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var patients []types.Patient
	err = json.NewDecoder(rr.Body).Decode(&patients)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []types.Patient{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}, patients)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreatePatient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO patients (name, email) VALUES ($1, $2) RETURNING id")).
		WithArgs("Sarah", "teste@teste.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	reqBody := types.Patient{Name: "Sarah", Email: "teste@teste.com"}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/patients", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/patients", handlers.CreatePatient(db)).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []byte(`{"id":1,"name":"Sarah","email":"teste@teste.com"}` + "\n")
	if !bytes.Equal(rr.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
