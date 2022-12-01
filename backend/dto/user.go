package dto

import (
	"hospitie/core"
	"hospitie/model"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateUserDTO(c echo.Context) model.User {

	date, error := time.Parse("2006-01-02", c.FormValue("birthday"))
	if error != nil {
		core.ErrorHandler(error)
	}

	emailVal := c.FormValue("email")
	var email *string
	if emailVal == "" {
		email = nil
	} else {
		email = &emailVal
	}

	password := core.HashPassword(c.FormValue("password"))

	user := model.User{
		Name:      c.FormValue("name"),
		Password:  password,
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		TelNo:     c.FormValue("tel_no"),
		Birthday:  &date,
		Email:     email,
	}
	return user
}
