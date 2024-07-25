package controllers

import (
	"src/models"
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

	var users *[]models.User
	users, _ = this.userRepo.GetAll()

	this.ResponseOk(users)
}

func (this *UserController) GetUser() {
	var user *models.User
	user, _ = this.userRepo.FindById(1)

	this.ResponseOk(user)
}
