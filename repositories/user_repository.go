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


func GetUserByUsername(username string) models.User {
	user := models.User{} 
	models.DB.Find(&user, models.DB.Where("username = ?", username)) 
	return user 
}