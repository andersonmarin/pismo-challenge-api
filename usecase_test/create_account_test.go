package usecase

import (
	"reflect"
	"testing"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"github.com/andersonmarin/pismo-challenge-api/usecase"
	"gorm.io/gorm"
)

func TestCreateAccount(t *testing.T) {
	type args struct {
		db             *gorm.DB
		documentNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "create account ok",
			args: args{
				db:             infrastructure.MemoryDatabase(),
				documentNumber: "12345678900",
			},
			want: &model.Account{
				ID:             1,
				DocumentNumber: "12345678900",
			},
			wantErr: false,
		},
		{
			name: "create account error",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.Account{
						DocumentNumber: "12345678900",
					})
				}),
				documentNumber: "12345678900",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.CreateAccount(tt.args.db, tt.args.documentNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
