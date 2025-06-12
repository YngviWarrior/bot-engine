package botengine

import (
	"fmt"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

type SellCoinParams struct {
	Operation    uint64
	ClosePrice   string
	CoinQuantity string
	Exchange     uint64
	OpAmount     string
	OpFee        string
	Symbol       string
	OrderType    string
}

func (b *botengine) SellCoin(params *SellCoinParams) bool {
	switch params.Exchange {
	case 1: //Binance
		// var BinanceInterface services.BinanceInterface = &services.Binance{}
		// var p bR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "SELL"
		// p.Type = "MARKET"
		// p.Quantity = params.CoinQuantity

		// res := BinanceInterface.CreateOrder(p)
		// log.Println("SELL ", res)

		// if res.ClientOrderId == "" {
		// 	// *channel <- true
		// 	return false
		// }

		// var pp repository.OperationHistoryDto

		// pp.Operation = params.Operation
		// pp.TransactionType = 2
		// pp.CoinPrice = utils.ParseFloat(res.Fills[0].Price)
		// pp.CoinQuantity = res.OrigQty
		// pp.StablePrice = utils.ParseFloat(res.CummulativeQuoteQty)
		// pp.StableQuantity = utils.ParseFloat(res.CummulativeQuoteQty)
		// pp.Fee = ((params.OpAmount - utils.ParseFloat(res.CummulativeQuoteQty)) / params.OpAmount)
		// pp.OperationExchangeId = res.OrderId

		// if !operationHistory.Create(tx, &pp) {
		// 	log.Panicln("SBC 02: ")
		// 	return false
		// }
	case 2: //ByBit
		fmt.Printf("%+v\n", params)
		b.External.CreateOperationHistory(&pb.CreateOperationHistoryRequest{
			OperationHistory: &pb.OperationHistory{
				Operation:               params.Operation,
				TransactionType:         2,
				CoinPrice:               params.ClosePrice,
				CoinQuantity:            params.CoinQuantity,
				StablePrice:             params.OpAmount,
				StableQuantity:          params.OpAmount,
				Fee:                     params.OpFee,
				OperationExchangeId:     fmt.Sprint(params.Operation),
				OperationExchangeStatus: 1,
			},
		})

		timestamp := b.Bybit.GetServerTimestamp()
		b.OrderChannel <- &bybitstructs.OrderRequest{
			ReqID: fmt.Sprint("Sell ", params.Operation),
			Header: bybitstructs.RequestHeader{
				Timestamp:  fmt.Sprint(timestamp),
				RecvWindow: "60000",
				Referer:    "demo",
			},
			Op: "order.create",
			Args: []bybitstructs.OrderArgument{
				{
					Symbol:      params.Symbol,
					Side:        "Sell",
					OrderType:   params.OrderType,
					Qty:         fmt.Sprint(params.CoinQuantity),
					Price:       fmt.Sprintf("%v", params.ClosePrice),
					Category:    "spot",
					TimeInForce: "GTC",
					OrderLinkId: fmt.Sprintf("%d-%v", params.Operation, timestamp),
				},
			},
		}
	case 3: //TEST
		// var p bBR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "SELL"
		// p.OrderType = "LIMIT"
		// p.OrderPrice = fmt.Sprintf("%v", params.ClosePrice)
		// p.OrderQty = fmt.Sprintf("%v", params.CoinQuantity)
		// p.TimeInForce = "GTC"

		// // res := bybittrade.SellOrder(&p)
		// // log.Println("SELL ", res)

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
		// 	log.Panicln("SBC 02: ")
		// 	return false
		// }
		// // }
	}

	return true
}
