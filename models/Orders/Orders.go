package order_models

import (
	"time"
)

type Order struct {
	ID              uint            `gorm:"primaryKey"`
	MrnID           *uint           `gorm:"column:mrn_id"`
	EncounterID     *uint           `gorm:"column:encounter_id"`
	OrderName       string          `gorm:"size:255"`
	OrderNo         string          `gorm:"size:255"`
	Details         string          `gorm:"size:255"`
	Category        string          `gorm:"size:255"`
	Status          string          `gorm:"size:255"`
	OrderType       string          `gorm:"size:255"`
	FrequencyStart  *time.Time      `gorm:"column:frequency_start"`
	FrequencyEnd    *time.Time      `gorm:"column:frequency_end"`
	Frequency       string          `gorm:"size:255"`
	OrderBy         string          `gorm:"size:255"`
	CareProviderID  *uint           `gorm:"column:care_provider_id"`
	CareProvider    *User           `gorm:"foreignKey:CareProviderID"`
	OrderDepartment string          `gorm:"size:255"`
	UnitPrice       float64         `gorm:"default:0"`
	Vat             float64         `gorm:"default:0"`
	Quantity        string          `gorm:"size:255;default:1"`
	FormCacheJson   string          `gorm:"type:text"`
	MedicineOrders  []MedicineOrder `gorm:"foreignKey:OrderID"`
}

type MedicineOrder struct {
	ID                uint     `gorm:"primaryKey"`
	MedicineID        uint     `gorm:"column:medicine_id"`
	Medicine          Medicine `gorm:"foreignKey:MedicineID"`
	OrderID           *uint    `gorm:"column:order_id"`
	Order             *Order   `gorm:"foreignKey:OrderID"`
	StoreID           *uint    `gorm:"column:store_id"`
	MrnID             *uint    `gorm:"column:mrn_id"`
	EncounterID       *uint    `gorm:"column:encounter_id"`
	OrderedByID       uint     `gorm:"column:ordered_by_id"`
	OrderedBy         User     `gorm:"foreignKey:OrderedByID"`
	Dosage            string   `gorm:"size:255"`
	DosageUom         string   `gorm:"size:255"`
	Route             string   `gorm:"size:255"`
	Quantity          int      `gorm:"default:1"`
	Priority          string   `gorm:"size:255;default:stat"`
	OrderToDepartment string   `gorm:"size:255"`
	Status            string   `gorm:"size:255;default:Drafted"`
	Comments          string   `gorm:"type:text"`
}

type ImagingOrder struct {
	ID              uint     `gorm:"primaryKey"`
	MrnID           *uint    `gorm:"column:mrn_id"`
	OrderID         *uint    `gorm:"column:order_id"`
	Order           *Order   `gorm:"foreignKey:OrderID"`
	ImagingID       *uint    `gorm:"column:imaging_id"`
	Imaging         *Imaging `gorm:"foreignKey:ImagingID"`
	EncounterID     *uint    `gorm:"column:encounter_id"`
	ProcedureString string   `gorm:"size:255"`
	Priority        string   `gorm:"size:255;default:stat"`
	UnitPrice       int      `gorm:"default:0"`
	Comments        string   `gorm:"type:text"`
	CareProviderID  *uint    `gorm:"column:care_provider_id"`
	CareProvider    *User    `gorm:"foreignKey:CareProviderID"`
}

type LaboratoryOrder struct {
	ID                  uint         `gorm:"primaryKey"`
	MrnID               *uint        `gorm:"column:mrn_id"`
	EncounterID         *uint        `gorm:"column:encounter_id"`
	OrderID             *uint        `gorm:"column:order_id"`
	Order               *Order       `gorm:"foreignKey:OrderID"`
	LoincID             uint         `gorm:"column:loinc_id"`
	Loinc               Panel        `gorm:"foreignKey:LoincID"`
	FhirID              string       `gorm:"size:255;default:'No FHIR Resource created for this Order'"`
	OrderName           string       `gorm:"size:255"`
	OrderedByID         uint         `gorm:"column:ordered_by_id"`
	OrderedBy           User         `gorm:"foreignKey:OrderedByID"`
	Priority            string       `gorm:"size:255;default:stat"`
	PathologySection    string       `gorm:"size:255"`
	OrderToDepartmentID *uint        `gorm:"column:order_to_department_id"`
	PeriodID            *uint        `gorm:"column:period_id"`
	Period              *OrderPeriod `gorm:"foreignKey:PeriodID"`
	AccessionNo         string       `gorm:"size:255;default:random_order_no"`
	Quantity            int          `gorm:"default:0"`
	Status              string       `gorm:"size:255;default:Drafted"`
	OrderStatusID       *uint        `gorm:"column:order_status_id"`
	OrderStatus         *OrderStatus `gorm:"foreignKey:OrderStatusID"`
	DepartmentStatus    string       `gorm:"size:255;default:Drafted"`
	Specimen            string       `gorm:"size:255"`
	SpecimenSite        string       `gorm:"size:255"`
	ClinicalInstruction string       `gorm:"size:255"`
	CareProviderID      *uint        `gorm:"column:care_provider_id"`
	CareProvider        *User        `gorm:"foreignKey:CareProviderID"`
	UnitPrice           int          `gorm:"default:0"`
	SignedByID          *uint        `gorm:"column:signed_by_id"`
	SignedBy            *User        `gorm:"foreignKey:SignedByID"`
	DiscPercentage      *int         `gorm:"column:disc_percentage"`
	ForDiagnosisID      *uint        `gorm:"column:for_diagnosis_id"`
	Percentage          *bool        `gorm:"default:false"`
	Favourite           bool         `gorm:"default:false"`
	Continous           bool         `gorm:"default:false"`
	FutureOrder         bool         `gorm:"default:false"`
	Collected           bool         `gorm:"default:false"`
	ProcessTime         string       `gorm:"size:255"`
	ReasonForCancel     string       `gorm:"size:255"`
	OrderDate           time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	ContinousDate       *time.Time   `gorm:"column:continous_date"`
	LastUpdated         time.Time    `gorm:"autoUpdateTime"`
	CollectedDate       *time.Time   `gorm:"column:collected_date"`
	RejectedDate        *time.Time   `gorm:"column:rejected_date"`
	AcceptDate          *time.Time   `gorm:"column:accept_date"`
	AssignedToID        *uint        `gorm:"column:assigned_to_id"`
	AssignedTo          *User        `gorm:"foreignKey:AssignedToID"`
	Comments            string       `gorm:"type:text"`
	CommentedByID       *uint        `gorm:"column:commented_by_id"`
	CommentedBy         *User        `gorm:"foreignKey:CommentedByID"`
	PackageID           *uint        `gorm:"column:package_id"`
	OrderSetID          *uint        `gorm:"column:order_set_id"`
	Active              bool         `gorm:"default:true"`
	Duplicate           bool         `gorm:"default:false"`
	PaymentType         *string      `gorm:"size:255"`
	PayorDetailsID      *uint        `gorm:"column:payor_details_id"`
}

type OrderStatus struct {
	Status *string `gorm:"size:255"`
}

type OrderPeriod struct {
	StartDateTime *string `gorm:"size:255"`
	StopDateTime  *string `gorm:"size:255"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:255"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
}

type Medicine struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:500"`
	Code string `gorm:"size:500"`
}

type Panel struct {
	ID         uint   `gorm:"primaryKey"`
	PanelNamme string `gorm:"size:500"`
	Loinc      string `gorm:"size:500"`
	LoincNo    string `gorm:"size:500"`
}

type Imaging struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Code string `gorm:"size:255"`
}
