package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Define your model structs here, if you haven't already.
	// For example:
	// type Appointments struct {
	//     gorm.Model
	//     // Define your fields here
	// }
	// Migrate the schema
	database.AutoMigrate(
		&Appointment{},
		&Member{},
		&Dentist{},
	)
	db = database
	
}
