package repositories

import (
	"github.com/beego/beego/v2/client/orm"
	"src/models"
)

type IUserRepository interface {
	IRepository[models.User]
	GetAll() (*[]models.User, error)
}

type UserRepository struct {
	Repository[models.User]
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (this UserRepository) GetAll() (*[]models.User, error) {
	o := orm.NewOrm()
	var results []models.User
	_, err := o.QueryTable(new(models.User)).All(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
