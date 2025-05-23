package botengine

import (
	"fmt"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

type BuyCoinParams struct {
	Operation  uint64
	ClosePrice float64
	Exchange   uint64
	OpAmount   string
	Symbol     string
	OrderType  string
	OpFee      float64
}

// var operationHistory repository.OperationHistoryRepositoryInterface = &repository.OperationHistoryRepository{}

func (b *botengine) BuyCoin(params *BuyCoinParams) bool {
	switch params.Exchange {
	case 1:
		// var BinanceInterface services.BinanceInterface = &services.Binance{}
		// var p bR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "BUY"
		// p.Type = "MARKET"
		// p.QuoteOrderQty = params.OpAmount

		// res := BinanceInterface.CreateOrder(p)

		// var pp repository.OperationHistoryDto

		// pp.Operation = params.Operation
		// pp.TransactionType = 1
		// pp.CoinPrice = utils.ParseFloat(res.Fills[0].Price)
		// pp.CoinQuantity = res.OrigQty
		// pp.StablePrice = utils.ParseFloat(res.CummulativeQuoteQty)
		// pp.StableQuantity = utils.ParseFloat(res.CummulativeQuoteQty)
		// pp.Fee = ((params.OpAmount - utils.ParseFloat(res.CummulativeQuoteQty)) / params.OpAmount)
		// pp.OperationExchangeId = res.OrderId

		// if !operationHistory.Create(tx, &pp) {
		// 	log.Panicln("BBC 02: ")
		// 	return false
		// }
	case 2:
		fmt.Printf("%+v\n", params)
		b.External.CreateOperationHistory(&pb.CreateOperationHistoryRequest{
			OperationHistory: &pb.OperationHistory{
				Operation:           params.Operation,
				TransactionType:     1,
				CoinPrice:           params.ClosePrice,
				CoinQuantity:        0,
				StablePrice:         utils.ParseFloat(params.OpAmount),
				StableQuantity:      utils.ParseFloat(params.OpAmount),
				Fee:                 params.OpFee,
				OperationExchangeId: fmt.Sprint(params.Operation),
			},
		})

		timestamp := b.Bybit.GetServerTimestamp()
		b.OrderChannel <- &bybitstructs.OrderRequest{
			ReqID: fmt.Sprint("Buy ", params.Operation),
			Header: bybitstructs.RequestHeader{
				Timestamp:  fmt.Sprint(timestamp),
				RecvWindow: "60000",
				Referer:    "testnet",
			},
			Op: "order.create",
			Args: []bybitstructs.OrderArgument{
				{
					Symbol:    params.Symbol,
					Side:      "Buy",
					OrderType: params.OrderType,
					Qty:       params.OpAmount,
					// Price:       fmt.Sprintf("%v", params.ClosePrice),
					Category:    "spot",
					TimeInForce: "GTC",
					OrderLinkId: fmt.Sprintf("%d-%v", params.Operation, timestamp),
				},
			},
		}
	case 3:
		// var p bBR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "BUY"
		// p.OrderType = "LIMIT"
		// p.OrderPrice = fmt.Sprintf("%v", params.ClosePrice)
		// p.OrderQty = fmt.Sprintf("%v", params.OpAmount)
		// p.TimeInForce = "GTC"

		// // res := bybittrade.BuyOrder(&p)

		// // if res.RetCode == 0 {
		// orderId, _ := strconv.ParseInt("132546", 10, 64)
		// var pp repository.OperationHistoryDto

		// pp.Operation = params.Operation
		// pp.TransactionType = 1
		// pp.CoinPrice = params.ClosePrice
		// pp.CoinQuantity = p.OrderQty
		// pp.StablePrice = params.OpAmount
		// pp.StableQuantity = params.OpAmount
		// pp.Fee = 0
		// pp.OperationExchangeId = orderId

		// if !operationHistory.Create(tx, &pp) {
		// 	log.Panicln("BBC 02: ")
		// 	return false
		// }
		// }
	}

	return true
}
