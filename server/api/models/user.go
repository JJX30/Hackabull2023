// Defines a database model for a user and event and a function to create and
// migrate the table to the database using the GORM library

package models

import (
	"gorm.io/gorm"
)

// Struct that defines the fields that represent a user in the application
type User struct {

	// The gorm.Model is a struct provided by the GORM library that
	// contains common fields used for most database tables.
	// By embedding gorm.Model in the User struct, the fields ID,
	// CreatedAt, UpdatedAt, and DeletedAt are added to the User struct
	// automatically.
	gorm.Model
	Username           string   `json:"username"`
	Passwordhash       string   `json:"passwordhash"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	ProfilePicturePath string   `json:"profilepicturepath"`
	Text               []string `json:"text"`
	ConfidenceLevel    []int    `json:"confidencelevel"`
}

// Struct that defines the fields that represent a user in the application
type Event struct {

	// The gorm.Model is a struct provided by the GORM library that
	// contains common fields used for most database tables.
	// By embedding gorm.Model in the User struct, the fields ID,
	// CreatedAt, UpdatedAt, and DeletedAt are added to the User struct
	// automatically.
	gorm.Model
	Name      string `json:"name"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Location  string `json:"location"`
	Interests string `json:"interests"`
}

// The DBMigrate function takes a GORM database instance (db) and uses
// the AutoMigrate method to automatically create the necessary table(s)
// and migrate the schema to match the struct. This ensures that the
// database schema is up-to-date with the latest changes in the model.
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Event{})
	return db
}
