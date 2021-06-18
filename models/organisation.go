package models

type Organisation struct {
	ID            string   `json:"organisation_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Description   string   `json:"description"`
	Logo          string   `json:"logo"`
	Website       string   `json:"website"`
	Causes        []string `json:"causes"`
	Opportunities []string `json:"opportunities"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}