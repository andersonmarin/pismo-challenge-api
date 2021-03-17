package usecase

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func TestCreateTransaction(t *testing.T) {
	hash := func(t *model.Transaction) string {
		if t == nil {
			return ""
		}
		return fmt.Sprintf(
			"AccountID: %d, OperationTypeID: %d, Amount: %f",
			t.AccountID,
			t.OperationTypeID,
			t.Amount,
		)
	}

	type args struct {
		db              *gorm.DB
		accountID       uint64
		operationTypeID uint64
		amount          float64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Transaction
		wantErr bool
	}{
		{
			name: "create ok",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.Account{
						DocumentNumber: "12345678900",
					})
					db.Create(&model.OperationType{
						Name:     "Positive",
						Negative: false,
					})
				}),
				accountID:       1,
				operationTypeID: 1,
				amount:          123.45,
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       time.Now(),
			},
			wantErr: false,
		},
		{
			name: "create error account",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Positive",
						Negative: false,
					})
				}),
				accountID:       1,
				operationTypeID: 1,
				amount:          123.45,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create error operation type",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.Account{
						DocumentNumber: "12345678900",
					})
				}),
				accountID:       1,
				operationTypeID: 1,
				amount:          123.45,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.CreateTransaction(tt.args.db, tt.args.accountID, tt.args.operationTypeID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(hash(got), hash(tt.want)) {
				t.Errorf("CreateTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
