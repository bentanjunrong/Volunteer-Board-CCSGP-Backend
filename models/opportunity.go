package models

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// TODO: abstract these models such like in https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/model/model.go
type Opportunity struct {
	Name             string   `json:"name" binding:"required"`
	Description      string   `json:"description" binding:"required"`
	OrganisationName string   `json:"organisation_name" binding:"required"`
	AgeRequirement   string   `json:"age_requirement" binding:"required"`
	Location         string   `json:"location" binding:"required"`
	PostingDate      string   `json:"posting_date" binding:"required"`
	Shifts           []Shift  `json:"shifts"`
	Causes           []string `json:"causes"`
	IsApproved       bool     `json:"is_approved"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}

func (o *Opportunity) Create(opp Opportunity) (Opportunity, error) {
	// Set up the request object.
	body, err := json.Marshal(opp)
	if err != nil {
		log.Fatalf("Error marshalling opportunity: %s", err)
	}
	req := esapi.IndexRequest{
		Index: "opps",
		Body:  strings.NewReader(string(body)),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), db.GetDB())
	res.Body.Close()
	return opp, err
}
