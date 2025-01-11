package main

import (
	allergy_controllers "his_apis_go/controllers/Allergy"
	diagnosis_controllers "his_apis_go/controllers/Diagnosis"
	eavs_controllers "his_apis_go/controllers/Eav"
	task_controllers "his_apis_go/controllers/Task"
	"his_apis_go/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.New()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/list-task", task_controllers.TaskIndex)
	r.GET("/list-forms", eavs_controllers.FormIndex)
	r.GET("/get-allergies", allergy_controllers.GetAllergyIntolerances)
	r.POST("/get-allergies", allergy_controllers.CreateAllergyIntolerance)
	r.PUT("/get-allergies/:id", allergy_controllers.UpdateAllergyIntolerance)
	r.GET("/get-diagnoses", diagnosis_controllers.GetDiagnoses)
	r.POST("/get-diagnoses", diagnosis_controllers.CreateDiagnosis)
	r.PUT("/get-diagnoses/:id", diagnosis_controllers.UpdateDiagnosis)
	r.DELETE("/get-diagnoses/:id", diagnosis_controllers.DeleteDiagnosis)

	r.Run(":4000")
}
