package models

import (
	"gorm.io/gorm"
)

type Riddle struct {
	gorm.Model
	Riddle   string `json:"riddle"`
	Solution string `json:"-"`
}

type AssignedRiddle struct {
	gorm.Model
	Riddle     Riddle `json:"riddle"`
	AssignedTo string `json:"assigned_to"`
}

type SolvedRiddle struct {
	gorm.Model
	Riddle   Riddle `json:"riddle"`
	SolvedBy string `json:"solved_by"`
}
