package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shift struct {
	ID                    primitive.ObjectID `bson:"_id"`
	Date                  string             `json:"date" bson:"date" binding:"required"`
	StartTime             string             `json:"start_time" bson:"start_time" binding:"required"`
	EndTime               string             `json:"end_time" bson:"end_time" binding:"required"`
	RegistrationCloseDate string             `json:"registration_close_date" bson:"registration_close_date" binding:"required"`
	AcceptedUsers         []string           `json:"accepted_users" bson:"accepted_users"`
	Vacancies             int16              `json:"vacancies" bson:"vacancies" binding:"required"`
	CreatedAt             string             `json:"created_at" bson:"created_at"`
	UpdatedAt             string             `json:"updated_at" bson:"updated_at"`
}
