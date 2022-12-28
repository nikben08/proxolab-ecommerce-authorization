package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID     `gorm:"primaryKey"`
	Hash        string        `gorm:"hash"`
	Email       string        `gorm:"email"`
	Name        string        `gorm:"name"`
	Sername     string        `gorm:"sername"`
	PhoneNumber []PhoneNumber `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;" json:"-"`
	Address     []Address     `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;" json:"-"`
}
