package ipaymu_go_api

import (
	"reflect"
	"testing"
)

func TestClient_DirectPaymentVA(t *testing.T) {
	cl := NewClient()
	cl.EnvApi = Sandbox
	cl.VirtualAccount = "1179000899"
	cl.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"

	notifyUrl := "https://your-website.com/callback"

	req := NewRequestDirectVA(BNI)
	req.AddBuyer("prima", "082313131", "test@email.com")
	req.NotifyUrl = &notifyUrl
	req.Amount = 100000

	type args struct {
		body RequestDirectVA
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "test direct va",
			args:    args{body: *req},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cl
			got, err := c.DirectPaymentVA(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectPaymentVA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Status != tt.want {
				t.Errorf("DirectPaymentVA() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RedirectPayment(t *testing.T) {
	cl := NewClient()
	cl.EnvApi = Sandbox
	cl.VirtualAccount = "1179002284460840"
	cl.ApiKey = "CDC1AD1E-A19C-40E6-998D-9736BF4E42FA"

	notifyUrl := "https://your-website.com/callback"
	returnUrl := "https://your-website.com/thank-you-page"
	cancelUrl := "https://your-website.com/failed-page"

	req := NewRequestRedirect()
	req.ReturnUrl = &returnUrl
	req.NotifyUrl = &notifyUrl
	req.CancelUrl = &cancelUrl
	req.AddProduct("product", 1, 100000, nil, nil)

	type args struct {
		request RequestRedirect
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		{
			name:    "test redirect payment",
			args:    args{*req},
			want:    Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cl
			got, err := c.RedirectPayment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedirectPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedirectPayment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
