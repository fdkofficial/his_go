package diagnosis_controllers

import (
	"his_apis_go/initializers"
	"log"
	"net/http"

	diagnosis_models "his_apis_go/models/Diagnosis"

	"github.com/gin-gonic/gin"
)

func GetDiagnoses(c *gin.Context) {
	mrn := c.Query("mrn")
	encounter := c.Query("encounter")
	log.Print("Fetching diagnoses for MRN:", mrn, "and Encounter:", encounter)

	var diagnoses []diagnosis_models.Diagnosis
	query := initializers.DB.Preload("Diagnosis")

	if mrn != "" {
		query = query.Where("mrn_id = ?", mrn)
	}
	if encounter != "" {
		query = query.Where("encounter_id = ?", encounter)
	}

	if err := query.Find(&diagnoses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  diagnoses,
		"count": len(diagnoses),
	})
}

// POST: Create a new Diagnosis
func CreateDiagnosis(c *gin.Context) {
	var input diagnosis_models.Diagnosis
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// PUT: Update an existing Diagnosis
func UpdateDiagnosis(c *gin.Context) {
	id := c.Param("id")

	var existing diagnosis_models.Diagnosis
	if err := initializers.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diagnosis record not found"})
		return
	}

	var input diagnosis_models.Diagnosis
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Preserve the existing ID
	input.Id = existing.Id

	if err := initializers.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// DELETE: Delete a Diagnosis
func DeleteDiagnosis(c *gin.Context) {
	id := c.Param("id")

	if err := initializers.DB.Delete(&diagnosis_models.Diagnosis{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diagnosis record deleted successfully"})
}
