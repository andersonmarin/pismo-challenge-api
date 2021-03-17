package usecase

import "testing"

func TestCheckAccountDocumentNumber(t *testing.T) {
	type args struct {
		documentNumber string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "document number ok",
			args:    args{"12345678900"},
			wantErr: false,
		},
		{
			name:    "document number empty error",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "document number length error",
			args:    args{"123456789"},
			wantErr: true,
		},
		{
			name:    "document number format error",
			args:    args{"123456789AB"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAccountDocumentNumber(tt.args.documentNumber); (err != nil) != tt.wantErr {
				t.Errorf("CheckAccountDocumentNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
