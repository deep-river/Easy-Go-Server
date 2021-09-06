package user

import (
	"EasyGo-Server/dto"
	"EasyGo-Server/models"
)

func CreateUser(dto dto.UserDto) error {
	user := models.User{}
	user.Username = dto.Username
	user.Password = dto.Password

	err := models.DB.Create(&user).Error
	return err
}
