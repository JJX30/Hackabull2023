// Defines several functions for managing user data in a web API
// using the Go programming language and the Gorilla web toolkit.

package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"github.com/JJX30/Rizzerator-App/server/api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUserProfilePicture(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// vars["username"] is used to extract the value of this variable.
	username := vars["username"]
	user := findUser(db, username, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user.ProfilePicturePath)
}

func LogIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// Creates a new empty models.User struct
	user := models.User{}

	// Decodes the JSON data from the request body into the user struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Makes sure that the request body is closed before the function exits.
	defer r.Body.Close()

	var checkUser models.User
	db.Where("username = ?", user.Username).First(&checkUser)

	if user.Username != checkUser.Username || !(CheckPasswordHash(user.Passwordhash, checkUser.Passwordhash)) {
		respondError(w, http.StatusInternalServerError, "Username or password is incorrect. Please try again.")
		return
	}

	respondJSON(w, http.StatusCreated, "Successful Login")
}
func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// Creates a new empty models.User struct
	user := models.User{}

	// Decodes the JSON data from the request body into the user struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Makes sure that the request body is closed before the function exits.
	defer r.Body.Close()

	// Checks if password meets requirements for good entropy
	if !Password(user.Passwordhash) {
		respondError(w, http.StatusBadRequest, "Password does not meet the requirements. Please try again.")
		return
	}

	// Hashes the user's password
	h, err := HashPassword(user.Passwordhash)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user.Passwordhash = h

	username := user.Username
	var checkUsername models.User
	db.Where("username = ?", username).First(&checkUsername)

	if username == checkUsername.Username {
		respondError(w, http.StatusInternalServerError, "Username already taken. Try another one.")
		return
	}

	// Saves the user to the database
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, "Successful Registration")
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) is used to extract the variables from the incoming
	// request's URL. The variables are returned as a map where the keys
	// are the variable names defined in the URL path and the values are
	// the actual values specified in the request.
	vars := mux.Vars(r)

	// vars["username"] is used to extract the value of this variable.
	username := vars["username"]
	user := findUser(db, username, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Creates a new empty models.User struct
	user := models.User{}

	// Decodes the JSON data from the request body into the user struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Makes sure that the request body is closed before the function exits.
	defer r.Body.Close()

	// Hashes the user's password
	h := sha256.New()
	h.Write([]byte(user.Passwordhash))
	user.Passwordhash = string(h.Sum(nil))

	// Saves the user to the database
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	user := findUser(db, username, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	user := findUser(db, username, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func findUser(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *models.User {
	user := models.User{}
	if err := db.Find(&user, models.User{Username: username}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
