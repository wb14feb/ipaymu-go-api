package payment

import (
	"fmt"

	ipaymu "github.com/wb14feb/ipaymu-api"
)

func Balance() error {
	// initiate client
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	// api call
	balance, err := client.GetBalance()
	if err != nil {
		return err
	}

	fmt.Printf("%v", balance)

	return nil
}
