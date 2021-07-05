package app

import "github.com/labstack/echo"

var (
	e *echo.Echo
)

func StartApplocation(address string) {
	e = echo.New()
	router()
	e.Logger.Print("starting server...")
	e.Logger.Fatal(e.Start(address))
}
