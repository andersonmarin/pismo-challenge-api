package usecase

import (
	"errors"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

var ErrTransactionAmountNotAccepted = errors.New("amount not accepted")

func CheckTransactionAmount(db *gorm.DB, operationTypeID uint64, amount float64) error {
	var operationType model.OperationType
	if err := db.First(&operationType, operationTypeID).Error; err != nil {
		return err
	}

	if (operationType.Negative && amount > 0) || !operationType.Negative && amount < 0 {
		return ErrTransactionAmountNotAccepted
	}

	return nil
}
