package controllers

import (
	"github.com/IBM/sarama"
	"log"
	"src/app/repositories"
	"src/app/services"
)

type UserController struct {
	BaseController
	userRepo repositories.UserRepository
}

func (uc *UserController) GetUsers() {
	//Get auth user
	//authUser := uc.Ctx.Input.GetData("authUser").(ultils.AuthUser)
	//fmt.Println(authUser.Email)

	key := "test"
	value := "Hello, World!"
	message := &sarama.ProducerMessage{
		Topic: "hello",
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	partition, offset, err := services.Producer.SendMessage(message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	} else {
		log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}

	users := uc.userRepo.GetAll()
	uc.ResponseOk(users)
}

func (uc *UserController) GetUser() {
	user := uc.userRepo.FindById(1)

	uc.ResponseOk(user)
}
