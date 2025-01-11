package allergy_controllers

import (
	"his_apis_go/initializers"
	allergy_models "his_apis_go/models/Allergy"

	// "his_apis_go/models"
	"log"

	"github.com/gin-gonic/gin"
)

// GET: Fetch all AllergyIntolerances
func GetAllergyIntolerances(c *gin.Context) {
	mrn := c.Query("mrn")
	encounter := c.Query("encounter")
	log.Print(mrn, encounter)
	var allergyIntolerances []allergy_models.AllergyIntolerance
	if err := initializers.DB.Preload("Substance").Where("mrn_id = ?", mrn).
		Where("encounter_id = ?", encounter).Find(&allergyIntolerances).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data":  allergyIntolerances,
		"count": len(allergyIntolerances),
	})
}

// POST: Create a new AllergyIntolerance
func CreateAllergyIntolerance(c *gin.Context) {
	var input allergy_models.AllergyIntolerance
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	if err := initializers.DB.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"data": input})
}

// PUT: Update an existing AllergyIntolerance
func UpdateAllergyIntolerance(c *gin.Context) {
	id := c.Param("id")
	var existing allergy_models.AllergyIntolerance
	if err := initializers.DB.First(&existing, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	var input allergy_models.AllergyIntolerance
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	input.Id = existing.Id // Preserve the ID
	if err := initializers.DB.Save(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": input})
}
