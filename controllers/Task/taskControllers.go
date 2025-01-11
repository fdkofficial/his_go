package task_controllers

import (
	"his_apis_go/initializers"
	task_models "his_apis_go/models/Task"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	task := task_models.Task{}
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		c.Status(400)
		return
	}
}

// func TaskIndex(c *gin.Context) {
// 	// Get the posts
// 	mrn := c.Query("mrn")
// 	task := c.Query("task")
// 	// encounter := c.Query("encounter")
// 	// log.Print(mrn)
// 	var Tasks_task []task_models.Task
// 	result := initializers.DB.Where("patient_id = ? ", mrn).Where("task_id = ?", task).Find(&Tasks_task)
// 	// .Where("encounter_id = ? ", encounter).Find(&Tasks_task)
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	// Respond with the filtered tasks
// 	c.JSON(200, gin.H{
// 		"tasks": Tasks_task,
// 		"count": len(Tasks_task),
// 	})

// }
func TaskIndex(c *gin.Context) {
	// Get the query parameters
	mrn := c.Query("mrn")
	encounter := c.Query("encounter")

	// Define a slice to hold the tasks
	var Tasks_task []task_models.Task

	// Query with Preload to load the related Tasks_categorie
	result := initializers.DB.
		Preload("Task"). // Preload the related table
		Where("patient_id = ?", mrn).
		Where("encounter_id = ?", encounter).
		Find(&Tasks_task)

	// Handle query errors
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Respond with the filtered tasks
	c.JSON(200, gin.H{
		"data":  Tasks_task,
		"count": len(Tasks_task),
	})
}
