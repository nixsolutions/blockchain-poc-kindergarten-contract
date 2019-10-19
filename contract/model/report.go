package model

const CREATED_STATUS = "created"
const APPROVED_STATUS = "approved"

type Report struct {
	Id     string     `json:"id"`
	Status string     `json:"status"`
	CardId string     `json:"card_id"`
	Data   ReportData `json:"report_data,omitempty"`
}

type ReportData struct {
	Name        string            `json:"name"`
	Surname     string            `json:"surname"`
	Vaccination []VaccinationItem `json:"vaccination"`
}

type VaccinationItem struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}
