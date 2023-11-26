package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Principal string `gorm:"unique"`
	Password  string
}
