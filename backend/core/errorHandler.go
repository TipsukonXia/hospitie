package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
