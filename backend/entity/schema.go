package entity

import (
	"time"

	"gorm.io/gorm"
)
type Appointment struct {
	gorm.Model
	Datie time.Time 
	Time time.Time 
	Problem  	string 

	MemberID *uint
	Member 	Member `gorm:"foreignKey:MemberID"`
	DentistID *uint
	Dentist Dentist `gorm:"foreignKey:DentistID"`

}
type Member struct {
	gorm.Model
	Username string
	Name string
	Appointment []Appointment `gorm:"foreignKey:MemberID"`
}

type Dentist struct {
	gorm.Model
	Username string
	Name string
	Appointment []Appointment `gorm:"foreignKey:DentistID"`
}