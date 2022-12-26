package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      int    `json:"id" gorm:"primaryKey"`
	Hash    string `json:"hash" gorm:"hash"`
	Email   string `json:"email" gorm:"email"`
	Name    string `json:"name" gorm:"name"`
	Sername string `json:"sername" gorm:"sername"`
}
