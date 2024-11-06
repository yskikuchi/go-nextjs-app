package repository

import (
	"log"

	"github.com/yskikuchi/go-nextjs-app/infra"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/service"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	db, err := infra.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	return &AdminRepository{
		DB: db,
	}
}

func (repo *AdminRepository) Create(admin *model.Admin) (model.Admin, error) {
	hashedPassword, err := service.EncryptPassword(admin.Password)
	if err != nil {
		return model.Admin{}, err
	}

	admin.Password = hashedPassword
	if err := repo.DB.Create(&admin).Error; err != nil {
		return model.Admin{}, err
	}

	return *admin, nil
}

func (repo *AdminRepository) FindByEmail(email string) (model.Admin, error) {
	var admin model.Admin
	if err := repo.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return model.Admin{}, err
	}

	return admin, nil
}
