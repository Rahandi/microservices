package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
