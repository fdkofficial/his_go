package procedure_models

type ProcedureInfo struct {
	Id                              int     `json:"id"`
	VisitNo                         *string `json:"visit_no"`
	VisitDate                       *string `json:"visit_date"`
	Mrn                             *string `json:"mrn"`
	Encounter                       *string `json:"encounter"`
	FhirId                          *string `json:"fhir_id"`
	ProcedureScheme                 *string `json:"procedure_scheme"`
	ProcedureCode                   *string `json:"procedure_code"`
	Priority                        *string `json:"priority"`
	OrderNo                         *string `json:"order_no"`
	ProcedureType                   *string `json:"procedure_type"`
	ProcedureName                   *string `json:"procedure_name"`
	ProcedureResult                 *string `json:"procedure_result"`
	ProcedureReason                 *string `json:"procedure_reason"`
	ProphylacticAntibioticIndicated bool    `json:"prophylactic_antibiotic_indicated"`
	RemovalDate                     *string `json:"removal_date"`
	RemovalTime                     *string `json:"removal_time"`
	Options                         *string `json:"options"`
	StartDate                       *string `json:"start_date"`
	StartTime                       *string `json:"start_time"`
	EndDate                         *string `json:"end_date"`
	EndTime                         *string `json:"end_time"`
	BodySite                        *string `json:"body_site"`
	AnesthesiaType                  *string `json:"anesthesia_type"`
	ProcedureBy                     *string `json:"procedure_by"`
	Surgeon                         *string `json:"surgeon"`
	Status                          *string `json:"status"`
}
