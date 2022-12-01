package service

import (
	"hospitie/constant"
	"hospitie/core"
	"hospitie/model"
)

func GetAllUser() []model.User {
	var users []model.User
	res := constant.DB.Find(&users)
	println(res.RowsAffected)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return users
}

func FindUser(filter map[string]interface{}) []model.User {
	var users []model.User
	res := constant.DB.Where(filter).Find(&users)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return users
}

func CreateUser(user model.User) model.User {
	res := constant.DB.Create(&user)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}

func UpdateUser(user_id string, value map[string]interface{}) model.User {
	user := model.User{}
	constant.DB.Find(&user, user_id)
	res := constant.DB.Model(&user).Updates(value)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}

func DeleteUser(user_id string) model.User {
	user := model.User{}
	res := constant.DB.Delete(&user, user_id)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}
