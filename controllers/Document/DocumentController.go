package document_controllers

import (
	// "encoding/json"
	"his_apis_go/initializers"
	document_models "his_apis_go/models/Document"

	"log"

	"strconv"

	"fmt"

	"strings"

	"time"

	"github.com/gin-gonic/gin"
)

func GetDocument(c *gin.Context) {
	// Get query parameters
	mrn := c.Query("mrn")
	encounter := c.Query("encounter")
	page := c.DefaultQuery("page", "1")           // Default to page 1 if not provided
	pageSize := c.DefaultQuery("page_size", "10") // Default to 10 records per page if not provided

	log.Print("MRN:", mrn, "Encounter:", encounter, "Page:", page, "PageSize:", pageSize)

	var documents []document_models.Document
	// Initialize the query
	query := initializers.DB.Model(&document_models.Document{})

	// Apply filters if present
	if mrn != "" {
		query = query.Where("mrn_id = ?", mrn)
	}
	if encounter != "" {
		query = query.Where("encounter_id = ?", encounter)
	}

	// Convert page and pageSize to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		c.JSON(400, gin.H{"error": "Invalid page_size parameter"})
		return
	}

	// Apply pagination
	offset := (pageInt - 1) * pageSizeInt
	query = query.Offset(offset).Limit(pageSizeInt)

	// Execute the query
	if err := query.Find(&documents).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Count total records for pagination metadata
	var totalRecords int64
	if err := initializers.DB.Model(&document_models.Document{}).Count(&totalRecords).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return the result as JSON
	c.JSON(200, gin.H{
		"data":        documents,
		"count":       len(documents),
		"total":       totalRecords,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (totalRecords + int64(pageSizeInt) - 1) / int64(pageSizeInt), // Calculate total pages
	})
}

// POST: Create a new AllergyIntolerance
func CreateDocument(c *gin.Context) {
	var input document_models.Document
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
func UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var existing document_models.Document
	if err := initializers.DB.First(&existing, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	var input document_models.Document
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

func HandleProgressNote(c *gin.Context) {
	method := c.Request.Method

	switch method {
	case "GET":
		// Handle GET request with pagination
		mrn := c.Query("mrn")
		encounter := c.Query("encounter")
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("page_size", "10")

		var progressNotes []document_models.ProgressNote
		query := initializers.DB.Model(&document_models.ProgressNote{})

		if mrn != "" {
			query = query.Where("mrn_id = ?", mrn)
		}
		if encounter != "" {
			query = query.Where("encounter_id = ?", encounter)
		}

		pageInt, _ := strconv.Atoi(page)
		pageSizeInt, _ := strconv.Atoi(pageSize)
		offset := (pageInt - 1) * pageSizeInt

		query = query.Offset(offset).Limit(pageSizeInt)
		if err := query.Find(&progressNotes).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var totalRecords int64
		initializers.DB.Model(&document_models.ProgressNote{}).Count(&totalRecords)

		c.JSON(200, gin.H{
			"data":        progressNotes,
			"count":       len(progressNotes),
			"total":       totalRecords,
			"page":        pageInt,
			"page_size":   pageSizeInt,
			"total_pages": (totalRecords + int64(pageSizeInt) - 1) / int64(pageSizeInt),
		})

	case "POST":
		// Handle POST request
		// var newProgressNote document_models.ProgressNote
		var newProgressNote struct {
			EntryDate string `json:"entry_date"`
			EntryTime string `json:"entry_time"`
			document_models.ProgressNote
		}
		if err := c.ShouldBindJSON(&newProgressNote); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if newProgressNote.EntryDate != "" {
			// Parse the date string to time.Time
			parsedDate, err := time.Parse("2006-01-02", newProgressNote.EntryDate)
			if err != nil {
				// Handle error: invalid date format
				c.JSON(400, gin.H{"error": "Invalid date format", "details": err.Error()})
				return
			}

			// Assign the parsed time to the Certification struct
			newProgressNote.ProgressNote.EntryDate = parsedDate
		}

		// Check if EntryTime is not empty and convert it to time.Time
		if newProgressNote.EntryTime != "" {
			// Validate the time format (HH:MM:SS) without converting to time.Time
			_, err := time.Parse("15:04:05", newProgressNote.EntryTime)
			if err != nil {
				// Handle error: invalid time format
				c.JSON(400, gin.H{"error": "Invalid time format", "details": err.Error()})
				return
			}

			// Assign the validated time string to the Certification struct
			newProgressNote.ProgressNote.EntryTime = newProgressNote.EntryTime
		} else {
			// If EntryTime is empty, set it to a default value (e.g., "00:00:00")
			newProgressNote.ProgressNote.EntryTime = "00:00:00"
		}

		if err := initializers.DB.Create(&newProgressNote.ProgressNote).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "ProgressNote created successfully", "data": newProgressNote})

	case "PUT":
		// Handle PUT request
		var updatedProgressNote document_models.ProgressNote
		if err := c.ShouldBindJSON(&updatedProgressNote); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if err := initializers.DB.Model(&document_models.ProgressNote{}).Where("id = ?", updatedProgressNote.Id).Updates(updatedProgressNote).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "ProgressNote updated successfully", "data": updatedProgressNote})

	default:
		c.JSON(405, gin.H{"error": "Method not allowed"})
	}
}

func Certifications(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		// Handle GET request
		mrn := c.DefaultQuery("mrn", "")
		id := c.DefaultQuery("id", "")
		from := c.DefaultQuery("from", "")
		to := c.DefaultQuery("to", "")
		status := c.DefaultQuery("status[]", "default_value")

		encounter := c.DefaultQuery("encounter", "")

		filters := []string{}

		// Build the query filters based on provided query params
		if id != "" {
			filters = append(filters, fmt.Sprintf("id = '%s'", id))
		}
		if encounter != "" && encounter != "null" {
			filters = append(filters, fmt.Sprintf("encounter_id = '%s'", encounter))
			if mrn != "" {
				filters = append(filters, fmt.Sprintf("mrn_id = '%s'", mrn))
			}
		} else if mrn != "" {
			filters = append(filters, fmt.Sprintf("mrn_id = '%s'", mrn))
		}
		if from != "" {
			filters = append(filters, fmt.Sprintf("entry_date >= '%s'", from))
		}
		if to != "" {
			filters = append(filters, fmt.Sprintf("entry_date <= '%s'", to))
		}
		if status != "default_value" && status != "" {
			statusSlice := strings.Split(status, ",") // Handle multiple statuses
			statusFilter := fmt.Sprintf("status IN ('%s')", strings.Join(statusSlice, "','"))
			filters = append(filters, statusFilter)
		}

		fmt.Println("Filters:", filters)

		// Build query string
		whereClause := strings.Join(filters, " AND ")
		query := initializers.DB.Model(&document_models.Certification{}).
			Preload("ValueSet").
			Preload("Template").
			Order("id DESC")

		if whereClause != "" {
			query = query.Where(whereClause)
		}

		// Execute query
		var certifications []document_models.Certification
		if err := query.Find(&certifications).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Serialize and return the data
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "successful",
			"data":    certifications,
		})

	case "POST":
		// Define the struct to bind JSON data
		var newCertification struct {
			EntryDate string `json:"entry_date"`
			EntryTime string `json:"entry_time"`
			document_models.Certification
		}

		// Bind JSON payload to the struct
		if err := c.ShouldBindJSON(&newCertification); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
			return
		}

		if newCertification.ValueSetId != 0 {
			newCertification.Certification.ValueSetId = int(newCertification.ValueSet.Id)
		}

		// Ensure TemplateId is properly set
		if newCertification.Template.Id != 0 {
			newCertification.Certification.TemplateId = newCertification.Template.Id
		}

		// Validate TemplateId before saving
		if newCertification.Certification.TemplateId == 0 {
			c.JSON(400, gin.H{"error": "TemplateId must be provided"})
			return
		}

		if newCertification.EntryDate != "" {
			// Parse the date string to time.Time
			parsedDate, err := time.Parse("2006-01-02", newCertification.EntryDate)
			if err != nil {
				// Handle error: invalid date format
				c.JSON(400, gin.H{"error": "Invalid date format", "details": err.Error()})
				return
			}

			// Assign the parsed time to the Certification struct
			newCertification.Certification.EntryDate = parsedDate
		}

		// Check if EntryTime is not empty and convert it to time.Time
		if newCertification.EntryTime != "" {
			// Validate the time format (HH:MM:SS) without converting to time.Time
			_, err := time.Parse("15:04:05", newCertification.EntryTime)
			if err != nil {
				// Handle error: invalid time format
				c.JSON(400, gin.H{"error": "Invalid time format", "details": err.Error()})
				return
			}

			// Assign the validated time string to the Certification struct
			newCertification.Certification.EntryTime = newCertification.EntryTime
		} else {
			// If EntryTime is empty, set it to a default value (e.g., "00:00:00")
			newCertification.Certification.EntryTime = "00:00:00"
		}

		// Convert empty strings in integer fields to NULL or default value
		if newCertification.AuthorId == nil || *newCertification.AuthorId == 0 {
			newCertification.AuthorId = nil // Set to NULL if no valid author_id
		}

		if newCertification.ReviewedById == nil || *newCertification.ReviewedById == 0 {
			newCertification.ReviewedById = nil // Set to NULL if no valid reviewed_by_id
		}

		if newCertification.ApproverId == nil || *newCertification.ApproverId == 0 {
			newCertification.ApproverId = nil // Set to NULL if no valid approver_id
		}

		fmt.Printf("Received Certification: %+v\n", newCertification)

		// Now, insert into the database
		if err := initializers.DB.Create(&newCertification.Certification).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Return success response
		c.JSON(201, gin.H{
			"status":  "ok",
			"message": "Certification created successfully",
			"data":    newCertification,
		})

	default:
		c.JSON(405, gin.H{"error": "Method Not Allowed"})
	}
}

func ScannedDocHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		mrn := c.DefaultQuery("mrn", "")
		id := c.DefaultQuery("id", "")
		from := c.DefaultQuery("from", "")
		to := c.DefaultQuery("to", "")
		status := c.QueryArray("status[]")
		encounter := c.DefaultQuery("encounter", "")

		var documents []document_models.ScannedDocument
		query := initializers.DB

		if id != "" {
			query = query.Where("id = ?", id)
		}
		if encounter != "" && encounter != "null" {
			query = query.Where("encounter_id = ?", encounter)
			if mrn != "" {
				query = query.Where("mrn_id = ?", mrn)
			}
		} else if mrn != "" {
			query = query.Where("mrn_id = ?", mrn)
		}
		if from != "" {
			query = query.Where("entry_date >= ?", from)
		}
		if to != "" {
			query = query.Where("entry_date <= ?", to)
		}
		if len(status) > 0 {
			query = query.Where("status IN ?", status)
		}

		query.Preload("Notes").Order("id DESC").Find(&documents)

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "successful",
			"data":    documents,
		})

	case "POST":
		var newDoc struct {
			CreatedBy uint `json:"created_by"`
			document_models.ScannedDocument
		}
		if err := c.ShouldBindJSON(&newDoc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newDoc.CreatedById = &newDoc.CreatedBy
		if err := initializers.DB.Create(&newDoc).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		notes := c.PostFormArray("notes")
		for _, note := range notes {
			// Remove any carriage return characters
			cleanedNote := strings.ReplaceAll(note, "\r", "")

			scannedDoc := document_models.ScannedDoc{
				Doc:     cleanedNote, // Use the cleaned note
				Details: newDoc.Id,
			}
			initializers.DB.Create(&scannedDoc)
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "successful",
			"data":    newDoc,
		})

	case "PUT":
		id := c.Query("id")
		if id == "" {
			c.JSON(400, gin.H{"error": "id is required"})
			return
		}

		var existingDoc document_models.ScannedDocument
		if err := initializers.DB.First(&existingDoc, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "document not found"})
			return
		}

		if err := c.ShouldBindJSON(&existingDoc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := initializers.DB.Save(&existingDoc).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "successful",
			"data":    existingDoc,
		})

	case "DELETE":
		id := c.Query("id")
		if id == "" {
			c.JSON(400, gin.H{"error": "id is required"})
			return
		}

		if err := initializers.DB.Delete(&document_models.ScannedDocument{}, id).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "successful",
		})

	default:
		c.JSON(405, gin.H{"error": "method not allowed"})
	}
}
