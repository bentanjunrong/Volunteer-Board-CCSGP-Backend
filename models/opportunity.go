package models

import (
	"context"
	"time"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: abstract these models such like in https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/model/model.go
type Opportunity struct {
	Name             string   `json:"name" bson:"name" binding:"required"`
	Description      string   `json:"description" bson:"description" binding:"required"`
	OrganisationName string   `json:"organisation_name" bson:"organisation_name" binding:"required"`
	AgeRequirement   int16    `json:"age_requirement" bson:"age_requirement" binding:"required"`
	Location         string   `json:"location" bson:"location" binding:"required"`
	PostingDate      string   `json:"posting_date" bson:"posting_date" binding:"required"`
	Shifts           []Shift  `json:"shifts"  bson:"shifts" binding:"required"` // TODO: this validation not working. fix here: https://stackoverflow.com/questions/58585078/binding-validations-does-not-work-when-request-body-is-array-of-objects
	Causes           []string `json:"causes" bson:"causes" `
	IsApproved       bool     `json:"is_approved" bson:"is_approved" `
	CreatedAt        string   `json:"created_at" bson:"created_at" `
	UpdatedAt        string   `json:"updated_at" bson:"updated_at" `
}

func (o *Opportunity) Create(ctx context.Context, opp Opportunity) (interface{}, error) {
	for i := 0; i < len(opp.Shifts); i++ {
		opp.Shifts[i].ID = primitive.NewObjectID()
	}
	result, err := db.GetCollection("opps").InsertOne(ctx, opp)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}

func (o *Opportunity) GetAll() ([]bson.M, error) {
	projection := bson.M{
		"description":     0,
		"shifts":          0,
		"age_requirement": 0,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.GetCollection("opps").Find(
		ctx,
		bson.M{"is_approved": true},
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

func (o *Opportunity) GetAllUnapproved() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.GetCollection("opps").Find(
		ctx,
		bson.M{"is_approved": false},
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

func (o *Opportunity) CreateShift(ctx context.Context, id string, shift Shift) error {
	shift.ID = primitive.NewObjectID()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = db.GetCollection("opps").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$push": bson.M{"shifts": shift},
		},
	)
	return err
}

// TODO: improve this garbage function.
// TODO: return error if shift not found.
func (o *Opportunity) DeleteShift(oppID string, shiftID string) error {
	oppId, err := primitive.ObjectIDFromHex(oppID)
	if err != nil {
		return err
	}
	shiftId, err := primitive.ObjectIDFromHex(shiftID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opp := &Opportunity{}
	if err = db.GetCollection("opps").FindOne(ctx, bson.M{"_id": oppId}).Decode(&opp); err != nil {
		return err
	}

	for _, shift := range opp.Shifts {
		// find shift
		if shift.ID == shiftId {
			// for each accepted user, remove shiftId from their accepted_opps
			for _, acceptedUserId := range shift.AcceptedUsers {
				userId, err := primitive.ObjectIDFromHex(acceptedUserId)
				if err != nil {
					return err
				}
				user := &User{}
				if err = db.GetCollection("users").FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
					return err
				}
				for i := 0; i < len(user.AcceptedOpportunities); i++ {
					if user.AcceptedOpportunities[i].OppID == oppID {
						acceptedOpp := user.AcceptedOpportunities[i]
						for j := 0; j < len(acceptedOpp.ShiftIDs); j++ {
							// remove shiftId
							if acceptedOpp.ShiftIDs[j] == shiftID {
								user.AcceptedOpportunities[i].ShiftIDs = append(acceptedOpp.ShiftIDs[:j], acceptedOpp.ShiftIDs[j+1:]...)
								break
							}
						}
						break
					}
				}
				// update user
				if _, err = db.GetCollection("users").ReplaceOne(ctx, bson.M{"_id": userId}, user); err != nil {
					return err
				}
			}
		}
	}

	// delete shift
	for i := 0; i < len(opp.Shifts); i++ {
		if opp.Shifts[i].ID == shiftId {
			opp.Shifts = append(opp.Shifts[:i], opp.Shifts[i+1:]...)
			break
		}
	}

	return nil
}
