package models

import "time"

// Document struct
type Document struct {
	Id            uint    `gorm:"primaryKey" json:"id"`
	MrnId         *string `gorm:"index" json:"mrn_id"`       // ForeignKey to Patients
	EncounterId   *uint   `gorm:"index" json:"encounter_id"` // ForeignKey to Encounter
	AuthorId      *uint   `gorm:"index" json:"author_id"`    // ForeignKey to User
	FhirId        *string `gorm:"default:'No FHIR Resource created for this Order'" json:"fhir_id"`
	ValueSetId    *uint   `gorm:"index" json:"value_set_id"` // ForeignKey to ValueSetDocs
	EntryDate     *string `gorm:"autoCreateTime" json:"entry_date"`
	LastUpdated   *string `gorm:"autoUpdateTime" json:"last_updated"`
	EntryTime     *string `json:"entry_time"`
	Status        *string `json:"status"`
	Type          *string `json:"type"`
	Subject       *string `json:"subject"`
	Document      *string `json:"document"`
	PowerDocument *string `json:"power_document"`
	Comment       *string `json:"comment"`
	AutherSignId  *uint   `gorm:"index" json:"author_sign_id"` // ForeignKey to User
	ReviewedById  *uint   `gorm:"index" json:"reviewed_by_id"` // ForeignKey to User
	ApproverId    *uint   `gorm:"index" json:"approver_id"`    // ForeignKey to User
}

// ScannedDocument struct
type ScannedDocument struct {
	Id          uint  `gorm:"primaryKey" json:"id"`
	MrnId       *uint `gorm:"index" json:"mrn_id"`       // ForeignKey to Patients
	EncounterId *uint `gorm:"index" json:"encounter_id"` // ForeignKey to Encounter
	Title       *string
	Status      *string
	Type        *string
	CreatedById *uint   `gorm:"index" json:"created_by"` // ForeignKey to User
	EntryDate   *string `gorm:"autoCreateTime"`

	Notes []ScannedDoc `gorm:"foreignKey:Details"`
}

// ScannedDocs struct
type ScannedDoc struct {
	Id      uint `gorm:"primaryKey"`
	Doc     string
	Details uint // ForeignKey to ScannedDocument
}

// Templates struct
type Template struct {
	Id    uint    `gorm:"primaryKey"`
	Title string  `json:"title"`
	Text  *string `json:"text"`
}

// DocumentHistory struct
type DocumentHistory struct {
	Id           uint  `gorm:"primaryKey"`
	DocId        *uint `gorm:"index"` // ForeignKey to Document
	MrnId        *uint `gorm:"index"` // ForeignKey to Patients
	AuthorId     *uint `gorm:"index"` // ForeignKey to User
	EntryDate    *string
	EntryTime    *string
	Status       *string
	Subject      *string
	Document     *string
	Comment      *string
	AutherSignId *uint  `gorm:"index"` // ForeignKey to User
	ReviewedById *uint  `gorm:"index"` // ForeignKey to User
	ApproverId   *uint  `gorm:"index"` // ForeignKey to User
	UpdatedById  *uint  `gorm:"index"` // ForeignKey to User
	UpdatedDate  string `gorm:"autoUpdateTime"`
}

// ValueSetDocs struct
type ValueSetDoc struct {
	Id      uint    `json:"id,omitempty"` // Auto-generated
	Code    *string `json:"code"`
	System  *string `json:"system"`
	Display *string `json:"display"`
	Status  string  `gorm:"default:'InActive'"`
}

// Certification struct
type Certification struct {
	Id           int       `json:"id,omitempty"` // Auto-generated
	MrnId        string    `json:"mrn"`
	EncounterId  int       `json:"encounter"`
	AuthorId     *int      `json:"author_id"`
	ValueSetId   int       `json:"value_set_id"`
	TemplateId   uint      `json:"template_id"` // Make sure TemplateId is an integer
	EntryDate    time.Time `json:"entry_date"`
	EntryTime    string    `json:"entry_time"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	Subject      string    `json:"subject"`
	Document     string    `json:"document"`
	AutherSignId *string   `json:"auther_signed_id"`
	ReviewedById *int      `json:"reviewed_by_id"`
	ApproverId   *int      `json:"approver_id"`

	// Relationship with Template
	ValueSet ValueSetDoc `gorm:"foreignKey:ValueSetId" json:"value_set"`
	Template Template    `gorm:"foreignKey:TemplateId" json:"template"`
}

// ProgressNote struct
type ProgressNote struct {
	Id           uint      `json:"id"`
	MrnId        *string   `json:"mrn"`          // ForeignKey to Patients
	EncounterId  *uint     `json:"encounter"`    // ForeignKey to Encounter
	AuthorId     *uint     `json:"author_id"`    // ForeignKey to User
	ValueSetId   *uint     `json:"value_set_id"` // ForeignKey to ValueSetDocs
	TemplateId   *uint     `json:"template_id"`  // ForeignKey to Templates
	FhirId       string    `gorm:"default:'No FHIR Resource created for this Document'" json:"fhir_id"`
	EntryDate    time.Time `json:"entry_date"`
	EntryTime    string    `json:"entry_time"`
	LastUpdated  time.Time `gorm:"autoUpdateTime"`
	Status       *string   `json:"status"`
	Type         *string   `json:"type"`
	Subject      *string   `json:"subject"`
	Document     *string   `json:"document"`
	AutherSignId *string   `json:"auther_signed_id"`
	ReviewedById *uint     `json:"reviewed_by_id"`
	ApproverId   *uint     `json:"approver_id"`

	// Relationship with Template
	ValueSet ValueSetDoc `gorm:"foreignKey:ValueSetId" json:"value_set"`
	Template Template    `gorm:"foreignKey:TemplateId" json:"template"`
}

// Specify the table name
func (Document) TableName() string {
	return "Document_document" // Set your table name here if it's different
}

// Specify the table name
func (ProgressNote) TableName() string {
	return "Document_progressnote" // Set your table name here if it's different
}

// Specify the table name
func (Certification) TableName() string {
	return "Document_certification" // Set your table name here if it's different
}

// Specify the table name
func (ValueSetDoc) TableName() string {
	return "Document_valuesetdocs" // Set your table name here if it's different
}

// Specify the table name
func (Template) TableName() string {
	return "Document_templates" // Set your table name here if it's different
}

// Specify the table name
func (ScannedDocument) TableName() string {
	return "Document_scanneddocument" // Set your table name here if it's different
}

// Specify the table name
func (ScannedDoc) TableName() string {
	return "Document_scanneddocs" // Set your table name here if it's different
}

// Specify the table name
func (DocumentHistory) TableName() string {
	return "Document_documenthistory" // Set your table name here if it's different
}
