package payment

import (
	"fmt"
	"strings"

	ipaymu "github.com/wb14feb/ipaymu-api"
)

func HistoryTransaction() error {
	// initiate client
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	// prepare request
	status := ipaymu.Success
	orderBy := ipaymu.Paid
	listID := []string{"68369", "44396", "44389"}
	listBulk := strings.Join(listID, ",")
	request := ipaymu.NewRequestTransactionHistory()
	request.Status = &status
	request.OrderBy = &orderBy
	request.BulkId = &listBulk

	// api call
	api, err := client.HistoryTransaction(*request)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", api)
	return nil
}
