package entities

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
