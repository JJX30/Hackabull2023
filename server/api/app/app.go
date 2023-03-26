// Defines an HTTP server application that handles user data by creating,
// reading, updating, and deleting user information in a SQLite database.

package app

import (
	"log"
	"net/http"

	"github.com/JJX30/Rizzerator-App/server/api/controllers"
	"github.com/JJX30/Rizzerator-App/server/api/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Server container
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Opens database and router
func (a *App) Initialize() {

	db, err := gorm.Open(sqlite.Open("User.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Could not connect to database")
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all routers
func (a *App) setRouters() {
	//User CRUD routes
	a.Get("/api/users", a.GetAllUsers)
	a.Post("/api/users", a.CreateUser)
	a.Get("/api/users/{username}", a.GetUser)
	a.Put("/api/users/{username}", a.UpdateUser)
	a.Delete("/api/users/{username}", a.DeleteUser)

	//User authentication routes
	a.Post("/api/users/register", a.RegisterUser)
	a.Post("/api/users/login", a.LogIn)

	//User Profile Picture
	a.Get("/api/users/profilepicture/{username}", a.GetUserProfilePicture)

	//User Audio
	a.Get("/api/users/audio-file/{username}/{number}", a.GetUserAudio)

	//Event routes
	a.Post("/api/events", a.CreateEvent)
	a.Get("/api/events", a.GetAllEvents)
	a.Get("/api/events/{name}", a.GetEvent)
	a.Put("/api/events/{name}", a.UpdateEvent)
	a.Delete("/api/events/{name}", a.DeleteEvent)
	a.Post("/api/events/search", a.SearchEvent)

	//File upload routes
	a.Post("/upload", a.UploadFile)
	a.Post("/api/upload/profilepicture/{username}", a.UploadUserPFP)
	a.Post("/api/upload/audio-file/{username}/{number}", a.UploadUserAudio)
}

// Router wrapper functions
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage user data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllUsers(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	controllers.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	controllers.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateUser(a.DB, w, r)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteUser(a.DB, w, r)
}

func (a *App) RegisterUser(w http.ResponseWriter, r *http.Request) {
	controllers.RegisterUser(a.DB, w, r)
}

func (a *App) LogIn(w http.ResponseWriter, r *http.Request) {
	controllers.LogIn(a.DB, w, r)
}

func (a *App) GetUserProfilePicture(w http.ResponseWriter, r *http.Request) {
	controllers.GetUserProfilePicture(a.DB, w, r)
}

func (a *App) GetUserAudio(w http.ResponseWriter, r *http.Request) {
	controllers.GetUserAudio(a.DB, w, r)
}

// Handlers to manage event data
func (a *App) CreateEvent(w http.ResponseWriter, r *http.Request) {
	controllers.CreateEvent(a.DB, w, r)
}

func (a *App) GetEvent(w http.ResponseWriter, r *http.Request) {
	controllers.GetEvent(a.DB, w, r)
}

func (a *App) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllEvents(a.DB, w, r)
}

func (a *App) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateEvent(a.DB, w, r)
}

func (a *App) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteEvent(a.DB, w, r)
}

func (a *App) SearchEvent(w http.ResponseWriter, r *http.Request) {
	controllers.SearchEvent(a.DB, w, r)
}

// Handlers to manage image/file data
func (a *App) UploadFile(w http.ResponseWriter, r *http.Request) {
	controllers.UploadFile(w, r)
}

func (a *App) UploadUserPFP(w http.ResponseWriter, r *http.Request) {
	controllers.UploadUserPFP(a.DB, w, r)
}

func (a *App) UploadUserAudio(w http.ResponseWriter, r *http.Request) {
	controllers.UploadUserAudio(a.DB, w, r)
}

// Run http server
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
