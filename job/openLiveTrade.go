package job

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) OpenLiveTrade(orderChannel <-chan *bybitstructs.OrderRequest) {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveTrade(orderChannel, make(<-chan struct{}))

	messages := j.RabbitMQ.Listen("", "execution")
	for order := range messages {
		fmt.Println(order)
		var msg bybitstructs.OrderResponse
		err := json.Unmarshal(order.Body, &msg)
		if err != nil {
			log.Println("Failed to parse msg", err)
			continue
		}
		fmt.Println("Parsed msg", msg)

		// j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
		// 	Operation: &pb.Operation{
		// 		Operation:       op.GetOperation(),
		// 		User:            op.GetUser(),
		// 		Parity:          op.GetParity(),
		// 		Exchange:        op.GetExchange(),
		// 		Strategy:        op.GetStrategy(),
		// 		StrategyVariant: op.GetStrategyVariant(),
		// 		MtsStart:        op.GetMtsStart(),
		// 		MtsFinish:       op.GetMtsFinish(),
		// 		OpenPrice:       op.GetOpenPrice(),
		// 		ClosePrice:      op.GetClosePrice(),
		// 		InvestedAmount:  op.GetInvestedAmount(),
		// 		ProfitAmount:    op.GetProfitAmount(),
		// 		Profit:          op.GetProfit(),
		// 		Closed:          op.GetClosed(),
		// 		Audit:           op.GetAudit(),
		// 		Enabled:         false,
		// 		TimesCanceled:   op.GetTimesCanceled(),
		// 	},
		// })
	}
}
