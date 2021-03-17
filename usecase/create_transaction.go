package usecase

import (
	"time"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func CreateTransaction(db *gorm.DB, accountID, operationTypeID uint64, amount float64) (*model.Transaction, error) {
	transaction := model.Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}
	if err := db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
