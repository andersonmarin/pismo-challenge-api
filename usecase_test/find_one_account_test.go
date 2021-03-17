package usecase

import (
	"reflect"
	"testing"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"gorm.io/gorm"
)

func TestFindOneAccount(t *testing.T) {
	type args struct {
		db *gorm.DB
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "find ok",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.Account{
						DocumentNumber: "12345678900",
					})
				}),
				id: 1,
			},
			want: &model.Account{
				ID:             1,
				DocumentNumber: "12345678900",
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				db: infrastructure.MemoryDatabase(),
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.FindOneAccount(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOneAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOneAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
