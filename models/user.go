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

type User struct {
	Name              string   `json:"name" binding:"required"`
	Email             string   `json:"email" binding:"required"`
	Password          string   `json:"password" binding:"required"`
	DateOfBirth       string   `json:"date_of_birth" binding:"required"`
	Gender            string   `json:"gender" binding:"required"`
	Age               int16    `json:"age" binding:"required"`
	Availability      []string `json:"availability"`
	SMSNotification   bool     `json:"sms_notification"`
	EmailNotification bool     `json:"email_notification"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
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

func (u *User) Update(user User) (*User, error) {
	return u, nil
}

func (u *User) Delete(user User) (*User, error) {
	return u, nil
}
