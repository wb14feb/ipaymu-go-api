package payment

import (
	"fmt"
	ipaymu "gitlab.ipaymu.com/plugin/ipaymu-go-api"
	"time"
)

func DirectVirtualAccount() error {
	// initiate client
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	// prepare the request
	var exp int8 = 24
	var expType ipaymu.ExpiredType = ipaymu.Hours
	var refId string = time.Now().Format("20060102150405") // change based on needs
	var notifUrl string = "http://localhost/notify-url"
	request := ipaymu.NewRequestDirectVA(ipaymu.CimbNiaga)
	request.AddBuyer("buyer", "08123456789", "email@test.com")
	request.NotifyUrl = &notifUrl
	request.Expired = &exp
	request.ExpiredType = &expType
	request.ReferenceId = &refId
	request.Amount = 100000

	// api call
	va, err := client.DirectPaymentVA(*request)
	if err != nil {
		return err
	}

	fmt.Println(*va.Data)

	return nil
}
