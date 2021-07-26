package models

import (
	"context"
	"time"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Organisation struct {
	ID                  primitive.ObjectID `json:"_id" bson:"_id"`
	Name                string             `json:"name" binding:"required"`
	Email               string             `json:"email" binding:"required"`
	Password            string             `json:"password" binding:"required"`
	Description         string             `json:"description"`
	Logo                string             `json:"logo"`
	Website             string             `json:"website"`
	Causes              []string           `json:"causes"` // TODO: replace with a struct with predifined vals
	ListedOpportunities []string           `json:"listed_opps"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
}

func (o *Organisation) GetOpps(orgID string) ([]bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	org := &Organisation{}
	if err = db.GetCollection("orgs").FindOne(ctx, bson.M{"_id": objID}).Decode(&org); err != nil {
		return nil, err
	}
	oppIDs := bson.A{}
	for _, listedOpp := range org.ListedOpportunities {
		oppId, err := primitive.ObjectIDFromHex(listedOpp)
		if err != nil {
			return nil, err
		}
		oppIDs = append(oppIDs, oppId)
	}
	cursor, err := db.GetCollection("opps").Find(
		ctx,
		bson.M{"_id": bson.M{"$in": oppIDs}},
	)
	if err != nil {
		return nil, err
	}
	opps := []bson.M{}
	if err = cursor.All(ctx, &opps); err != nil {
		return nil, err
	}
	return opps, nil
}

func (o *Organisation) Update(orgID string, orgUpdate Organisation) (Organisation, error) {
	orgId, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return Organisation{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	org := &Organisation{}
	if err = db.GetCollection("orgs").FindOne(ctx, bson.M{"_id": orgId}).Decode(&org); err != nil {
		return Organisation{}, err
	}

	if err := copier.Copy(org, orgUpdate); err != nil {
		return Organisation{}, err
	}

	returnUpdatedDoc := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnUpdatedDoc,
	}
	err = db.GetCollection("orgs").FindOneAndUpdate(ctx, bson.M{"_id": orgId}, bson.M{"$set": org}, opts).Decode(&org)

	return *org, nil
}
