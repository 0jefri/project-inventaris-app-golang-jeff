package model

import (
	"time"
)

type User struct {
	ID               string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Username         string    `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required,alphanum"`
	Email            string    `gorm:"type:varchar(255);not null;unique" json:"email,omitempty" binding:"required,email"`
	Password         string    `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required,alphanum"`
	RegistrationDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"registrationDate" binding:"omitempty"`
	LastLogin        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastLogin" binding:"omitempty"`
}
