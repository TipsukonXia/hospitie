package model

import (
	"time"

	"gorm.io/gorm"
)

type Hospitie struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	FirstName string     `json:"first_name" gorm:"not null"`
	LastName  string     `json:"last_name" gorm:"not null"`
	Password  string     `json:"password" gorm:"not null"`
	TelNo     string     `json:"tel_no" gorm:"size:15;not null;unique"`
	Birthday  *time.Time `json:"birthday" gorm:"not null"`
	Email     *string    `json:"email" gorm:"unique"`
	Cars      []Car      `json:"cars" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
