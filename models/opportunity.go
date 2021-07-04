package models

import (
	"context"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: abstract these models such like in https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/model/model.go
type Opportunity struct {
	Name             string   `json:"name" bson:"name,omitempty" binding:"required"`
	Description      string   `json:"description" bson:"description,omitempty" binding:"required"`
	OrganisationName string   `json:"organisation_name" bson:"organisation_name,omitempty" binding:"required"`
	AgeRequirement   int16    `json:"age_requirement" bson:"age_requirement,omitempty" binding:"required"`
	Location         string   `json:"location" bson:"location,omitempty" binding:"required"`
	PostingDate      string   `json:"posting_date" bson:"posting_date,omitempty" binding:"required"`
	Shifts           []Shift  `json:"shifts"  bson:"shifts,omitempty" binding:"required"` // TODO: this validation not working. fix here: https://stackoverflow.com/questions/58585078/binding-validations-does-not-work-when-request-body-is-array-of-objects
	Causes           []string `json:"causes" bson:"causes" `
	IsApproved       bool     `json:"is_approved" bson:"is_approved" `
	CreatedAt        string   `json:"created_at" bson:"created_at" `
	UpdatedAt        string   `json:"updated_at" bson:"updated_at" `
}

func (o *Opportunity) Create(ctx context.Context, opp Opportunity) (interface{}, error) {
	result, err := db.GetCollection("opps").InsertOne(ctx, opp)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}

func (o *Opportunity) GetAll(ctx context.Context) ([]bson.M, error) {
	projection := bson.M{
		"description":     0,
		"shifts":          0,
		"age_requirement": 0,
	}
	cursor, err := db.GetCollection("opps").Find(
		ctx,
		bson.M{},
		options.Find().SetProjection(projection),
	)
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

func (o *Opportunity) GetOne(ctx context.Context, id string) (bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var opp bson.M
	if err = db.GetCollection("opps").FindOne(ctx, bson.M{"_id": objID}).Decode(&opp); err != nil {
		return nil, err
	}
	return opp, nil
}
