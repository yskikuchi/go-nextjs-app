package repository

import (
	"errors"
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

// statusを指定して予約を取得
func (repo *BookingRepository) FindAll(status []string) ([]model.Booking, error) {
	bookings := []model.Booking{}

	var err error
	if status == nil {
		err = repo.DB.Preload("Car").Find(&bookings).Error
	} else {
		err = repo.DB.Preload("Car").Where("status IN ?", status).Find(&bookings).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return bookings, nil
}

func (repo *BookingRepository) FindByID(id string) (model.Booking, error) {
	booking := model.Booking{}

	if err := repo.DB.Preload("Car").Where("id = ?", id).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Booking{}, err
		}
	}

	return booking, nil
}

func (repo *BookingRepository) FindByReferenceNumber(referenceNumber string) (model.Booking, error) {
	booking := model.Booking{}

	if err := repo.DB.Preload("Car").Where("reference_number = ?", referenceNumber).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Booking{}, err
		}
	}

	return booking, nil
}

func (repo *BookingRepository) FindByReferenceNumberAndEmail(referenceNumber string, email string) (model.Booking, error) {
	booking := model.Booking{}

	if err := repo.DB.Preload("Car").Where("reference_number = ? AND user_email = ?", referenceNumber, email).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Booking{}, err
		}
	}

	return booking, nil
}

// 既存の予約と重複しないかの検証に使用
func (repo *BookingRepository) FindByCarIDAndTimeRange(carID string, startTime string, endTime string) (model.Booking, error) {
	booking := model.Booking{}

	if err := repo.DB.Where("car_id = ? AND start_time < ? AND end_time > ?", carID, endTime, startTime).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Booking{}, err
		}
	}

	return booking, nil
}

func (repo *BookingRepository) Create(booking *model.Booking) (model.Booking, error) {
	if err := repo.DB.Create(&booking).Error; err != nil {
		return model.Booking{}, err
	}

	return *booking, nil
}

func (repo *BookingRepository) Update(booking model.Booking) (model.Booking, error) {
	if err := repo.DB.Save(&booking).Error; err != nil {
		return model.Booking{}, err
	}
	return booking, nil
}
