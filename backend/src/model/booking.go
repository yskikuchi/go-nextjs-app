package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	ReferenceNumber string    `gorm:"not null" json:"referenceNumber"`
	StartTime       time.Time `gorm:"not null" json:"startTime" validate:"required,mustBeAfterNow"`
	EndTime         time.Time `gorm:"not null" json:"endTime" validate:"required,gtfield=StartTime"`
	Status          string    `gorm:"not null;default:'unconfirmed'" json:"status"`
	Amount          float64   `gorm:"not null" json:"amount"`
	UserName        string    `gorm:"not null" json:"userName"`
	UserEmail       string    `gorm:"not null" json:"userEmail"`
	CarID           uuid.UUID `json:"carId"`
	Car             Car       `gorm:"foreignKey:CarID" json:"car"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
