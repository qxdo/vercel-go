package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/aprs_passcode")
	e.Start(":8001")

}
