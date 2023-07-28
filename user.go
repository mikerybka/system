package system

import "strings"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (u *User) SendSMS(message string) error {
	return SendSMS(u.Phone, message)
}

func (u *User) SendEmail(subject, message string) error {
	return SendEmail(u.Email, u.FullName(), subject, message, message)
}

func (u *User) FullName() string {
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}
