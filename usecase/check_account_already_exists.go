package usecase

import (
	"errors"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

var ErrAccountDocumentNumberAlreadyExists = errors.New("document number already exists")

func CheckAccountAlreadyExists(db *gorm.DB, documentNumber string) error {
	err := db.First(&model.Account{}, "document_number = ?", documentNumber).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		return ErrAccountDocumentNumberAlreadyExists
	default:
		return err
	}
}
