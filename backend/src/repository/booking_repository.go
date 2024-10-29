package repository

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	CarNumber string
	Capacity  int
}

type Booking struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ReferenceNumber string
	StartTime       time.Time `gorm:"not null"`
	EndTime         time.Time `gorm:"not null"`
	Status          string
	Amount          float64
	Car             Car
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
