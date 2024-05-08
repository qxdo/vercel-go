package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qxdo/vercel-go/api"
)

func main() {
	e := echo.New()
	e.GET("/aprs_passcode", api.Test)
	e.Start("127.0.0.1")
}
