package model

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	// ID           uint   `json:"id" gorm:"primaryKey"`
	LicensePlate string `json:"lisense_plate" gorm:"primaryKey"`
	Details      string `json:"details"`
	HospitieID   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	// OwnerID      uint
	// OwnerType    string
}
