package entities

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primaryKey; coolumn:id" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
