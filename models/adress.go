package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  uuid.UUID `gorm:"user_id"`
	ID      uuid.UUID `gorm:"primaryKey"`
	Address string    `gorm:"address"`
}
