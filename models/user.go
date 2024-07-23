package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type User struct {
	ID         int         `orm:"column(id);pk"`
	Name       string      `orm:"column(name)"`
	Email      string      `orm:"column(email)"`
	Password   string      `orm:"column(password);null"`
	DeletedAt  string      `orm:"column(deleted_at);auto_now;type(datetime);null"`
	Department *Department `orm:"rel(fk);column(department_id)"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}
