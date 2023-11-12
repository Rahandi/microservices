package models

import (
	"FinancialService/types"
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	UserId        uuid.UUID
	Name          string
	AccountNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Balance struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	AccountId uuid.UUID
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Budget struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Date        datatypes.Date
	Budget      float64
	Realization float64   `gorm:"default:0"`
	AccountId   uuid.UUID `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Transaction struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	AccountId   uuid.UUID
	DateTime    datatypes.Date
	Description sql.NullString
	Amount      float64
	Type        types.TransactionType
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
