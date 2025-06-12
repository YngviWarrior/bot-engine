package botengine

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	"github.com/YngviWarrior/bot-engine/utils"
)

func (b *botengine) ByBitOpenClose(operation *pb.OperationJoin, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	b.OperateOpenClose(operation, kline, wg)
}

func (b *botengine) OperateOpenClose(operation *pb.OperationJoin, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	wallet := b.External.GetWalletWithCoin(&pb.GetWalletWithCoinRequest{
		Exchange: 2,
		Coin:     1,
	})
	if wallet == nil {
		log.Panicln("BBOC 00: ", "Wallet not found")
	}
	// fmt.Println(wallet)

	lastRegister := b.External.GetLastBuyRegisterByOperation(&pb.GetLastBuyRegisterByOperationRequest{
		Operation: operation.Operation,
	})
	// fmt.Println("lastRegister: ", lastRegister)

	sellCotation := utils.FindSellCotation(operation.InvestedAmount, lastRegister.CoinQuantity, lastRegister.Fee, operation.OpenPrice, "0.001")

	var newestRegister bool = (lastRegister.Status == 0)
	currentPrice := utils.ParseFloat(kline.Data[0].Close)
	var isProfitable bool = (currentPrice >= sellCotation) && sellCotation != 0

	walletAmount, err := strconv.ParseFloat(wallet.GetWallet().GetAmount(), 64)
	if err != nil {
		log.Panic("BBOC 01: ", err)
	}

	fmt.Printf("%v OP: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.Operation, sellCotation)
	if walletAmount > (50+1000) && newestRegister && !operation.InTransaction {
		var bc BuyCoinParams

		bc.Operation = operation.Operation
		bc.ClosePrice = fmt.Sprintf("%0.2f", currentPrice)
		bc.Exchange = 2
		bc.Symbol = "BTCUSDT"
		bc.OrderType = "Limit"
		bc.OpFee = "0.001"

		switch operation.Parity {
		case 1: // BTCUSDT
			bc.OpAmount = fmt.Sprintf("%f", 50.0/currentPrice)
		case 4: // ETHUSDT
			bc.OpAmount = fmt.Sprintf("%f", 50.0/currentPrice)
		}

		if b.BuyCoin(&bc) {
			b.External.UpdateOperation(&pb.UpdateOperationRequest{
				Operation: &pb.Operation{
					Operation:       operation.GetOperation(),
					User:            operation.GetUser(),
					Parity:          operation.GetParity(),
					Exchange:        operation.GetExchange(),
					Strategy:        operation.GetStrategy(),
					StrategyVariant: operation.GetStrategyVariant(),
					MtsStart:        operation.GetMtsStart(),
					InvestedAmount:  operation.GetInvestedAmount(),
					OpenPrice:       fmt.Sprintf("%0.2f", currentPrice),
					ClosePrice:      operation.GetClosePrice(),
					Profit:          "0.001",
					ProfitAmount:    operation.GetProfitAmount(),
					Closed:          false,
					Audit:           operation.GetAudit(),
					Enabled:         true,
					InTransaction:   true,
				},
			})

			walletUpdate := &pb.Wallet{
				Wallet:   wallet.GetWallet().GetWallet(),
				Exchange: wallet.GetWallet().GetExchange(),
				Coin:     wallet.GetWallet().GetCoin(),
				Amount:   fmt.Sprintf("%0.2f", walletAmount-50.0),
			}

			b.External.UpdateWallet(&pb.UpdateWalletRequest{
				Wallet: []*pb.Wallet{walletUpdate},
			})
		}
	} else if isProfitable && lastRegister.CoinQuantity != "0" {
		var p SellCoinParams

		investedAmount, err := strconv.ParseFloat(operation.GetInvestedAmount(), 64)
		if err != nil {
			log.Panic("BBOC 03: ", err)
		}

		p.Operation = operation.Operation
		p.ClosePrice = fmt.Sprintf("%0.2f", currentPrice)
		p.Exchange = 2
		p.OpAmount = fmt.Sprintf("%0.11f", investedAmount/currentPrice)
		p.Symbol = "BTCUSDT"
		p.OrderType = "Limit"
		p.OpFee = "0.001"

		coinQuantity, err := strconv.ParseFloat(lastRegister.CoinQuantity, 64)
		if err != nil {
			log.Panic("BBOC 04: ", err)
		}

		switch operation.Parity {
		case 1: // BTCUSDT
			p.CoinQuantity = fmt.Sprintf("%0.11f", utils.Truncate(coinQuantity, 5))
		case 4: // ETHUSDT
			p.CoinQuantity = fmt.Sprintf("%0.11f", utils.Truncate(coinQuantity, 4))
		}

		if b.SellCoin(&p) {
			b.External.UpdateOperation(&pb.UpdateOperationRequest{
				Operation: &pb.Operation{
					Operation:       operation.Operation,
					User:            operation.User,
					Parity:          operation.Parity,
					Exchange:        operation.Exchange,
					Strategy:        operation.Strategy,
					StrategyVariant: operation.StrategyVariant,
					MtsStart:        operation.MtsStart,
					MtsFinish:       uint64(time.Now().UnixMicro()),
					InvestedAmount:  fmt.Sprintf("%0.11f", investedAmount*(1+0.001)),
					OpenPrice:       operation.OpenPrice,
					ClosePrice:      fmt.Sprintf("%0.2f", currentPrice),
					Profit:          "0.001",
					ProfitAmount:    operation.ProfitAmount,
					Closed:          true,
					Audit:           operation.Audit,
					Enabled:         false,
					InTransaction:   true,
				},
			})

			walletUpdate := &pb.Wallet{
				Wallet:   wallet.GetWallet().GetWallet(),
				Exchange: wallet.GetWallet().GetExchange(),
				Coin:     wallet.GetWallet().GetCoin(),
				Amount:   fmt.Sprintf("%0.2f", walletAmount+50.0),
			}

			b.External.UpdateWallet(&pb.UpdateWalletRequest{
				Wallet: []*pb.Wallet{walletUpdate},
			})
		}
	}
}
