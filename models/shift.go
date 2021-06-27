package models

type Shift struct {
	Date                  string   `json:"date" binding:"required"`
	StartTime             string   `json:"start_time" binding:"required"`
	EndTime               string   `json:"end_time" binding:"required"`
	RegistrationCloseDate string   `json:"registration_close_date" binding:"required"`
	AcceptedUsers         []string `json:"accepted_users"`
	Vacancies             int16    `json:"vacancies" binding:"required"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
}
