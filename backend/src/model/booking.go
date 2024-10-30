package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	ReferenceNumber string    `gorm:"not null" json:"referenceNumber"`
	StartTime       time.Time `gorm:"not null" json:"startTime"`
	EndTime         time.Time `gorm:"not null" json:"endTime"`
	Status          string    `gorm:"not null" json:"status"`
	Amount          float64   `gorm:"not null" json:"amount"`
	CarID           uuid.UUID `json:"carId"`
	Car             Car       `gorm:"foreignKey:CarID"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
