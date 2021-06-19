package models

type User struct {
	ID                string   `json:"user_id"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Password          string   `json:"password"`
	DateOfBirth       string   `json:"date_of_birth"`
	Gender            string   `json:"gender"`
	Age               int16    `json:"age"`
	Availability      []string `json:"availability"`
	SMSNotification   bool     `json:"sms_notification"`
	EmailNotification bool     `json:"email_notification"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
}

func (u User) Signup() *User {
	return &u
}
