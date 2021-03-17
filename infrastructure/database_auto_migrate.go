package infrastructure

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func DatabaseAutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Account{},
		&model.OperationType{},
		&model.Transaction{},
	)
}
