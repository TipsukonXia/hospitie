package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"hospitie/constant"
	"hospitie/core"
	"hospitie/model"
	"hospitie/route"
)

func main() {
	godotenv.Load()

	/* create DB */
	DB := constant.DBInstance()
	if DB != nil {
		core.ErrorHandler(DB.Error)
	}

	/* migrate DB */
	err := model.Migrate()
	if err != nil {
		core.ErrorHandler(err)
	}

	/* init echo instant*/
	constant.InitEchoInstance()

	/*init route*/
	route.Init()

	constant.EchoInstance.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!!!")
	})
	constant.EchoInstance.Logger.Fatal(constant.EchoInstance.Start(":8080"))
}
