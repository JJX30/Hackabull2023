package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/JJX30/Rizzerator-App/server/api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateEvent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	checkEvent := models.Event{}
	db.Where("name = ? AND date = ? AND time = ? AND location = ? AND interests = ?", event.Name, event.Date, event.Time, event.Location, event.Interests).Find(&checkEvent)

	// Does the Find function find all structs with the same
	// structure or can it also find same info specifically?

	// FIX THIS
	if event.Name == checkEvent.Name {
		respondError(w, http.StatusInternalServerError, "Event has already been created.")
		return
	}

	if err := db.Save(&event).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, event)
}

func GetAllEvents(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	events := []models.Event{}
	db.Find(&events)
	respondJSON(w, http.StatusOK, events)
}

func GetEvent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) is used to extract the variables from the incoming
	// request's URL. The variables are returned as a map where the keys
	// are the variable names defined in the URL path and the values are
	// the actual values specified in the request.
	vars := mux.Vars(r)

	// vars["name"] is used to extract the value of this variable.
	name := vars["name"]
	event := findEvent(db, name, w, r)
	if event == nil {
		return
	}
	respondJSON(w, http.StatusOK, event)
}

func UpdateEvent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	event := findEvent(db, name, w, r)
	if event == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&event).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, event)
}

func DeleteEvent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	event := findEvent(db, name, w, r)
	if event == nil {
		return
	}
	if err := db.Delete(&event).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func SearchEvent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	event := models.Event{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	allEvents := []models.Event{}

	where := buildWhereClause(&event)

	db.Where(where).Find(&allEvents)

	// Or for individual events

	//db.Where(db.Where("name LIKE ?", event.Name).Where(db.Where("date LIKE ?", event.Date).Or("time LIKE ?", event.Time).Or("location LIKE ?", event.Location).Or("interests LIKE ?", event.Interests))).Find(allEvents)
	// db.Where("name LIKE ?", event.Name).Find(&allEvents)
	// db.Where("date LIKE ?", event.Date).Find(&allEvents)
	// db.Where("time LIKE ?", event.Time).Find(&allEvents)
	// db.Where("location LIKE ?", event.Location).Find(&allEvents)
	// db.Where("interests LIKE ?", event.Interests).Find(&allEvents)

	// db.Where(db.Where("name LIKE ?", event.Name).Where(db.Where("date LIKE ?", event.Date).Or("time LIKE ?", event.Time).Or("location LIKE ?", event.Location).Or("interests LIKE ?", event.Interests))).Find(&allEvents)

	//db.Where("name LIKE ? AND date LIKE ? AND time LIKE ? AND location LIKE ? AND interests LIKE ?", event.Name, event.Date, event.Time, event.Location, event.Interests).Find(&allEvents)

	respondJSON(w, http.StatusOK, allEvents)
}

func findEvent(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *models.Event {
	event := models.Event{}
	if err := db.Find(&event, models.Event{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &event
}

// Helper function for SearchEvent function
func buildWhereClause(e *models.Event) string {
	var temp []string
	if e.Name != "" {
		temp = append(temp, "name LIKE '%"+e.Name+"%'")
	}
	if e.Date != "" {
		temp = append(temp, "date LIKE '%"+e.Date+"%'")
	}
	if e.Time != "" {
		temp = append(temp, "time LIKE '%"+e.Time+"%'")
	}
	if e.Location != "" {
		temp = append(temp, "location LIKE '%"+e.Location+"%'")
	}
	if e.Interests != "" {
		temp = append(temp, "interests LIKE '%"+e.Interests+"%'")
	}
	return strings.Join(temp, " AND ")
}
