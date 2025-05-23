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
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (b *botengine) ByBitOpenClose(configs *pb.TradeConfig, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	b.OperateOpenClose(configs, kline, wg)
}

func (b *botengine) OperateOpenClose(tradeConfig *pb.TradeConfig, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	var newRegister bool
	var stableCoin uint64
	switch tradeConfig.GetParity() {
	case 1, 4:
		stableCoin = 1
	case 3:
		stableCoin = 2
	}

	wallet := b.External.GetWalletWithCoin(&pb.GetWalletWithCoinRequest{
		Exchange: tradeConfig.Exchange,
		Coin:     stableCoin,
	})
	if wallet == nil {
		log.Panicln("OC 00: ", "Wallet not found")
	}
	fmt.Println(wallet)

	operations := b.External.ListOperation(&pb.ListOperationRequest{
		User:            tradeConfig.GetUser(),
		Strategy:        tradeConfig.GetStrategy(),
		StrategyVariant: tradeConfig.GetStrategyVariant(),
		Parity:          tradeConfig.GetParity(),
		Exchange:        tradeConfig.GetExchange(),
		Closed:          false,
		Enabled:         true,
		InTransaction:   false,
	})

	fmt.Println("Operations: ", operations.Operations)

	// fmt.Println("Open Operations: ", len(operations.Operations)
	if len(operations.GetOperations()) < int(tradeConfig.GetOperationQuantity()) && wallet.GetWallet().GetAmount() > (tradeConfig.GetOperationAmount()+tradeConfig.GetWalletValueLimit()) {
		op := b.External.CreateOperation(&pb.CreateOperationRequest{
			Operation: &pb.Operation{
				User:            tradeConfig.GetUser(),
				Parity:          tradeConfig.GetParity(),
				Exchange:        tradeConfig.GetExchange(),
				Strategy:        tradeConfig.GetStrategy(),
				StrategyVariant: tradeConfig.GetStrategyVariant(),
				MtsStart:        uint64(time.Now().UnixMicro()),
				InvestedAmount:  tradeConfig.GetOperationAmount(),
				Enabled:         true,
			},
		})

		if op.Operation.Operation == 0 {
			log.Panicln("OC 04: ")
		}

		discordService.NewDiscordWebhook().SendNotification(
			&discordstructs.Notification{
				ChannelUrl: "/1127722725318856714/31CT1E3rKDE-U1Ip-4nuYX_aIKJKGpKu9D3_p122FRMD5zbKs719t4YFdlV1LtTNDfHL",
				Content:    fmt.Sprintf("(%v) <---> Registering New Order! %v Strategy %v (%v)", time.Now().Format("2006-01-02 15:04:05"), tradeConfig.GetParitySymbol(), tradeConfig.GetStrategyName(), tradeConfig.GetStrategyVariantName()),
			},
		)

		newRegister = true
	}

	for _, operation := range operations.GetOperations() {
		lastRegister := b.External.GetLastBuyRegisterByOperation(&pb.GetLastBuyRegisterByOperationRequest{
			Operation: operation.Operation,
		})
		fmt.Println("lastRegister: ", lastRegister)

		sellCotation := utils.FindSellCotation(operation.InvestedAmount, lastRegister.CoinQuantity, lastRegister.Fee, operation.OpenPrice, tradeConfig.GetDefaultProfitPercentage())

		var newestRegister bool = (lastRegister.Status == 0)
		currentPrice := utils.ParseFloat(kline.Data[0].Close)
		var isProfitable bool = (currentPrice >= sellCotation) && sellCotation != 0
		// fmt.Printf("%v OP: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.Operation, sellCotation)
		if wallet.GetWallet().GetAmount() > (tradeConfig.GetOperationAmount()+tradeConfig.GetWalletValueLimit()) && newestRegister && !operation.InTransaction {
			var bc BuyCoinParams

			bc.Operation = operation.Operation
			bc.ClosePrice = currentPrice
			bc.Exchange = tradeConfig.GetExchange()
			bc.Symbol = tradeConfig.GetParitySymbol()
			bc.OrderType = "Market"
			bc.OpFee = 0.001

			switch operation.Parity {
			case 1: // BTCUSDT
				bc.OpAmount = fmt.Sprintf("%f", tradeConfig.GetOperationAmount())
			case 4: // ETHUSDT
				bc.OpAmount = fmt.Sprintf("%f", tradeConfig.GetOperationAmount())
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
							Profit:          tradeConfig.GetDefaultProfitPercentage(),
							ProfitAmount:    operation.GetProfitAmount(),
							Closed:          false,
							Audit:           operation.GetAudit(),
							Enabled:         false,
							InTransaction:   true,
						},
					})

					walletUpdate := &pb.Wallet{
						Wallet:   wallet.GetWallet().GetWallet(),
						Exchange: wallet.GetWallet().GetExchange(),
						Coin:     wallet.GetWallet().GetCoin(),
						Amount:   (wallet.GetWallet().GetAmount() - tradeConfig.GetOperationAmount()),
					}

					b.External.UpdateWallet(&pb.UpdateWalletRequest{
						Wallet: []*pb.Wallet{walletUpdate},
					})

					fmt.Println("OP: ", operation.Operation)

					newRegister = true
				}
			}
		} else if isProfitable && lastRegister.CoinQuantity != 0 {
			var p SellCoinParams

			p.Operation = operation.Operation
			p.ClosePrice = currentPrice
			p.Exchange = tradeConfig.GetExchange()
			p.OpAmount = operation.InvestedAmount
			p.Symbol = tradeConfig.GetParitySymbol()
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
							InvestedAmount:  operation.InvestedAmount * (1 + tradeConfig.DefaultProfitPercentage),
							OpenPrice:       operation.OpenPrice,
							ClosePrice:      currentPrice,
							Profit:          tradeConfig.GetDefaultProfitPercentage(),
							ProfitAmount:    operation.ProfitAmount,
							Closed:          false,
							Audit:           operation.Audit,
							Enabled:         false,
							InTransaction:   true,
						},
					})

					walletUpdate := &pb.Wallet{
						Wallet:   wallet.GetWallet().GetWallet(),
						Exchange: wallet.GetWallet().GetExchange(),
						Coin:     wallet.GetWallet().GetCoin(),
						Amount:   (wallet.GetWallet().GetAmount() + tradeConfig.GetOperationAmount()),
					}

					b.External.UpdateWallet(&pb.UpdateWalletRequest{
						Wallet: []*pb.Wallet{walletUpdate},
					})

					newRegister = true
				}
			}
		}
	}

	if newRegister {
		time.Sleep(time.Millisecond * 500)
	}
}
