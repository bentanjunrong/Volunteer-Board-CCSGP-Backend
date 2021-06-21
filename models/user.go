package models

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
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

func (u *User) Create(user User) error {
	// Set up the request object.
	body, err := json.Marshal(user)
	req := esapi.IndexRequest{
		Index: "users",
		Body:  strings.NewReader(string(body)),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), db.GetDB())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	res.Body.Close()
	return nil
}

func (u *User) Update(user User) (*User, error) {
	return u, nil
}

func (u *User) Delete(user User) (*User, error) {
	return u, nil
}
