package payment

import "testing"

func TestListPaymentMethod(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test list payment method",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListPaymentMethod(); (err != nil) != tt.wantErr {
				t.Errorf("ListPaymentMethod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
