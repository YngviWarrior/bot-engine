package job

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) OpenLiveOrder() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveOrder(make(<-chan struct{}))

	messages := j.RabbitMQ.Listen("order", "order")
	for order := range messages {
		var msg []bybitstructs.OrderDetail
		err := json.Unmarshal(order.Body, &msg)
		if err != nil {
			log.Println("OLO 01: ", err)
			continue
		}
		// fmt.Printf("OLO: %+v\n ", msg)

		parts := strings.Split(msg[0].OrderLinkID, "-")
		linkId := parts[0]

		oph := j.ExchangeMS.GetOperationHistory(&pb.GetOperationHistoryRequest{
			OrderId: utils.ParseInt(linkId),
		})

		switch msg[0].OrderStatus {
		case "New":
			fmt.Println("Aqui")
			fmt.Println(msg[0])
			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               msg[0].AvgPrice,
					CoinQuantity:            msg[0].CumExecQty,
					StablePrice:             msg[0].CumExecValue,
					StableQuantity:          msg[0].CumExecValue,
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 1,
					Fee:                     msg[0].CumExecFee,
				},
			})
		case "PartiallyFilled":
			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               msg[0].AvgPrice,
					CoinQuantity:            msg[0].CumExecQty,
					StablePrice:             msg[0].CumExecValue,
					StableQuantity:          msg[0].CumExecValue,
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 2,
					Fee:                     msg[0].CumExecFee,
				},
			})
		case "Filled":
			var fee float64
			if msg[0].Side == "Sell" {
				operation := j.ExchangeMS.GetOperation(&pb.GetOperationRequest{
					OperationId: oph.GetOperationHistory().GetOperation(),
				})

				cumExecFee, err := strconv.ParseFloat(msg[0].CumExecFee, 64)
				if err != nil {
					log.Panic("OLO 02: ", err)
				}
				avgPrice, err := strconv.ParseFloat(msg[0].AvgPrice, 64)
				if err != nil {
					log.Panic("OLO 02: ", err)
				}

				fee = cumExecFee / avgPrice
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
			} else {
				cumExecFee, err := strconv.ParseFloat(msg[0].CumExecFee, 64)
				if err != nil {
					log.Panic("OLO 02: ", err)
				}
				fee = cumExecFee
			}

			cumExecQty, err := strconv.ParseFloat(msg[0].CumExecQty, 64)
			if err != nil {
				log.Panic("OLO 03 : ", err)
			}

			j.ExchangeMS.UpdateOperationHistory(&pb.UpdateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					OperationHistory:        oph.GetOperationHistory().GetOperationHistory(),
					Operation:               oph.GetOperationHistory().GetOperation(),
					TransactionType:         oph.GetOperationHistory().GetTransactionType(),
					CoinPrice:               msg[0].AvgPrice,
					CoinQuantity:            fmt.Sprintf("%0.11f", cumExecQty-fee),
					StablePrice:             msg[0].CumExecValue,
					StableQuantity:          msg[0].CumExecValue,
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 3,
					Fee:                     fmt.Sprintf("%0.11f", fee),
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
					CoinPrice:               msg[0].Price,
					CoinQuantity:            msg[0].Qty,
					StablePrice:             msg[0].CumExecValue,
					StableQuantity:          msg[0].CumExecValue,
					OperationExchangeId:     msg[0].OrderID,
					OperationExchangeStatus: 4,
					Fee:                     msg[0].CumExecFee,
				},
			})
		}

	}
}
