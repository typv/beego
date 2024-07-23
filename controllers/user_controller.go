package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"src/models"
)

type UserController struct {
	web.Controller
}

func (c *UserController) GetUser() {
	o := orm.NewOrm()
	u := &models.User{ID: 1}
	err := o.Read(u)
	fmt.Println(err)
	if err != nil {
		c.Ctx.WriteString("User not found")
		return
	}
	o.LoadRelated(u, "Department")

	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) GetUsers() {
	var users []models.User
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.User))

	_, err := qs.
		RelatedSel("Department").
		Filter("id", 1).
		All(&users, "ID", "Name", "Email")
	if err != nil {
		c.Ctx.WriteString("Error: " + err.Error())
		return
	}

	c.Data["json"] = users
	c.ServeJSON()
}
