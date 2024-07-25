package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"src/database"
	"src/models"
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
	results := new([]models.User)
	err := database.DB.Model(new(models.User)).Find(results).Error
	if err != nil {
		if !(errors.Is(err, gorm.ErrRecordNotFound)) {
			log.Printf("UserRepository.GetAll: Database error: %v", err)
			panic("Internal service error")
		}
		results = nil
	}
	return results
}
