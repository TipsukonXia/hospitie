package core

import (
	"fmt"
)

func ErrorHandler(err error) {
	errMsg := fmt.Sprintf("error with %s", err)
	print(errMsg)
}
