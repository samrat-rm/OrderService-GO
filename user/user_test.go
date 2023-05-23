package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUsers(t *testing.T) {
	// Create a new mock database
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	// Open a new GORM database connection with the mock DB
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	DB = db

	// Set up a test user
	user := User{
		Model:     gorm.Model{ID: 1},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	// Mock the DB Find method to return the test user
	mock.ExpectQuery("SELECT (.+) FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "email"}).
			AddRow(user.ID, user.FirstName, user.LastName, user.Email))

	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")

	// Define the test request
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	// Call the handler function
	router.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)

	// var response []User
	// json.Unmarshal(rec.Body.Bytes(), &response)
	// assert.Equal(t, []User{user}, response)
}

func TestCreateUser(t *testing.T) {
	// Create a new mock database
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	// Open a new GORM database connection with the mock DB
	DB, _ = gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	// Set up a test user
	user := User{
		Model:     gorm.Model{ID: 1},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	// Mock the DB Create method to return the test user
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")

	// Define the test request
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	rec := httptest.NewRecorder()

	// Call the handler function
	router.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response User
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, user, response)
}
