package repositories

import (
	"github.com/beego/beego/v2/client/orm"
)

type IRepository[Model any] interface {
	FindById(id int) (*Model, error)
}

type Repository[Model any] struct{}

func NewRepository[Model any]() IRepository[Model] {
	return &Repository[Model]{}
}

func (r *Repository[Model]) FindById(id int) (*Model, error) {
	o := orm.NewOrm()
	result := new(Model)
	err := o.QueryTable(result).Filter("ID", id).One(result)

	return result, err
}
