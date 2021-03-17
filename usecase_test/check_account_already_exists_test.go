package usecase

import (
	"testing"

	"github.com/andersonmarin/pismo-challenge-api/model"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"
	"github.com/andersonmarin/pismo-challenge-api/usecase"
	"gorm.io/gorm"
)

func TestCheckAccountAlreadyExists(t *testing.T) {
	type args struct {
		db             *gorm.DB
		documentNumber string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "account does not exist",
			args: args{
				db:             infrastructure.MemoryDatabase(),
				documentNumber: "12345678900",
			},
			wantErr: false,
		},
		{
			name: "account already exist",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.Account{DocumentNumber: "12345678900"})
				}),
				documentNumber: "12345678900",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := usecase.CheckAccountAlreadyExists(tt.args.db, tt.args.documentNumber); (err != nil) != tt.wantErr {
				t.Errorf("CheckAccountAlreadyExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
