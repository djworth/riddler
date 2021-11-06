package web

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/djworth/riddler/pkg/db/models"
	"github.com/djworth/riddler/pkg/hash"
)

type ValidateAnswerRequest struct {
	RiddleID      int    `json:"id"`
	WalletAddress string `json:"address"`
	Answer        string `json:"answer"`
}

type SuccessFailureResponse struct {
	Response string `json:"response"`
}

var success = SuccessFailureResponse{Response: "success"}
var failure = SuccessFailureResponse{Response: "failure"}

func logValidateRequest(db *gorm.DB, req *ValidateAnswerRequest) error {
	record := models.ValidateAnswers{
		RiddleID:   req.RiddleID,
		AssignedTo: req.WalletAddress,
		Answer:     req.Answer,
	}
	result := db.Create(&record)
	return result.Error

}

func GetRiddle(db *gorm.DB) echo.HandlerFunc {
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

func ValidateAnswer(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := ValidateAnswerRequest{}

		err := c.Bind(&req)
		if err != nil {
			return err
		}

		err = logValidateRequest(db, &req)
		if err != nil {
			log.Println(err)
		}

		assignedRiddle := models.AssignedRiddle{}
		results := db.Where("assigned_to = ?", req.Answer, "riddle_id = ?", req.RiddleID).Preload("Riddle").Take(&assignedRiddle)
		if results.RowsAffected != 1 {
			return c.JSON(http.StatusBadRequest, errors.New("unable to find riddle for addressed"))
		}

		hashedAnswer := hash.Hash(req.Answer)
		if assignedRiddle.Riddle.Answer != hashedAnswer {
			return c.JSON(http.StatusBadRequest, failure)
		}

		return c.JSON(http.StatusOK, success)
	}
}

func HashAnswer(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := map[string]string{}

		err := c.Bind(&data)
		if err != nil {
			return err
		}

		results := map[string]string{}
		results["hashed"] = hash.Hash(strings.ToLower(data["answer"]))

		return c.JSON(http.StatusOK, results)
	}
}
