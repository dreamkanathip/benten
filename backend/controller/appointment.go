package controller

import (
	"net/http"

	"time"

	"github.com/dreamkanathip/project-66/entity"
	"github.com/gin-gonic/gin"
)

//POST
func CreateAppointment(c *gin.Context) {
	var appointment entity.Appointment
	var member entity.Member
	var dentist entity.Dentist

	// Bind JSON request to the appointment variable
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the member exists
	if tx := entity.DB().Where("id = ?", appointment.MemberID).First(&member); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}

	// Check if the dentist exists
	if tx := entity.DB().Where("id = ?", appointment.DentistID).First(&dentist); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dentist not found"})
		return
	}

	// Create a new appointment object
	newAppointment := entity.Appointment{
		Datie:  appointment.Datie,
		Problem:   appointment.Problem,
		Time: appointment.Time,

		DentistID: appointment.DentistID,
		Dentist: dentist,
		MemberID: appointment.MemberID,
		Member: member,
	}

	// Save the new appointment to the database
	if err := entity.DB().Create(&newAppointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Appointment created successfully", "appointment": newAppointment})
}
	// GET /Member,Dentist/:id
func GetAppointment(c *gin.Context) {
	var appointment entity.Appointment
	id := c.Param("id")
	if err := entity.DB().Preload("Member").Preload("Dentist").Raw("SELECT * FROM appointments WHERE id = ?", id).Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}
	// GET /appointments	
func ListAppointment(c *gin.Context) {
	var appointment []entity.Appointment
	if err := entity.DB().Preload("Member").Preload("Dentist").Raw("SELECT * FROM appointments").Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}
// DELETE /apointments/:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM appointments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func FormatTime(timeValue time.Time) string {
    // Format the time.Time value as a string
    return timeValue.Format("DD/MM/YYYY")
}

// func GetAppointmentbyUsernameMember(c *gin.Context) {
// 	var member entity.Member
// 	var appointment entity.Appointment
// 	// usernameM := c.Param("username")
// 	// Select users with orders between 100 and 200

// 	if err := entity.DB().Where("Name = ?", ).First(&member); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
	
	
// 	if err := entity.DB().Raw("SELECT * FROM appointments WHERE id = ?", username).Find(&username).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": username})
// }