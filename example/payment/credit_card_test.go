package payment

import "testing"


func TestDirectPaymentCreditCard(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test direct payment credit card",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DirectPaymentCreditCard(); (err != nil) != tt.wantErr {
				t.Errorf("DirectPaymentCreditCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}