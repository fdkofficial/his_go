package eavs_controllers

import (
	"his_apis_go/initializers"
	eav_models "his_apis_go/models/EAV"

	// "his_apis_go/models"

	"github.com/gin-gonic/gin"
)

// func FormIndex(c *gin.Context) {
// 	// Get the query parameters
// 	mrn := c.Query("mrn")

// 	// Define a slice to hold the tasks
// 	var savedobservationform []eav_models.SavedObservationForm
// 	// result := initializers.DB.
// 	// 	Preload("Values"). // Preload the related table
// 	// 	Where("mrn_id = ?", mrn).
// 	// 	Find(&savedobservationform)

// 	// Query with Preload to load the related Tasks_categorie
// 	result := initializers.DB.Preload("Values").Raw(`
// 	   SELECT *
// 	    FROM "Forms_savedobservationform"
// 	    JOIN "Forms_savedobservationform_value"
// 	    ON "Forms_savedobservationform".id = "Forms_savedobservationform_value".savedobservationform_id
// 	    WHERE mrn_id = ?
// 	`, mrn).Find(&savedobservationform)

// 	// Handle query errors
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	// Respond with the filtered tasks
// 	c.JSON(200, gin.H{
// 		"forms": savedobservationform,
// 		"count": len(savedobservationform),
// 	})
// }

func FormIndex(c *gin.Context) {
	// Get the query parameters
	mrn := c.Query("mrn")

	// Define slices to hold the data
	var savedObservationForms []eav_models.SavedObservationForm
	var formValues []eav_models.Forms_savedobservationform_value

	// Raw query to load the main table data
	formQuery := `
        SELECT *
        FROM "Forms_savedobservationform"
        WHERE "mrn_id" = ?
    `
	result := initializers.DB.Raw(formQuery, mrn).Scan(&savedObservationForms)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Load the associated values (manual preload)
	if len(savedObservationForms) > 0 {
		var formIDs []uint
		for _, form := range savedObservationForms {
			formIDs = append(formIDs, form.ID)
		}

		valuesQuery := `
            SELECT *
            FROM "Forms_savedobservationform_value"
            WHERE "savedobservationform_id" IN ?
        `
		result = initializers.DB.Raw(valuesQuery, formIDs).Scan(&formValues)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		// Link values to their respective forms
		valueMap := make(map[uint][]eav_models.Forms_savedobservationform_value)
		for _, value := range formValues {
			valueMap[value.Savedobservationform_id] = append(valueMap[value.Savedobservationform_id], value)
		}

		for i, form := range savedObservationForms {
			savedObservationForms[i].Values = valueMap[form.ID]
		}
	}

	// Respond with the filtered data
	c.JSON(200, gin.H{
		"forms": savedObservationForms,
		"count": len(savedObservationForms),
	})
}
