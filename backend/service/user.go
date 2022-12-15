package service

import (
	"hospitie/constant"
	"hospitie/core"
	"hospitie/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService() *UserService {
	db := constant.DBInstance()
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetAllUser(c echo.Context) []model.User {
	var users []model.User
	res := s.DB.Find(&users)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return users
}

func (s *UserService) FindUser(c echo.Context, filter map[string]interface{}) []model.User {
	var users []model.User
	res := s.DB.Where(filter).Find(&users)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return users
}

func (s *UserService) CreateUser(c echo.Context, user model.User) model.User {
	res := s.DB.Create(&user)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}

func (s *UserService) UpdateUser(c echo.Context, user_id string, value map[string]interface{}) model.User {
	user := model.User{}
	s.DB.Find(&user, user_id)
	res := s.DB.Model(&user).Updates(value)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}

func (s *UserService) DeleteUser(c echo.Context, user_id string) model.User {
	user := model.User{}
	res := s.DB.Delete(&user, user_id)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return user
}
