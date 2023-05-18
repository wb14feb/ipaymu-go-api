package payment

import "testing"

func TestDirectVirtualAccount(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test direct payment virtual account",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DirectVirtualAccount(); (err != nil) != tt.wantErr {
				t.Errorf("DirectVirtualAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
