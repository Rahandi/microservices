package internals

import (
	"FinancialService/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	config *Config
}

func NewDatabase(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (d *Database) Connect() *gorm.DB {
	dsn := d.config.DatabaseUsername + ":" + d.config.DatabasePassword + "@tcp(" + d.config.DatabaseHost + ":" + d.config.DatabasePort + ")/" + d.config.DatabaseName + "?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic(err)
	}

	database.AutoMigrate(
		&models.DBAccount{},
		&models.DBBalance{},
		&models.DBBudget{},
		&models.DBTransaction{},
	)

	return database
}
