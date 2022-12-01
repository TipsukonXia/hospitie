package model

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	LicensePlate string   `json:"lisense_plate" gorm:"not null;unique;index"`
	HospitieID   uint     `json:"hospitie_id"`
	Hospitie     Hospitie `json:"hospitie" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Details      string   `json:"details"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
