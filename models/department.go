package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type Department struct {
	ID    int     `orm:"column(id);pk"`
	Name  string  `orm:"column(name)"`
	Users []*User `orm:"reverse(many);null"`
}

func (u *Department) TableName() string {
	return "departments"
}

func init() {
	orm.RegisterModel(new(Department))
}
