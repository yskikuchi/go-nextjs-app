package model

import "github.com/google/uuid"

type Car struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	CarNumber string `gorm:"not null" json:"carNumber"`
	Capacity  int `gorm:"not null" json:"capacity"`
	PriceForInitialTwelveHours int `gorm:"not null" json:"priceForInitialTwelveHours"`
	PricePerAdditionalSixHours int `gorm:"not null" json:"pricePerAdditionalSixHours"`
}
