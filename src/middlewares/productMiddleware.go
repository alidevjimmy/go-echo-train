package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
)

func OnlyUsers(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("you is user, so you can watch product")
	return func(c echo.Context) error {
		return next(c)
	}
}
