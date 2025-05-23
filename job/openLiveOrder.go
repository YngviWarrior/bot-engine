package job

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) OpenLiveOrder() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveOrder(make(<-chan struct{}))

	// transactionTypes := j.ExchangeMS.ListTransactionType(&pb.ListTransactionTypeRequest{})

	messages := j.RabbitMQ.Listen("order", "order")
	for order := range messages {
		fmt.Println(string(order.Body))
		var msg []bybitstructs.OrderDetail
		err := json.Unmarshal(order.Body, &msg)
		if err != nil {
			log.Println("OLO 01: ", err)
			continue
		}
		fmt.Printf("OLO: %+v\n ", msg)

		parts := strings.Split(msg[0].OrderLinkID, "-")
		linkId := parts[0]

		oph := j.ExchangeMS.GetOperationHistory(&pb.GetOperationHistoryRequest{
			OrderId: utils.ParseInt(linkId),
		})

		fmt.Printf("STATUS: %+v \n", msg[0].OrderStatus)
		switch msg[0].OrderStatus {
		case "New":
			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               utils.ParseFloat(msg[0].AvgPrice),
					CoinQuantity:            utils.ParseFloat(msg[0].CumExecQty),
					StablePrice:             utils.ParseFloat(msg[0].CumExecValue),
					StableQuantity:          utils.ParseFloat(msg[0].CumExecValue),
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 2,
					Fee:                     utils.ParseFloat(msg[0].CumExecFee),
				},
			})
		case "PartiallyFilled":
			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               utils.ParseFloat(msg[0].AvgPrice),
					CoinQuantity:            utils.ParseFloat(msg[0].CumExecQty),
					StablePrice:             utils.ParseFloat(msg[0].CumExecValue),
					StableQuantity:          utils.ParseFloat(msg[0].CumExecValue),
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 2,
					Fee:                     utils.ParseFloat(msg[0].CumExecFee),
				},
			})
		case "Filled":
			if msg[0].Side == "Sell" {
				operation := j.ExchangeMS.GetOperation(&pb.GetOperationRequest{
					OperationId: oph.GetOperationHistory().GetOperation(),
				})

				j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
					Operation: &pb.Operation{
						Operation:       operation.GetOperation().GetOperation(),
						User:            operation.GetOperation().GetUser(),
						Parity:          operation.GetOperation().GetParity(),
						Exchange:        operation.GetOperation().GetExchange(),
						Strategy:        operation.GetOperation().GetStrategy(),
						StrategyVariant: operation.GetOperation().GetStrategyVariant(),
						MtsStart:        operation.GetOperation().GetMtsStart(),
						MtsFinish:       operation.GetOperation().GetMtsFinish(),
						OpenPrice:       operation.GetOperation().GetOpenPrice(),
						ClosePrice:      operation.GetOperation().GetClosePrice(),
						InvestedAmount:  operation.GetOperation().GetInvestedAmount(),
						ProfitAmount:    operation.GetOperation().GetProfitAmount(),
						Profit:          operation.GetOperation().GetProfit(),
						Closed:          true,
						Audit:           operation.GetOperation().GetAudit(),
						Enabled:         false,
						TimesCanceled:   operation.GetOperation().GetTimesCanceled(),
						InTransaction:   false,
					},
				})
			}

			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               utils.ParseFloat(msg[0].AvgPrice),
					CoinQuantity:            utils.ParseFloat(msg[0].CumExecQty),
					StablePrice:             utils.ParseFloat(msg[0].CumExecValue),
					StableQuantity:          utils.ParseFloat(msg[0].CumExecValue),
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 3,
					Fee:                     utils.ParseFloat(msg[0].CumExecFee),
				},
			})
		case "Cancelled":
			operation := j.ExchangeMS.GetOperation(&pb.GetOperationRequest{
				OperationId: oph.GetOperationHistory().GetOperation(),
			})

			j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
				Operation: &pb.Operation{
					Operation:       operation.GetOperation().GetOperation(),
					User:            operation.GetOperation().GetUser(),
					Parity:          operation.GetOperation().GetParity(),
					Exchange:        operation.GetOperation().GetExchange(),
					Strategy:        operation.GetOperation().GetStrategy(),
					StrategyVariant: operation.GetOperation().GetStrategyVariant(),
					MtsStart:        operation.GetOperation().GetMtsStart(),
					MtsFinish:       operation.GetOperation().GetMtsFinish(),
					OpenPrice:       operation.GetOperation().GetOpenPrice(),
					ClosePrice:      operation.GetOperation().GetClosePrice(),
					InvestedAmount:  operation.GetOperation().GetInvestedAmount(),
					ProfitAmount:    operation.GetOperation().GetProfitAmount(),
					Profit:          operation.GetOperation().GetProfit(),
					Closed:          true,
					Audit:           operation.GetOperation().GetAudit(),
					Enabled:         false,
					TimesCanceled:   operation.GetOperation().GetTimesCanceled(),
					InTransaction:   false,
				},
			})

			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               utils.ParseFloat(msg[0].Price),
					CoinQuantity:            utils.ParseFloat(msg[0].Qty),
					StablePrice:             utils.ParseFloat(msg[0].CumExecValue),
					StableQuantity:          utils.ParseFloat(msg[0].CumExecValue),
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 4,
					Fee:                     utils.ParseFloat(msg[0].CumExecFee),
				},
			})
		}

	}
}
