package dto

import (
	"time"
)

type UserResponse struct {
	ID               string    `json:"id,omitempty"`
	Username         string    `json:"username,omitempty"`
	Email            string    `json:"email,omitempty"`
	Password         string    `json:"password,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
	LastLogin        time.Time `json:"lastLogin,omitempty"`
}
