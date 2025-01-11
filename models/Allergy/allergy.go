package allergy_models

type AllergyCategory struct {
	Id       int     `json:"id"`
	Category *string `json:"category"`
}
type Substance struct {
	Id              int     `json:"id"`
	Name            *string `json:"name"`
	Terminology     *string `json:"terminology"`
	TerminologyAxis *string `json:"terminology_axis"`
	Code            *string `json:"code"`
	PrincipleType   *string `json:"principle_type"`
}
type AllergyIntolerance struct {
	Id           int     `json:"id"`
	Mrn_id       string  `gorm:"column:mrn_id"`
	Encounter_id int     `gorm:"column:encounter_id"`
	CreatedBy_id int     `json:"created_by_id"`
	FhirId       *string `json:"fhir_id"`
	// Substance       *string `json:"substance"`
	Substance_fk *string   `gorm:"column:substance_id"` // Foreign key column `json:"task_id_fk"`
	Substance    Substance `gorm:"foreignKey:Substance_fk;references:Id"`
	Category_fk  *string   `gorm:"column:category_id"`
	// Category        AllergyCategory `gorm:"foreignKey:Category_fk;references:Id"`
	Severity        *string `json:"severity"`
	ReactionType    *string `json:"reaction_type"`
	Active          bool    `json:"active"`
	ObservedInVisit bool    `json:"observed_in_visit"`
	OnsetDate       *string `json:"onset_date"`
	ClouserDate     *string `json:"clouser_date"`
	Status          *string `json:"status"`
	Confirmation    *string `json:"confirmation"`
	UpdatedAt       *string `json:"updated_at"`
	ReviewedAt      *string `json:"reviewed_at"`
	MarkAsReviewed  bool    `json:"mark_as_reviewed"`
	Reviewed        *string `json:"reviewed"`
	Criticality     *string `json:"criticality"`
	RecordedDate    *string `json:"recorded_date"`
	InfoSource      *string `json:"info_source"`
	Comment         *string `json:"comment"`
	Details         *string `json:"details"`
	IsDrafted       bool    `json:"is_drafted"`
}

func (AllergyIntolerance) TableName() string {
	return "Allergy_allergyintolerance"
}

func (Substance) TableName() string {
	return "Allergy_substance"
}
func (AllergyCategory) TableName() string {
	return "Allergy_allergycategory"
}
