package web

import (
	"log"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.GET("/:address", GetRiddle(conn))
	e.POST("/hash", HashAnswer(conn))
	e.POST("/validate", ValidateAnswer(conn))

	return e.Start(addr)
}
