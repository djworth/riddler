package models

import (
	"time"

	"gorm.io/gorm"
)

type Riddle struct {
	ID        int            `json:"id" gorm:"type:int;primaryKey;autoIncrement:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Riddle    string         `json:"riddle"`
	Answer    string         `json:"-"`
}

type AssignedRiddle struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
	Riddle     Riddle    `json:"riddle" gorm:"foreignKey:RiddleID;references:ID"`
	RiddleID   int       `json:"-" gorm:"type:int;primaryKey"`
	AssignedTo string    `json:"assigned_to" gorm:"type:text;primaryKey"`
}

type ValidateAnswers struct {
	CreatedAt  time.Time `json:"created_at"`
	RiddleID   int       `json:"-" gorm:"foreignKey:RiddleID;references:ID"`
	AssignedTo string    `json:"assigned_to" gorm:"type:text;primaryKey"`
	Answer     string    `json:"answer" gorm:"type:text"`
}

func (r *Riddle) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":     r.ID,
		"riddle": r.Riddle,
	}
}

func (ar *AssignedRiddle) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          ar.RiddleID,
		"riddle":      ar.Riddle.Riddle,
		"assigned_to": ar.AssignedTo,
	}
}
