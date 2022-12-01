package model

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	UserID       uint     `json:"user_id"`
	User         User     `json:"user" gorm:"not null"`
	HospitieID   uint     `json:"hospitie_id"`
	Hospitie     Hospitie `json:"hospitie" gorm:"not null"`
	CarID        uint     `json:"car_id"`
	Car          Car      `json:"car" gorm:"not null"`
	IsCanceled   bool     `json:"is_canceled" gorm:"not null;default:false"`
	CancelDetail *string  `json:"cancel_detail"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
