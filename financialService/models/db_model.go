package models

import (
	"database/sql"
	"financialService/types"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type DBAccount struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId        uuid.UUID
	Name          string
	AccountNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type DBBalance struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	AccountId uuid.UUID
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DBBudget struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Date        datatypes.Date
	Budget      float64
	Realization float64   `gorm:"default:0"`
	AccountId   uuid.UUID `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type DBTransaction struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	AccountId   uuid.UUID
	DateTime    datatypes.Date
	Description sql.NullString
	Amount      float64
	Type        types.TransactionType
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
