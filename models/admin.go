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

type Admin struct{}

func (a *Admin) Approve(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	opp := &Opportunity{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.GetCollection("opps").FindOne(ctx, bson.M{"_id": objID}).Decode(&opp); err != nil {
		return err
	}

	opp.Status = "approved"
	if _, err = db.GetCollection("opps").ReplaceOne(ctx, bson.M{"_id": objID}, opp); err != nil {
		return err
	}

	return nil
}

func (a *Admin) Reject(id string, rejReason string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	opp := &Opportunity{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.GetCollection("opps").FindOne(ctx, bson.M{"_id": objID}).Decode(&opp); err != nil {
		return err
	}

	opp.Status = "rejected"
	opp.RejectionReason = rejReason
	if _, err = db.GetCollection("opps").ReplaceOne(ctx, bson.M{"_id": objID}, opp); err != nil {
		return err
	}

	return nil
}

func (a *Admin) Update(adminID string, adminUpdate Admin) (Admin, error) {
	adminId, err := primitive.ObjectIDFromHex(adminID)
	if err != nil {
		return Admin{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	admin := &Admin{}
	if err = db.GetCollection("admins").FindOne(ctx, bson.M{"_id": adminId}).Decode(&admin); err != nil {
		return Admin{}, err
	}

	if err := copier.Copy(admin, adminUpdate); err != nil {
		return Admin{}, err
	}

	returnUpdatedDoc := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnUpdatedDoc,
	}
	err = db.GetCollection("admins").FindOneAndUpdate(ctx, bson.M{"_id": adminId}, bson.M{"$set": admin}, opts).Decode(&admin)

	return *admin, nil
}
