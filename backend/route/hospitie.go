package route

import (
	"encoding/json"
	"fmt"
	"hospitie/constant"
	"hospitie/dto"
	"hospitie/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitHospitieRoute() {
	hospitieGroup := constant.EchoInstance.Group("/hospitie")
	hs := service.NewHospitieService()

	hospitieGroup.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusFound, hs.GetAllHospities())
	})

	hospitieGroup.GET("/filter", func(c echo.Context) error {
		filter := map[string]interface{}{}
		json.NewDecoder(c.Request().Body).Decode(&filter)
		return c.JSON(http.StatusFound, hs.FilterHospitie(filter))
	})

	hospitieGroup.POST("", func(c echo.Context) error {
		hd := dto.CreateHospitieDTO(c)
		h, err := hs.CreateHospitie(hd)
		if err != nil {
			fmt.Println("err")
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, h)
	})

	hospitieGroup.PATCH("/:hospitie_id", func(c echo.Context) error {
		val := map[string]interface{}{}
		json.NewDecoder(c.Request().Body).Decode(&val)
		return c.JSON(http.StatusOK, hs.UpdateHospitie(c.Param("hospitie_id"), val))
	})

	hospitieGroup.DELETE("/:hospitie_id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, hs.DeleteHospitie(c.Param("hospitie_id")))
	})
}
