package eav_models

type SavedObservationForm struct {
	ID          uint    `gorm:"primaryKey"`
	MrnID       string  `gorm:"column:mrn_id"`
	EncounterID string  `gorm:"column:encounter_id"`
	FormID      uint    `gorm:"column:form_id"`
	Form        HisForm `gorm:"foreignKey:FormID;references:ID"`
	// ValuesId    list.List `gorm:"column:form_id"`
	Values []Forms_savedobservationform_value `gorm:"many2many:Forms_savedobservationform_value"`
	// Values_data []Value `gorm:"many2many:Forms_savedobservationform_value;joinForeignKey:SavedObservationFormID;joinReferences:ValueID"`
	CreatedDate string `gorm:"column:created_date"`
}

type HisForm struct {
	ID         uint     `gorm:"primaryKey"`
	Name       string   `gorm:"column:name"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID uint     `gorm:"column:category_id"`
}

type Category struct {
	ID       uint   `gorm:"primaryKey"`
	Category string `gorm:"column:category"`
}

type Value struct {
	ID        uint   `gorm:"primaryKey"`
	Attribute string `gorm:"column:attribute"`
	ValueText string `gorm:"column:value_text"`
	Created   string `gorm:"column:created"`
}

type Attribute struct {
	ID        uint   `gorm:"primaryKey"`
	Attribute string `gorm:"column:attribute"`
	ValueText string `gorm:"column:value_text"`
	Created   string `gorm:"column:created"`
}

type ContentType struct {
	ID        uint `gorm:"primaryKey"`
	App_Label string
	Model     string
}

type Forms_savedobservationform_value struct {
	ID                      uint `gorm:"primaryKey"`
	Savedobservationform_id uint
	Value_id                uint
}

func (ContentType) TableName() string {
	return "django_content_type"
}

func (Attribute) TableName() string {
	return "eav_attribute"
}

func (Value) TableName() string {
	return "eav_value"
}

func (Category) TableName() string {
	return "Forms_category"
}

func (HisForm) TableName() string {
	return "Forms_his_form"
}

func (SavedObservationForm) TableName() string {
	return "Forms_savedobservationform"
}

func (Forms_savedobservationform_value) TableName() string {
	return "Forms_savedobservationform_value"
}
