package models

// TODO: abstract these models such like in https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/model/model.go
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
