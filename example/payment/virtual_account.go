package payment

import (
	"fmt"
	ipaymu "gitlab.ipaymu.com/plugin/ipaymu-go-api"
)

func DirectVirtualAccount() error {
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	request := ipaymu.NewRequestDirectVA(ipaymu.CimbNiaga)
	request.AddBuyer("buyer", "phone", "email@test.com")
	request.NotifyUrl = "http://localhost/notify-url"
	request.Expired = 24
	request.ExpiredType = ipaymu.Hours
	request.ReferenceId = "trx-123"
	request.Amount = 100000

	va, err := client.DirectPaymentVA(*request)
	if err != nil {
		return err
	}

	fmt.Println(va)

	return nil
}
