package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"src/database"
)

type IRepository[Model any] interface {
	FindById(id int) *Model
}

type Repository[Model any] struct{}

func NewRepository[Model any]() IRepository[Model] {
	return &Repository[Model]{}
}

func (r *Repository[Model]) FindById(id int) *Model {
	result := new(Model)
	err := database.DB.Model(result).Where("id = ?", id).First(result).Error
	if err != nil {
		if !(errors.Is(err, gorm.ErrRecordNotFound)) {
			log.Printf("Repository.FindById: Database error: %v", err)
			panic("Internal service error")
		}
		result = nil
	}

	return result
}
