package allergy_controllers_test

import (
	"encoding/json"

	allergy_controllers "his_apis_go/controllers/Allergy"
	"his_apis_go/initializers"
	allergy_models "his_apis_go/models/Allergy"
	"log"

	// "his_apis_go/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=NadaaIT password=Nadaa123 dbname=his port=5432 sslmode=disable"
	initializers.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}
func ptr(s string) *string {
	return &s
}
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get-allergies", allergy_controllers.GetAllergyIntolerances)
	r.POST("/get-allergies", allergy_controllers.CreateAllergyIntolerance)
	r.PUT("/get-allergies/:id", allergy_controllers.UpdateAllergyIntolerance)
	// Existing router setup logic
	return r
}

func TestGetAllergyIntolerances(t *testing.T) {
	// Setup
	router := SetupRouter()

	// Perform the GET request
	req, _ := http.NewRequest("GET", "/get-allergies?mrn=202402238501&encounter=173&pg_no=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["data"])
	assert.GreaterOrEqual(t, int(response["count"].(float64)), 0)
}

func seedTestDatabase() {
	// Seed the database with necessary data for the tests
	substance := allergy_models.Substance{
		Id: 14878, Name: ptr("Substance Name"),
	}
	category := allergy_models.AllergyCategory{
		Id: 10, Category: ptr("Category Name"),
	}
	initializers.DB.Create(&substance)
	initializers.DB.Create(&category)
}

// func TestCreateAllergyIntolerance(t *testing.T) {
// 	// Setup
// 	router := SetupRouter() // Use the router from main.go
// 	seedTestDatabase()

// 	// Create a unique ID using xid
// 	newID := xid.New().Counter()

// 	// Prepare the input payload
// 	allergyIntolerance := map[string]interface{}{
// 		"mrn_id":            "202402238501",
// 		"created_by_id":     1,
// 		"fhir_id":           "da5de5415021476ca74488f2",
// 		"encounter_id":      173,
// 		"substance_fk":      14878,
// 		"severity":          "Mild",
// 		"active":            true,
// 		"observed_in_visit": false,
// 		"onset_date":        "2025-01-08T00:00:00Z",
// 		"clouser_date":      "2025-01-23T00:00:00Z",
// 		"status":            "Suspect",
// 		"confirmation":      "Definitive Adr",
// 		"updated_at":        time.Now().Format(time.RFC3339),
// 		"recorded_date":     "2025-01-10T13:01:00+03:00",
// 		"info_source":       "Patient Survey",
// 		"is_drafted":        false,
// 	}

// 	payload, _ := json.Marshal(allergyIntolerance)

// 	// Perform the POST request
// 	req, _ := http.NewRequest("POST", "/get-allergies", bytes.NewBuffer(payload))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusCreated, w.Code)

// 	var response map[string]interface{}
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response["data"])
// 	assert.Equal(t, newID, int(response["data"].(map[string]interface{})["id"].(float64)))
// }
