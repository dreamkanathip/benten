package controller

import (
	"net/http"
	"github.com/dreamkanathip/project-66/entity"
	"github.com/gin-gonic/gin"
)

// POST /dentists
func CreateDentist(c *gin.Context) {
	var dentist entity.Dentist

	// Bind JSON request to the dentist variable
	if err := c.ShouldBindJSON(&dentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new dentist to the database
	if err := entity.DB().Create(&dentist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dentist"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Dentist created successfully", "dentist": dentist})
}

// GET /dentists/:id
func GetDentist(c *gin.Context) {
	var dentist entity.Dentist
	id := c.Param("id")
	if err := entity.DB().Where("id = ?", id).First(&dentist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dentist})
}

// GET /dentists
func ListDentists(c *gin.Context) {
	var dentists []entity.Dentist
	if err := entity.DB().Find(&dentists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dentists})
}

// DELETE /dentists/:id
func DeleteDentist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM dentists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
