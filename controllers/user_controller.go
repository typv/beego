package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"src/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) GetUsers() {
	var users []models.User
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.User))

	_, err := qs.
		RelatedSel("Department").
		Filter("id", 1).
		All(&users, "ID", "Name", "Email")
	if err != nil {
		this.ResponseInternalServerError(nil, nil)
		return
	}

	this.ResponseOk(users)
}

func (this *UserController) GetUser() {
	o := orm.NewOrm()
	u := &models.User{ID: 1}
	err := o.Read(u)
	if err != nil {
		this.ResponseNotFound("User not found", nil)
		return
	}
	o.LoadRelated(u, "Department")

	this.ResponseOk(u)
}
