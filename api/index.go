package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var srv http.Handler

func init() {
	e := echo.New()
	e.GET("/books", Test)
	srv = e
}

func Test(e echo.Context) error {
	return e.String(200, "abcdefgh")
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
