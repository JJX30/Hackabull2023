package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a new file with a specific name and path
	f, err := os.Create("./uploaded-file/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func UploadUserPFP(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

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

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Make dir for user to store their pfp
	path := "./user-profile-images/" + username + "/"

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Create a new file with a specific name and path
	f, err := os.Create(path + handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add pfp path to models.User{} for that username
	frontendPath := path + handler.Filename

	user.ProfilePicturePath = frontendPath

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	respondJSON(w, http.StatusCreated, frontendPath)
}

func UploadUserAudio(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")

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

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Make dir for user to store their pfp
	path := "./user-audio-file/" + username + "/"

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Create a new file with a specific name and path
	f, err := os.Create(path + handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new file with a specific name and path
	g, err := os.Create("./uploaded-file/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(g, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run the Python ML model
	RunPython()

	e := os.Remove("./uploaded-file/" + handler.Filename)
	if e != nil {
		log.Fatal(e)
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	respondJSON(w, http.StatusOK, "Uploaded Audio")
}
