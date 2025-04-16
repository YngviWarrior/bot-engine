package botengine

import (
	"database/sql"
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

func (b *botengine) BuyCoin(tx *sql.Tx, params *BuyCoinParams) bool {
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
		// fmt.Println(params)
		// p.Symbol = params.Symbol
		// p.Side = "BUY"
		// p.OrderType = params.OrderType
		// p.OrderPrice = fmt.Sprintf("%v", params.ClosePrice)
		// p.OrderQty = params.OpAmount
		// p.TimeInForce = "GTC"
		// // fmt.Println(p)

		// res := b.ByBit.CreateOrder(p)

		// log.Println("BUY ", res, " -> OP: ", params.Operation, params.Symbol)

		// if res.RetCode == 0 {
		// 	orderId, _ := strconv.ParseInt(res.Result.OrderID, 10, 64)
		// 	var pp repository.OperationHistoryDto
		// 	opAmount, _ := strconv.ParseFloat(params.OpAmount, 64)

		// 	pp.Operation = params.Operation
		// 	pp.TransactionType = 1
		// 	pp.CoinPrice = params.ClosePrice
		// 	pp.CoinQuantity = fmt.Sprintf("%.7f", utils.Truncate((opAmount/params.ClosePrice)*(1-params.OpFee), 6))
		// 	pp.StablePrice = opAmount
		// 	pp.StableQuantity = opAmount
		// 	pp.Fee = params.OpFee
		// 	pp.OperationExchangeId = orderId

		// 	if !operationHistory.Create(tx, &pp) {
		// 		log.Panicln("BBC 02: ")
		// 		return false
		// 	}
		// } else {
		// 	n.ChannelUrl = constants.DISCORD_ERROR_TEST
		// 	n.Content = fmt.Sprintf("(%v) <---> ERR Buy %v: Send %v <----> Get %v ", time.Now().Format("2006-01-02 15:04:05"), params.Operation, p, res)
		// 	b.Discord.SendNotification(&n)
		// 	return false
		// }
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
