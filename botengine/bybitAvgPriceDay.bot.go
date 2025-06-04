package botengine

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	"github.com/YngviWarrior/bot-engine/utils"
)

func (b *botengine) ByBitAvgPriceDay(operation *pb.OperationJoin, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	b.OperateAvgPriceDay(operation, kline, wg)
}

func (b *botengine) OperateAvgPriceDay(operation *pb.OperationJoin, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	var newRegister bool

	wallet := b.External.GetWalletWithCoin(&pb.GetWalletWithCoinRequest{
		Exchange: 2,
		Coin:     1,
	})
	if wallet == nil {
		log.Panicln("OC 00: ", "Wallet not found")
	}
	// fmt.Println(wallet)

	avg := b.External.GetAvgPriceByParityExchange(&pb.GetAvgPriceByParityExchangeRequest{
		Parity:   1,
		Exchange: 2,
	})

	fmt.Println("Average Price: ", avg)

	lastRegister := b.External.GetLastBuyRegisterByOperation(&pb.GetLastBuyRegisterByOperationRequest{
		Operation: operation.Operation,
	})
	// fmt.Println("lastRegister: ", lastRegister)

	sellCotation := utils.FindSellCotation(operation.InvestedAmount, lastRegister.CoinQuantity, lastRegister.Fee, operation.OpenPrice, 0.001)

	var newestRegister bool = (lastRegister.Status == 0)
	currentPrice := utils.ParseFloat(kline.Data[0].Close)
	var isProfitable bool = (currentPrice >= sellCotation) && sellCotation != 0
	// isProfitable = true // For testing purposes, set isProfitable to true

	fmt.Printf("%v AVG Day Op: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.GetOperation(), sellCotation)
	fmt.Println(currentPrice <= avg.GetDay(), "Current Price: ", currentPrice, "Average Day: ", avg.GetDay())
	if wallet.GetWallet().GetAmount() > (50+1000) && newestRegister && !operation.InTransaction && currentPrice <= avg.GetDay() {
		var bc BuyCoinParams

		bc.Operation = operation.Operation
		bc.ClosePrice = currentPrice
		bc.Exchange = 2
		bc.Symbol = "BTCUSDT"
		bc.OrderType = "Market"
		bc.OpFee = 0.001

		switch operation.Parity {
		case 1: // BTCUSDT
			bc.OpAmount = fmt.Sprintf("%f", 50.0)
		case 4: // ETHUSDT
			bc.OpAmount = fmt.Sprintf("%f", 50.0)
		}

		switch os.Getenv("ENVIROMENT") {
		case "server", "local", "testnet":
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
						OpenPrice:       currentPrice,
						ClosePrice:      operation.GetClosePrice(),
						Profit:          0.001,
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
					Amount:   (wallet.GetWallet().GetAmount() - 50.0),
				}

				b.External.UpdateWallet(&pb.UpdateWalletRequest{
					Wallet: []*pb.Wallet{walletUpdate},
				})

				newRegister = true
			}
		}
	} else if isProfitable && lastRegister.CoinQuantity != 0 && operation.InTransaction {
		var p SellCoinParams

		p.Operation = operation.Operation
		p.ClosePrice = currentPrice
		p.Exchange = 2
		p.OpAmount = operation.InvestedAmount
		p.Symbol = "BTCUSDT"
		p.OrderType = "Market"
		p.OpFee = 0.001

		switch operation.Parity {
		case 1: // BTCUSDT
			p.CoinQuantity = utils.Truncate(lastRegister.CoinQuantity, 5)
		case 4: // ETHUSDT
			p.CoinQuantity = utils.Truncate(lastRegister.CoinQuantity, 4)
		}

		switch os.Getenv("ENVIROMENT") {
		case "server", "local", "testnet":
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
						InvestedAmount:  operation.InvestedAmount,
						OpenPrice:       operation.OpenPrice,
						ClosePrice:      currentPrice,
						Profit:          0.001,
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
					Amount:   (wallet.GetWallet().GetAmount() + 50.0),
				}

				b.External.UpdateWallet(&pb.UpdateWalletRequest{
					Wallet: []*pb.Wallet{walletUpdate},
				})

				newRegister = true
			}
		}
	}

	if newRegister {
		time.Sleep(time.Millisecond * 500)
	}

	// wg.Done()
}
