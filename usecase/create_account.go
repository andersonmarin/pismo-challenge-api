package usecase

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB, documentNumber string) (*model.Account, error) {
	account := model.Account{DocumentNumber: documentNumber}
	if err := db.Create(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
