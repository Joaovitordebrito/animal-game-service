package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserName   string
	CurrentBet string
	BetDate    time.Time
}
