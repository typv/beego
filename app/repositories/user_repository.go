package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"src/app/models"
	"src/database"
)

type IUserRepository interface {
	IRepository[models.User]
	GetAll() *[]models.User
}

type UserRepository struct {
	Repository[models.User]
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (ur UserRepository) GetAll() *[]models.User {
	results := make([]models.User, 0)
	err := database.DB.Model(&models.User{}).Find(&results).Error
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		log.Printf("UserRepository.GetAll: Database error: %v", err)
		panic("Internal service error")
	}

	return &results
}
