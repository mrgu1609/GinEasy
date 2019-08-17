package service

import (
	"gineasy/models"
	"gineasy/views"
)

type SignMsg struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	InviteCode string `json:"invite_code" binding:"required"`
}

func (service *SignMsg) valid() *views.Response {
	return nil
}

func (service *SignMsg) Register() *views.Response {
	user := models.User{
		Username: service.Username,
		Password: service.Username,
	}
	//
	if err := models.DB.Create(&user).Error; err != nil{
		return views.ErrorRes("register failed")
	}
	return nil
}
