package payment

import "testing"

func TestHistoryTransaction(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test history transaction",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HistoryTransaction(); (err != nil) != tt.wantErr {
				t.Errorf("HistoryTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
