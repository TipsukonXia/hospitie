package dto

import (
	"encoding/json"
	"fmt"
	"hospitie/core"
	"hospitie/model"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateHospitieDTO(c echo.Context) model.Hospitie {

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

	cars := []model.Car{}
	fmt.Println(c.FormValue("cars"))
	err := json.Unmarshal([]byte(c.FormValue("cars")), &cars)
	if err != nil {
		core.ErrorHandler(err)
	}
	fmt.Println("cars")
	fmt.Println(cars)

	hospitie := model.Hospitie{
		Name:      c.FormValue("name"),
		Password:  password,
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		TelNo:     c.FormValue("tel_no"),
		Birthday:  &date,
		Email:     email,
		Cars:      cars,
	}
	return hospitie
}
