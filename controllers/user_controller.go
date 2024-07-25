package controllers

import (
	"src/repositories"
)

type UserController struct {
	BaseController
	userRepo repositories.UserRepository
}

func (this *UserController) GetUsers() {
	//Get auth user
	//authUser := this.Ctx.Input.GetData("authUser").(ultils.AuthUser)
	//fmt.Println(authUser.Email)

	users := this.userRepo.GetAll()
	this.ResponseOk(users)
}

func (this *UserController) GetUser() {
	user := this.userRepo.FindById(1)

	this.ResponseOk(user)
}
