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

	messages := j.RabbitMQ.Listen("livetrade", "order.create")
	for order := range messages {
		fmt.Println(order)
		var msg bybitstructs.OrderResponse
		err := json.Unmarshal(order.Body, &msg)
		if err != nil {
			log.Println("Failed to parse msg", err)
			continue
		}

		// opId := utils.ParseInt(strings.Fields(msg.ReqID)[1])

		// // Find order in Bybit API
		// orderHistory := bybit.OrderHistory(&bybitstructs.OrderHistoryParams{
		// 	OrderId:  msg.Data.OrderID,
		// 	Category: "spot",
		// })

		// var closed bool
		// switch orderHistory.Result.List[0].Side {
		// case "Buy":
		// 	closed = false
		// case "Sell":
		// 	closed = true
		// }

		// operation := j.ExchangeMS.GetOperation(&pb.GetOperationRequest{
		// 	OperationId: opId,
		// })

		// if operation.Operation == nil {
		// 	log.Panicln("OLT 01: No Operation Found")
		// }

		// j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
		// 	Operation: &pb.Operation{
		// 		Operation:       opId,
		// 		User:            operation.GetOperation().GetUser(),
		// 		Parity:          operation.GetOperation().GetParity(),
		// 		Exchange:        operation.GetOperation().GetExchange(),
		// 		Strategy:        operation.GetOperation().GetStrategy(),
		// 		StrategyVariant: operation.GetOperation().GetStrategyVariant(),
		// 		MtsStart:        operation.GetOperation().GetMtsStart(),
		// 		MtsFinish:       operation.GetOperation().GetMtsFinish(),
		// 		OpenPrice:       operation.GetOperation().GetOpenPrice(),
		// 		ClosePrice:      operation.GetOperation().GetClosePrice(),
		// 		InvestedAmount:  operation.GetOperation().GetInvestedAmount(),
		// 		ProfitAmount:    operation.GetOperation().GetProfitAmount(),
		// 		Profit:          operation.GetOperation().GetProfit(),
		// 		Closed:          closed,
		// 		Audit:           operation.GetOperation().GetAudit(),
		// 		Enabled:         true,
		// 		InTransaction:   true,
		// 	},
		// })

	}
}
