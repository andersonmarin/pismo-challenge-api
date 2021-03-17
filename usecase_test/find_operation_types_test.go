package usecase

import (
	"reflect"
	"testing"

	"github.com/andersonmarin/pismo-challenge-api/infrastructure"

	"github.com/andersonmarin/pismo-challenge-api/model"
	"github.com/andersonmarin/pismo-challenge-api/usecase"
	"gorm.io/gorm"
)

func TestFindOperationTypes(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.OperationType
		wantErr bool
	}{
		{
			name: "find ok",
			args: args{
				db: infrastructure.MemoryDatabase(func(db *gorm.DB) {
					db.Create(&model.OperationType{
						Name:     "Positive",
						Negative: false,
					}).Create(&model.OperationType{
						Name:     "Negative",
						Negative: true,
					})
				}),
			},
			want: []*model.OperationType{
				{
					ID:       1,
					Name:     "Positive",
					Negative: false,
				},
				{
					ID:       2,
					Name:     "Negative",
					Negative: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.FindOperationTypes(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOperationTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOperationTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
