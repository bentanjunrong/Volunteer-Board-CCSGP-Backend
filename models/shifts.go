package models

type Shift struct {
	ID                    string   `json:"opp_id"`
	Date                  string   `json:"date"`
	StartTime             string   `json:"start_time"`
	EndTime               string   `json:"end_time"`
	RegistrationCloseDate string   `json:"registration_close_date"`
	AcceptedUsers         []string `json:"accepted_users"`
	Vacancies             int16    `json:"vacancies"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
}