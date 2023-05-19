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

	notifyUrl := "http://localhost/notify-url"
	expired := int8(24)
	expiredType := ipaymu.Hours
	referenceId := "trx-123"

	request := ipaymu.NewRequestDirectVA(ipaymu.CimbNiaga)
	request.AddBuyer("buyer", "08123123123", "email@test.com")
	request.NotifyUrl = &notifyUrl
	request.Expired = &expired
	request.ExpiredType = &expiredType
	request.ReferenceId = &referenceId
	request.Amount = 100000

	va, err := client.DirectPaymentVA(*request)
	if err != nil {
		return err
	}

	fmt.Println(va)

	return nil
}
