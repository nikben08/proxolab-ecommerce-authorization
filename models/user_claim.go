package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserClaim struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"primaryKey"`
	AccessLevel int       `gorm:"access_level;default:0"`
}
