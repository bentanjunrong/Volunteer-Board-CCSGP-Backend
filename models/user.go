package models

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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
	AcceptedOpportunities []AcceptedOpportunity `json:"accepted_opportunities" bson:"accepted_opportunities"`
	SMSNotification       bool                  `json:"sms_notification" bson:"sms_notification"`
	EmailNotification     bool                  `json:"email_notification" bson:"email_notification"`
	CreatedAt             string                `json:"created_at" bson:"created_at"`
	UpdatedAt             string                `json:"updated_at" bson:"updated_at"`
}

func (u *User) Create(user User) (User, error) {
	// Set up the request object.
	user.Password = utils.Hash(user.Password)
	body, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling user: %s", err)
	}
	req := esapi.IndexRequest{
		Index: "users",
		Body:  strings.NewReader(string(body)),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), db.GetDB())
	res.Body.Close()
	return user, err
}

func (u *User) Read(email string) (map[string]interface{}, error) {
	user, err := db.GetOneByField("users", map[string]string{
		"email": email,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Update(user User) (*User, error) {
	return u, nil
}

func (u *User) Delete(user User) (*User, error) {
	return u, nil
}
