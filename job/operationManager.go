package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j *job) OperationManager(loopChannel *chan bool) {
	operations := j.ExchangeMS.ListAllOperation(&pb.ListAllOperationRequest{})
	fmt.Println(operations)
	for _, op := range operations.GetOperations() {
		tradeConfig := j.ExchangeMS.GetTradeConfig(&pb.GetTradeConfigRequest{
			User:            op.GetUser(),
			Strategy:        op.GetStrategy(),
			StrategyVariant: op.GetStrategyVariant(),
			Parity:          op.GetParity(),
			Exchange:        op.GetExchange(),
		})

		queueName := ""
		if op.GetParity() == 1 {
			queueName = "kline.1.BTCUSDT"
		}
		currentPrice := 0.0
		msg := j.RabbitMQ.Listen("klines", queueName)
		select {
		case msg := <-msg:
			currentPrice = float64(msg.Body[0])
		}

		if op.GetEnabled() && !op.GetClosed() { // Desable
			if currentPrice < (op.GetOpenPrice() * (1 - 0.013)) {
				j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
					Operation: &pb.Operation{
						Operation:       op.GetOperation(),
						User:            op.GetUser(),
						Parity:          op.GetParity(),
						Exchange:        op.GetExchange(),
						Strategy:        op.GetStrategy(),
						StrategyVariant: op.GetStrategyVariant(),
						MtsStart:        op.GetMtsStart(),
						MtsFinish:       op.GetMtsFinish(),
						OpenPrice:       op.GetOpenPrice(),
						ClosePrice:      op.GetClosePrice(),
						InvestedAmount:  op.GetInvestedAmount(),
						ProfitAmount:    op.GetProfitAmount(),
						Profit:          op.GetProfit(),
						Closed:          op.GetClosed(),
						Audit:           op.GetAudit(),
						Enabled:         false,
						TimesCanceled:   op.GetTimesCanceled(),
					},
				})
			}

			if tradeConfig.TradeConfig.GetEnabled() && tradeConfig.TradeConfig.GetOperationQuantity() > 0 {
				j.ExchangeMS.UpdateTradeConfig(&pb.UpdateTradeConfigRequest{
					TradeConfig:             uint64(tradeConfig.TradeConfig.GetTradeConfig()),
					User:                    uint64(tradeConfig.TradeConfig.GetUser()),
					Modality:                uint64(tradeConfig.TradeConfig.GetModality()),
					Strategy:                uint64(tradeConfig.TradeConfig.GetStrategy()),
					StrategyVariant:         uint64(tradeConfig.TradeConfig.GetStrategyVariant()),
					Parity:                  uint64(tradeConfig.TradeConfig.GetParity()),
					Exchange:                uint64(tradeConfig.TradeConfig.GetExchange()),
					OperationQuantity:       uint64(tradeConfig.TradeConfig.GetOperationQuantity() - 1),
					OperationAmount:         tradeConfig.TradeConfig.GetOperationAmount(),
					DefaultProfitPercentage: tradeConfig.TradeConfig.GetDefaultProfitPercentage(),
					WalletValueLimit:        tradeConfig.TradeConfig.GetWalletValueLimit(),
					Enabled:                 tradeConfig.TradeConfig.GetEnabled(),
				})
			}

			discord := discordService.NewDiscordWebhook()
			discord.SendNotification(&discordstructs.Notification{
				ChannelUrl: "",
				Content:    fmt.Sprintf("(%v) Operation %v with OpenPrice -> %v was Desabled.", time.Now().Format("2006-01-02 15:04:05"), op.GetOperation(), op.GetOpenPrice()),
			})
		} else { // Enable
			if op.GetOpenPrice() <= currentPrice && op.GetOpenPrice() != 0 && !op.GetClosed() && op.GetEnabled() {
				j.ExchangeMS.UpdateOperation(&pb.UpdateOperationRequest{
					Operation: &pb.Operation{
						Operation:       op.GetOperation(),
						User:            op.GetUser(),
						Parity:          op.GetParity(),
						Exchange:        op.GetExchange(),
						Strategy:        op.GetStrategy(),
						StrategyVariant: op.GetStrategyVariant(),
						MtsStart:        op.GetMtsStart(),
						MtsFinish:       op.GetMtsFinish(),
						OpenPrice:       op.GetOpenPrice(),
						ClosePrice:      op.GetClosePrice(),
						InvestedAmount:  op.GetInvestedAmount(),
						ProfitAmount:    op.GetProfitAmount(),
						Profit:          op.GetProfit(),
						Closed:          op.GetClosed(),
						Audit:           op.GetAudit(),
						Enabled:         true,
						TimesCanceled:   op.GetTimesCanceled(),
					},
				})
			}

			if tradeConfig.GetTradeConfig().GetEnabled() {
				j.ExchangeMS.UpdateTradeConfig(&pb.UpdateTradeConfigRequest{
					TradeConfig:             uint64(tradeConfig.TradeConfig.GetTradeConfig()),
					User:                    uint64(tradeConfig.TradeConfig.GetUser()),
					Modality:                uint64(tradeConfig.TradeConfig.GetModality()),
					Strategy:                uint64(tradeConfig.TradeConfig.GetStrategy()),
					StrategyVariant:         uint64(tradeConfig.TradeConfig.GetStrategyVariant()),
					Parity:                  uint64(tradeConfig.TradeConfig.GetParity()),
					Exchange:                uint64(tradeConfig.TradeConfig.GetExchange()),
					OperationQuantity:       uint64(tradeConfig.TradeConfig.GetOperationQuantity() + 1),
					OperationAmount:         tradeConfig.TradeConfig.GetOperationAmount(),
					DefaultProfitPercentage: tradeConfig.TradeConfig.GetDefaultProfitPercentage(),
					WalletValueLimit:        tradeConfig.TradeConfig.GetWalletValueLimit(),
					Enabled:                 tradeConfig.TradeConfig.GetEnabled(),
				})
			}

			discord := discordService.NewDiscordWebhook()
			discord.SendNotification(&discordstructs.Notification{
				ChannelUrl: "",
				Content:    fmt.Sprintf("(%v) Operation %v with OpenPrice -> %v was Desabled.", time.Now().Format("2006-01-02 15:04:05"), op.GetOperation(), op.GetOpenPrice()),
			})
		}
	}

	*loopChannel <- true
}
