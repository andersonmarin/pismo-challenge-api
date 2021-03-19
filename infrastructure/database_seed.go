package infrastructure

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func DatabaseSeed(db *gorm.DB) error {
	defaultOperationTypes := []*model.OperationType{
		{
			Name:     "COMPRA A VISTA",
			Negative: true,
		},
		{
			Name:     "COMPRA PARCELADA",
			Negative: true,
		},
		{
			Name:     "SAQUE",
			Negative: true,
		},
		{
			Name:     "PAGAMENTO",
			Negative: false,
		},
	}

	if err := db.First(&model.OperationType{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.CreateInBatches(&defaultOperationTypes, len(defaultOperationTypes)).Error
		}
		return err
	}

	return nil
}
