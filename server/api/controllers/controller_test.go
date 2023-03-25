//Provides a Golang Test Environment
//Tests can be executed by using "go test"
//The following functions test the controller functions used alongside a test sqlite database

package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blakeshelley10/CampUs/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegisterUser(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "crandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err2 := encoder.Encode(testUser); err2 != nil {
		t.Fatal(err2)
	}
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/users", &b)

	//Call function
	RegisterUser(testDB, wr, req)

	//Check response code and body
	if wr.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status code 201, got %d", wr.Code)
	}

	resultUser := models.User{}
	decoder := json.NewDecoder(wr.Body)
	if err3 := decoder.Decode(&resultUser); err != nil {
		t.Fatal(err3)
	}
	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}

func TestLogIn(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "nrandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err2 := encoder.Encode(testUser); err2 != nil {
		t.Fatal(err2)
	}
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/users", &b)

	//Call function
	LogIn(testDB, wr, req)

	//Check response code and body
	if wr.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status code 201, got %d", wr.Code)
	}

	resultUser := models.User{}
	decoder := json.NewDecoder(wr.Body)
	if err3 := decoder.Decode(&resultUser); err != nil {
		t.Fatal(err3)
	}
	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}

func TestCreateUser(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "mrandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err2 := encoder.Encode(testUser); err2 != nil {
		t.Fatal(err2)
	}
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/users", &b)

	//Call function
	CreateUser(testDB, wr, req)

	//Check response code and body
	if wr.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status code 201, got %d", wr.Code)
	}

	resultUser := models.User{}
	decoder := json.NewDecoder(wr.Body)
	if err3 := decoder.Decode(&resultUser); err != nil {
		t.Fatal(err3)
	}
	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}

func TestGetUser(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "mrandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users/mrandall24", nil)

	//Call function
	resultUser := findUser(testDB, testUser.Username, wr, req)

	//Check response code and body
	if wr.Code == http.StatusNotFound {
		t.Errorf("User not found")
	}

	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}

func TestUpdateUser(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "mrandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users/mrandall24", nil)

	//Call function
	resultUser := findUser(testDB, testUser.Username, wr, req)

	//Check response code and body
	if wr.Code == http.StatusNotFound {
		t.Errorf("User not found")
	}

	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}

func TestDeleteUser(t *testing.T) {

	//Open test database
	testdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	testDB := models.DBMigrate(testdb)

	//Create test user
	testUser := models.User{Username: "mrandall24", Passwordhash: "password", Firstname: "matt", Lastname: "randall", Email: "mr@gmail.com"}

	//Create test request and recorder
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users/mrandall24", nil)

	//Call function
	resultUser := findUser(testDB, testUser.Username, wr, req)

	//Check response code and body
	if wr.Code == http.StatusNotFound {
		t.Errorf("User not found")
	}

	if resultUser.Username != testUser.Username {
		t.Errorf("Expected user %s, got %s", testUser.Username, resultUser.Username)
	}
}
