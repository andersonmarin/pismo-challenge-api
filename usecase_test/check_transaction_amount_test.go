package usecase

import (
	"testing"

	"github.com/andersonmarin/pismo-challenge-api/model"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"
	"github.com/andersonmarin/pismo-challenge-api/usecase"
	"gorm.io/gorm"
)

func TestCheckTransactionAmount(t *testing.T) {
	type args struct {
		db              *gorm.DB
		operationTypeID uint64
		amount          float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "amount negative ok",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Negative",
						Negative: true,
					})
				}),
				operationTypeID: 1,
				amount:          -123.45,
			},
			wantErr: false,
		},
		{
			name: "amount negative error",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Negative",
						Negative: true,
					})
				}),
				operationTypeID: 1,
				amount:          123.45,
			},
			wantErr: true,
		},
		{
			name: "amount positive ok",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Positive",
						Negative: false,
					})
				}),
				operationTypeID: 1,
				amount:          123.45,
			},
			wantErr: false,
		},
		{
			name: "amount positive error",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Negative",
						Negative: false,
					})
				}),
				operationTypeID: 1,
				amount:          -123.45,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := usecase.CheckTransactionAmount(tt.args.db, tt.args.operationTypeID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CheckTransactionAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
