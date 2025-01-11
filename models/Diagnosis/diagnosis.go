package diagnosis_models

type Diagnosis struct {
	Id                  int     `json:"id"`
	MrnId               *string `json:"mrn"`
	FhirId              *string `json:"fhir_id"`
	EncounterId         *string `json:"encounter"`
	Diagnosis_id        *string `gorm:"column:diagnosis_id"` // Foreign key column `json:"task_id_fk"`
	Diagnosis           ICD10   `gorm:"foreignKey:Diagnosis_id;references:Id"`
	PrimaryDiagnosis    bool    `json:"primary_diagnosis"`
	OnSetDate           *string `json:"on_set_date"`
	ClouserDate         *string `json:"clouser_date"`
	Confirmation        *string `json:"confirmation"`
	Type                *string `json:"type"`
	ResponsibleProvider *string `json:"responsible_provider"`
	ClinicalService     *string `json:"clinical_service"`
	Stage               *string `json:"stage"`
	Comorbidity         *string `json:"comorbidity"`
	Severity            *string `json:"severity"`
	CancerStage         *string `json:"cancer_stage"`
	Active              bool    `json:"active"`
	ExternalCause       bool    `json:"external_cause"`
	UnderlyingCause     bool    `json:"underlying_cause"`
	ChronicDiagnosis    bool    `json:"chronic_diagnosis"`
	InfectiousDiagnosis bool    `json:"infectious_diagnosis"`
	AddToProblemList    bool    `json:"add_to_problem_list"`
	LastUpdate          *string `json:"last_update"`
	Comment             *string `json:"comment"`
	IsDrafted           bool    `json:"is_drafted"`
}

type ICD10 struct {
	Id     int      `json:"id"`
	Code   *string  `json:"code"`
	Name   *string  `json:"name"`
	Price  *float64 `json:"price"`
	Status *string  `json:"status"`
}

func (Diagnosis) TableName() string {
	return "Diagnosis_diagnosis"
}

func (ICD10) TableName() string {
	return "Billing_icd10"
}
