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

func UpdateUser(userDto dto.UserDto) error {
	user := models.User{} 
	if userDto.Username != "" {
		user.Username = userDto.Username 
	}

	if userDto.Password != "" {
		user.Password = userDto.Password
	}

	err := models.DB.Model(&user).Update(&user).Error 
	return err 
}

func DeleteUserById(id uint) error {
	err := models.DB.Delete(&models.User{ID: id}).Error
	return err
}