package models

type Organisation struct {
	Name          string        `json:"name" binding:"required"`
	Email         string        `json:"email" binding:"required"`
	Password      string        `json:"password" binding:"required"`
	Description   string        `json:"description"`
	Logo          string        `json:"logo"`
	Website       string        `json:"website"`
	Causes        []string      `json:"causes"` // TODO: replace with a struct with predifined vals
	Opportunities []Opportunity `json:"opportunities"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}
