package payment

import (
	"fmt"

	ipaymu "github.com/wb14feb/ipaymu-go-api"
)

func DirectPaymentCreditCard() error {
	// initiate client
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	// prepare the request
	var exp int8 = 24
	var expType ipaymu.ExpiredType = ipaymu.Hours
	var refId string = "20210901123456" // change based on needs
	var notifUrl string = "http://localhost/notify-url"
	request := ipaymu.NewRequestDirectCreditCard()

	request.AddBuyer("buyer", "08123456789", "email@test.com")
	request.NotifyUrl = &notifUrl
	request.Expired = &exp
	request.ExpiredType = &expType
	request.ReferenceId = &refId
	request.Amount = 100000

	// api call
	resp, err := client.DirectPaymentCreditCard(*request)

	if err != nil {
		return err
	}

	fmt.Println(*resp.Data)

	return nil
}
