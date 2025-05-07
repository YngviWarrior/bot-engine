package job

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) OpenLiveOrder() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveExec(make(<-chan struct{}))

	// transactionTypes := j.ExchangeMS.ListTransactionType(&pb.ListTransactionTypeRequest{})

	messages := j.RabbitMQ.Listen("", "order")
	for order := range messages {
		fmt.Println(order)
		var msg bybitstructs.OrderMessage
		err := json.Unmarshal(order.Body, &msg)
		if err != nil {
			log.Println("OLO 01: ", err)
			continue
		}
		fmt.Println("OLO: ", msg)

		oph := j.ExchangeMS.GetOperationHistory(&pb.GetOperationHistoryRequest{
			OrderId: utils.ParseInt(msg.Data[0].OrderID),
		})

		op := j.ExchangeMS.GetOperation(&pb.GetOperationRequest{
			OperationId: oph.GetOperationHistory().GetOperation(),
		})

		switch msg.Data[0].OrderStatus {
		case "NEW":

		case "PARTIALLY_FILLED":

		case "FILLED":
			var closed bool
			var enabled bool
			if oph.OperationHistory.GetTransactionType() == 2 {
				closed = true
				enabled = false
			} else {
				closed = false
				enabled = true
			}

			j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
				Operation: &pb.Operation{
					Operation:       oph.GetOperationHistory().GetOperation(),
					User:            op.GetOperation().GetUser(),
					Parity:          op.GetOperation().GetParity(),
					Exchange:        op.GetOperation().GetExchange(),
					Strategy:        op.GetOperation().GetStrategy(),
					StrategyVariant: op.GetOperation().GetStrategyVariant(),
					MtsStart:        op.GetOperation().GetMtsStart(),
					MtsFinish:       op.GetOperation().GetMtsFinish(),
					OpenPrice:       op.GetOperation().GetOpenPrice(),
					ClosePrice:      op.GetOperation().GetClosePrice(),
					InvestedAmount:  op.GetOperation().GetInvestedAmount(),
					ProfitAmount:    op.GetOperation().GetProfitAmount(),
					Profit:          op.GetOperation().GetProfit(),
					Closed:          closed,
					Audit:           false,
					Enabled:         enabled,
					TimesCanceled:   op.GetOperation().GetTimesCanceled(),
				},
			})

			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               oph.GetOperationHistory().GetCoinPrice(),
					CoinQuantity:            oph.GetOperationHistory().GetCoinQuantity(),
					StablePrice:             oph.GetOperationHistory().GetStablePrice(),
					StableQuantity:          oph.GetOperationHistory().GetStableQuantity(),
					OperationExchangeId:     oph.GetOperationHistory().GetOperationExchangeId(),
					OperationExchangeStatus: 3,
					Fee:                     oph.GetOperationHistory().GetFee(),
				},
			})

		case "CANCELED":
			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               oph.GetOperationHistory().GetCoinPrice(),
					CoinQuantity:            oph.GetOperationHistory().GetCoinQuantity(),
					StablePrice:             oph.GetOperationHistory().GetStablePrice(),
					StableQuantity:          oph.GetOperationHistory().GetStableQuantity(),
					OperationExchangeId:     oph.GetOperationHistory().GetOperationExchangeId(),
					OperationExchangeStatus: 4,
					Fee:                     oph.GetOperationHistory().GetFee(),
				},
			})

		case "PENDING_CANCEL":

		case "PENDING_NEW":

		case "REJECTED":

		case "ORDER_NEW":

		case "ORDER_CANCELED":

		case "ORDER_FILLED":

		case "ORDER_FAILED":

		}

	}
}
