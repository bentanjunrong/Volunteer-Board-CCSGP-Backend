package models

import (
	"context"
	"time"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcceptedOpportunity struct {
	OppID    string   `json:"opp_id" bson:"opp_id" binding:"required"`
	ShiftIDs []string `json:"shift_ids" bson:"shift_ids" binding:"required"`
}

type User struct {
	Name                  string                `json:"name" bson:"name" binding:"required"`
	Email                 string                `json:"email" bson:"email" binding:"required"`
	Password              string                `json:"password" bson:"password" binding:"required"`
	DateOfBirth           string                `json:"date_of_birth" bson:"date_of_birth"`
	Gender                string                `json:"gender" bson:"gender"`
	Age                   int16                 `json:"age" bson:"age"`
	Availability          []string              `json:"availability" bson:"availability"`
	AcceptedOpportunities []AcceptedOpportunity `json:"accepted_opps" bson:"accepted_opps"`
	SMSNotification       bool                  `json:"sms_notification" bson:"sms_notification"`
	EmailNotification     bool                  `json:"email_notification" bson:"email_notification"`
	CreatedAt             string                `json:"created_at" bson:"created_at"`
	UpdatedAt             string                `json:"updated_at" bson:"updated_at"`
}

func (u *User) GetOpps(userID string) ([]bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := &User{}
	if err = db.GetCollection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user); err != nil {
		return nil, err
	}
	var oppIDs []primitive.ObjectID
	for _, acceptedOpp := range user.AcceptedOpportunities {
		oppId, err := primitive.ObjectIDFromHex(acceptedOpp.OppID)
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
	var opps []bson.M
	if err = cursor.All(ctx, &opps); err != nil {
		return nil, err
	}
	return opps, nil
}
