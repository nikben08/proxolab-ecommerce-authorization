package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	gorm.Model
	UserID      uuid.UUID `gorm:"user_id"`
	ID          uuid.UUID `gorm:"primaryKey"`
	PhoneNumber string    `gorm:"phone_number"`
}
