package usecase

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func FindOperationTypes(db *gorm.DB) ([]*model.OperationType, error) {
	var operationTypes []*model.OperationType
	if err := db.Find(&operationTypes).Error; err != nil {
		return nil, err
	}
	return operationTypes, nil
}
