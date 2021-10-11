package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Serve(addr string) error {

	e := echo.New()

	e.GET("/:address", func(c echo.Context) error {
		return c.String(http.StatusOK, "Riddle Me This...")
	})

	return e.Start(addr)
}
