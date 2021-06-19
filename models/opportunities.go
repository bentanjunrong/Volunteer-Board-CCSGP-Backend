package models

type Opportunity struct {
	ID               string   `json:"opp_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	OrganisationName string   `json:"organisation_name"`
	AgeRequirement   string   `json:"age_requirement"`
	Location         string   `json:"location"`
	PostingDate      string   `json:"posting_date"`
	Shifts           []Shift  `json:"shifts"`
	Causes           []string `json:"causes"`
	IsApproved       bool     `json:"is_approved"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}
