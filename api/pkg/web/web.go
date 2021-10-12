package web

import (
	"log"

	echo "github.com/labstack/echo/v4"

	"github.com/djworth/riddler/pkg/db"
)

func Serve(addr string) error {

	conn, err := db.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(conn); err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	e.GET("/:address", GetRiddle(conn))
	e.PUT("/:address", AssignRiddle(conn))

	return e.Start(addr)
}
