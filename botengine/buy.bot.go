package botengine

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
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
		fmt.Println(params)
		res := b.Bybit.CreateOrder(&bybitstructs.OrderParams{
			Category:    "spot",
			Symbol:      params.Symbol,
			Side:        "Buy",
			OrderType:   params.OrderType,
			OrderPrice:  fmt.Sprintf("%v", params.ClosePrice),
			OrderQty:    params.OpAmount,
			TimeInForce: "GTC",
			// OrderLinkId: fmt.Sprintf("%d-BUY", params.Operation),
		})

		log.Println("BUY ", res, " -> OP: ", params.Operation, params.Symbol)

		if res.RetCode == 0 {
			orderId, _ := strconv.ParseInt(res.Data.OrderID, 10, 64)
			opAmount, _ := strconv.ParseFloat(params.OpAmount, 64)

			b.External.CreateOperationHistory(&pb.CreateOperationHistoryRequest{
				OperationHistory: &pb.OperationHistory{
					Operation:           params.Operation,
					TransactionType:     1,
					CoinPrice:           params.ClosePrice,
					CoinQuantity:        utils.Truncate((opAmount/params.ClosePrice)*(1-params.OpFee), 6),
					StablePrice:         opAmount,
					StableQuantity:      opAmount,
					Fee:                 params.OpFee,
					OperationExchangeId: uint64(orderId),
				},
			})
		} else {
			discordService.NewDiscordWebhook().SendNotification(&discordstructs.Notification{
				ChannelUrl: "/1127721886214795325/h00kYmfUrFTAIIIJasJi2MOnD0jPRmja_NYXoC1gsxSv4pxKKFBQGWsr9T0FDGIf_RCk",
				Content:    fmt.Sprintf("(%v) <---> ERR Buy %v: Send %v <----> Get %v ", time.Now().Format("2006-01-02 15:04:05"), params.Operation, params.Operation, res),
			})
			return false
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
