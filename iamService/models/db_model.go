package models

import "github.com/google/uuid"

type DBUser struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
