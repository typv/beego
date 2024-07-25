package controllers

import (
	"src/repositories"
)

type UserController struct {
	BaseController
	userRepo repositories.UserRepository
}

func (uc *UserController) GetUsers() {
	//Get auth user
	//authUser := uc.Ctx.Input.GetData("authUser").(ultils.AuthUser)
	//fmt.Println(authUser.Email)

	users := uc.userRepo.GetAll()
	uc.ResponseOk(users)
}

func (uc *UserController) GetUser() {
	user := uc.userRepo.FindById(1)

	uc.ResponseOk(user)
}
