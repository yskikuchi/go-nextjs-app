package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/yskikuchi/go-nextjs-app/infra"
	"github.com/yskikuchi/go-nextjs-app/model"
	"gorm.io/gorm"
)

type BookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository() *BookingRepository {
	db, err := infra.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	return &BookingRepository{
		DB: db,
	}
}

func (repo *BookingRepository) FindAll() ([]model.Booking, error) {
	bookings := []model.Booking{}

	if err := repo.DB.Find(&bookings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return bookings, nil
}

func (repo *BookingRepository) FindByReferenceNumber(referenceNumber string) (model.Booking, error) {
	booking := model.Booking{}

	if err := repo.DB.Where("reference_number = ?", referenceNumber).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Booking{}, err
		}
	}

	return booking, nil
}

func (repo *BookingRepository) Create(booking *model.Booking) (model.Booking, error) {
	fmt.Println("リポジトリ", booking)
	if err := repo.DB.Create(&booking).Error; err != nil {
		return model.Booking{}, err
	}

	return *booking, nil
}
