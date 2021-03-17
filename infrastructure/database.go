package infrastructure

import (
	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=pismo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Account{}, &model.OperationType{}, &model.Transaction{})
	if err != nil {
		panic(err)
	}

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

	if err = db.First(&model.OperationType{}).Error; err == gorm.ErrRecordNotFound {
		err = db.CreateInBatches(&defaultOperationTypes, len(defaultOperationTypes)).Error
		if err != nil {
			panic(err)
		}
	}

	return db
}
