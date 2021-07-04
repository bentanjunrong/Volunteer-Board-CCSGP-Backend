package models

import (
	"context"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: abstract these models such like in https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/model/model.go
type Opportunity struct {
	Name             string   `json:"name" binding:"required"`
	Description      string   `json:"description" binding:"required"`
	OrganisationName string   `json:"organisation_name" binding:"required"`
	AgeRequirement   int16    `json:"age_requirement" binding:"required"`
	Location         string   `json:"location" binding:"required"`
	Postin6gDate     string   `json:"posting_date" binding:"required"`
	Shifts           []Shift  `json:"shifts"  binding:"required"` // TODO: this validation not working. fix here: https://stackoverflow.com/questions/58585078/binding-validations-does-not-work-when-request-body-is-array-of-objects
	Causes           []string `json:"causes"`
	IsApproved       bool     `json:"is_approved"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}

func (o *Opportunity) Create(ctx context.Context, opp Opportunity) (*mongo.InsertOneResult, error) {
	result, err := db.GetCollection("opps").InsertOne(ctx, opp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *Opportunity) GetAll(ctx context.Context) ([]bson.M, error) {
	cursor, err := db.GetCollection("opps").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var opps []bson.M
	if err = cursor.All(ctx, &opps); err != nil {
		return nil, err
	}
	return opps, nil
}

func (o *Opportunity) Search(query string) ([]map[string]interface{}, error) {
	allOpps, err := db.Search(query, "opps")
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}
	for _, obj := range allOpps {
		opp := (obj["_source"]).(map[string]interface{})
		opp["id"] = obj["_id"]
		res = append(res, opp)
	}
	return res, nil
}

func (o *Opportunity) GetOne(id string) (map[string]interface{}, error) {
	allOpps, err := db.GetAllByField("opps", map[string]string{"_id": id})
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}
	for _, obj := range allOpps {
		opp := (obj["_source"]).(map[string]interface{})
		opp["id"] = obj["_id"]
		res = append(res, opp)
	}
	return res[0], nil
}
