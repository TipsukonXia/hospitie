package constant

import "github.com/labstack/echo/v4"

var EchoInstance *echo.Echo

func InitEchoInstance() {
	if EchoInstance == nil {
		print("init echo instance")
		EchoInstance = echo.New()
	}
	print("echo instance already init")
}
