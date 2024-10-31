package repository

import (
	"log"

	"github.com/yskikuchi/go-nextjs-app/infra"
	"github.com/yskikuchi/go-nextjs-app/model"
	"gorm.io/gorm"
)

type CarRepository struct {
	DB *gorm.DB
}

func NewCarRepository() *CarRepository {
	db, err := infra.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	return &CarRepository{
		DB: db,
	}
}

func (repo *CarRepository) FindAll() ([]model.Car, error) {
	cars := []model.Car{}

	if err := repo.DB.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (repo *CarRepository) Create(car *model.Car) (model.Car, error) {
	if err := repo.DB.Create(&car).Error; err != nil {
		return model.Car{}, err
	}

	return *car, nil
}
