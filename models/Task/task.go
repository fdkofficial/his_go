package task_models

type Task struct {
	// gorm.Model
	Id                  int             `json:"id"`
	Task_description    *string         `json:"task_description"`
	Mnemoric            *string         `json:"mnemoric"`
	Priority            *string         `json:"priority"`
	Task_category       *string         `json:"task_category"`
	Reference_set       *string         `json:"reference_set"`
	Status              *string         `json:"status"`
	Start               *string         `json:"start"`
	End                 *string         `json:"end"`
	Patient_id          *string         `json:"patient_id"`
	Task_id_fk          int             `gorm:"column:task_id"` // Foreign key column `json:"task_id_fk"`
	Task                Tasks_categorie `gorm:"foreignKey:Task_id_fk;references:Id"`
	User_id             *string         `json:"user_id"`
	Encounter_id        *string         `json:"encounter_id"`
	Frequency           *string         `json:"frequency"`
	Comment             *string         `json:"comment"`
	Is_drafted          bool            `json:"is_drafted"`
	Continous           *string         `json:"continous"`
	Indications         *string         `json:"indications"`
	Monitor_instruction *string         `json:"monitor_instruction"`
	Order_description   *string         `json:"order_description"`
	Order_no            *string         `json:"order_no"`
	Special_description *string         `json:"special_description"`
	Fhir_id             *string         `json:"fhir_id"`
}

type Tasks_categorie struct {
	Id             int
	Code           *string
	Name           *string
	Category       *string
	Order_category *string
	Sub_category   *string
}

func (Task) TableName() string {
	return "Tasks_task"
}

func (Tasks_categorie) TableName() string {
	return "Tasks_taskcategorie"
}
