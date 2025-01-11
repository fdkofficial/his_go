package diagnosis_controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	diagnosis_controllers "his_apis_go/controllers/Diagnosis"
	"his_apis_go/initializers"
	diagnosis_models "his_apis_go/models/Diagnosis"

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
	r.GET("/diagnoses", diagnosis_controllers.GetDiagnoses)
	r.POST("/diagnoses", diagnosis_controllers.CreateDiagnosis)
	r.PUT("/diagnoses/:id", diagnosis_controllers.UpdateDiagnosis)
	return r
}

func TestGetDiagnoses(t *testing.T) {
	// Setup
	router := SetupRouter()

	// Perform the GET request
	req, _ := http.NewRequest("GET", "/diagnoses?mrn=202402238501&encounter=173", nil)
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

func TestCreateDiagnosis(t *testing.T) {
	// Setup
	router := SetupRouter()

	// Sample Diagnosis JSON
	diagnosis := map[string]interface{}{
		"mrn":               "202402238501",
		"fhir_id":           "4fb0e567fb964531bf6a48c2",
		"encounter":         "173",
		"Diagnosis_id":      "40434",
		"primary_diagnosis": true,
		"on_set_date":       "2025-01-11T00:00:00Z",
		"confirmation":      "Final Diagnosis",
		"type":              "Working",
		"severity":          "Cosmetic",
		"active":            true,
		"last_update":       "2025-01-11T00:00:00Z",
		"comment":           "-",
		"is_drafted":        false,
	}

	payload, _ := json.Marshal(diagnosis)

	// Perform the POST request
	req, _ := http.NewRequest("POST", "/diagnoses", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["data"])
	assert.Equal(t, "202402238501", response["data"].(map[string]interface{})["mrn"])
}

func TestUpdateDiagnosis(t *testing.T) {
	// Setup
	router := SetupRouter()

	// Sample existing Diagnosis for Update
	existingDiagnosis := diagnosis_models.Diagnosis{
		MrnId:            ptr("202402238501"),
		FhirId:           ptr("4fb0e567fb964531bf6a48c2"),
		EncounterId:      ptr("173"),
		Diagnosis_id:     ptr("40434"),
		PrimaryDiagnosis: true,
		OnSetDate:        ptr("2025-01-11T00:00:00Z"),
		Confirmation:     ptr("Final Diagnosis"),
		Type:             ptr("Working"),
		Severity:         ptr("Cosmetic"),
		Active:           true,
		LastUpdate:       ptr("2025-01-11T00:00:00Z"),
		Comment:          ptr("-"),
		IsDrafted:        false,
	}
	initializers.DB.Create(&existingDiagnosis)

	// Update Payload
	updatePayload := map[string]interface{}{
		"severity":   "Critical",
		"comment":    "Updated comment",
		"is_drafted": true,
	}

	payload, _ := json.Marshal(updatePayload)

	// Perform the PUT request
	req, _ := http.NewRequest("PUT", "/diagnoses/"+string(existingDiagnosis.Id), bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["data"])
	assert.Equal(t, "Critical", response["data"].(map[string]interface{})["severity"])
	assert.Equal(t, true, response["data"].(map[string]interface{})["is_drafted"])
}
