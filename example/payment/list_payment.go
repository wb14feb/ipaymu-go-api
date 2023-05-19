package payment

import (
	"fmt"
	ipaymu "gitlab.ipaymu.com/plugin/ipaymu-go-api"
)

func ListPaymentMethod() error {
	// initiate client
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	// api call
	balance, err := client.ListPaymentMethod()
	if err != nil {
		return err
	}

	fmt.Printf("%v", balance)

	return nil
}
