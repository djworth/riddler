package web

import (
	"math/rand"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/djworth/riddler/pkg/db/models"
)

func AssignRiddle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")

		exists := models.AssignedRiddle{}
		results := db.Where("assigned_to = ?", address).Preload("Riddle").Take(&exists)
		if results.RowsAffected == 1 {
			return c.JSON(http.StatusOK, exists.ToMap())
		}

		riddles := []models.Riddle{}

		db.Find(&riddles)

		i := rand.Intn(len(riddles))

		ar := models.AssignedRiddle{
			Riddle:     riddles[i],
			AssignedTo: address,
		}
		db.Create(&ar)

		return c.JSON(http.StatusOK, ar.ToMap())
	}
}

func GetRiddle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")

		ar := models.AssignedRiddle{}

		results := db.Where("assigned_to = ?", address).Preload("Riddle").Take(&ar)

		if results.RowsAffected == 0 {
			return c.JSON(http.StatusOK, map[string]string{})
		}
		return c.JSON(http.StatusOK, ar.ToMap())
	}
}
