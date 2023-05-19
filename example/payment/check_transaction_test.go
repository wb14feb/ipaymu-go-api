package payment

import "testing"

func TestCheckTransaction(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test check transaction",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckTransaction(); (err != nil) != tt.wantErr {
				t.Errorf("CheckTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
