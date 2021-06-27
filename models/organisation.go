package models

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

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

func (o *Organisation) Create(org Organisation) (Organisation, error) {
	// Set up the request object.
	org.Password = utils.Hash(org.Password)
	body, err := json.Marshal(org)
	if err != nil {
		log.Fatalf("Error marshalling organisation: %s", err)
	}
	req := esapi.IndexRequest{
		Index: "orgs",
		Body:  strings.NewReader(string(body)),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), db.GetDB())
	res.Body.Close()
	return org, err
}

func (o *Organisation) Read(email string) (map[string]interface{}, error) {
	org, err := db.GetOneByField("orgs", map[string]string{
		"email": email,
	})
	if err != nil {
		return nil, err
	}
	return org, nil
}
