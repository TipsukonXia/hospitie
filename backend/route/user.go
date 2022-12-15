package route

import (
	"encoding/json"
	"hospitie/constant"
	"hospitie/dto"
	"hospitie/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserRoute() {

	userGroup := constant.EchoInstance.Group("/user")
	us := service.NewUserService()

	userGroup.GET("", func(c echo.Context) error {
		res := us.GetAllUser(c)
		return c.JSON(http.StatusOK, res)
	})

	userGroup.POST("", func(c echo.Context) error {
		user := dto.CreateUserDTO(c)
		res := us.CreateUser(c, user)
		return c.JSON(http.StatusOK, res)
	})

	userGroup.POST("/filter", func(c echo.Context) error {
		body := map[string]interface{}{}
		// ?? how to cast req body to map[string]interface{}
		json.NewDecoder(c.Request().Body).Decode(&body)
		res := us.FindUser(c, body)
		return c.JSON(http.StatusOK, res)
	})

	userGroup.PATCH("/:user_id", func(c echo.Context) error {
		body := map[string]interface{}{}
		json.NewDecoder(c.Request().Body).Decode(&body)
		res := us.UpdateUser(c, c.Param("user_id"), body)
		return c.JSON(http.StatusOK, res)
	})

	userGroup.DELETE("/:user_id", func(c echo.Context) error {
		res := us.DeleteUser(c, c.Param("user_id"))
		return c.JSON(http.StatusOK, res)
	})

}
