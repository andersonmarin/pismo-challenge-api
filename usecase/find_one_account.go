package usecase

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func FindOneAccount(db *gorm.DB, id uint64) (*model.Account, error) {
	var account model.Account
	if err := db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
