package entities

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey; coolumn:id"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
