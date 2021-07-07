package models

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

func (u *User) Update(user User) (*User, error) {
	return u, nil
}

func (u *User) Delete(user User) (*User, error) {
	return u, nil
}
