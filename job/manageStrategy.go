package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j *job) createOperation(tradeConfig *pb.TradeConfig) {
	j.ExchangeMS.CreateOperation(&pb.CreateOperationRequest{
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

	discordService.NewDiscordWebhook().SendNotification(
		&discordstructs.Notification{
			ChannelUrl: "/1127722725318856714/31CT1E3rKDE-U1Ip-4nuYX_aIKJKGpKu9D3_p122FRMD5zbKs719t4YFdlV1LtTNDfHL",
			Content:    fmt.Sprintf("(%v) <---> Registering New Order! %v Strategy %v (%v)", time.Now().Format("2006-01-02 15:04:05"), tradeConfig.GetParitySymbol(), tradeConfig.GetStrategyName(), tradeConfig.GetStrategyVariantName()),
		},
	)
}

func (j *job) ManageTradeConfigStrategy(loopChannel *chan bool) {
	tradeConfigList := j.ExchangeMS.ListTradeConfig()
	operations := j.ExchangeMS.ListOperationEnabled(&pb.ListOperationEnabledRequest{})

	for _, tradeConfig := range tradeConfigList.GetTradeConfig() {
		if !tradeConfig.GetEnabled() {
			continue
		}

		if !tradeConfig.GetStrategyEnabled() {
			continue
		}

		hasOperation := false

		for _, operation := range operations.GetOperations() {
			if operation.GetUser() == tradeConfig.GetUser() &&
				operation.GetExchange() == tradeConfig.GetExchange() &&
				operation.GetParity() == tradeConfig.GetParity() &&
				operation.GetStrategy() == tradeConfig.GetStrategy() &&
				operation.GetStrategyVariant() == tradeConfig.GetStrategyVariant() {
				hasOperation = true
			}
		}

		if !hasOperation {
			j.createOperation(tradeConfig)
			hasOperation = false
		}

		// if tradeConfig.GetStrategy() == 3 {
		// 	lastCandles := j.ExchangeMS.ListCandleLimit(&pb.ListCandleLimitRequest{
		// 		Exchange: uint64(tradeConfig.GetExchange()),
		// 		Parity:   uint64(tradeConfig.GetParity()),
		// 		Limit:    2,
		// 	})

		// 	if len(lastCandles.GetCandles()) > 0 {
		// 		var currentPrice, closedPrice float64
		// 		for i, v := range lastCandles.GetCandles() {
		// 			if i == 0 {
		// 				currentPrice = v.GetClose()
		// 			} else {
		// 				closedPrice = v.GetClose()
		// 			}
		// 		}

		// 		if currentPrice > (closedPrice * 1.00010) {
		// 			j.ExchangeMS.UpdateTradeConfig(&pb.UpdateTradeConfigRequest{
		// 				TradeConfig:             uint64(tradeConfig.GetTradeConfig()),
		// 				User:                    uint64(tradeConfig.GetUser()),
		// 				Modality:                uint64(tradeConfig.GetModality()),
		// 				Strategy:                uint64(tradeConfig.GetStrategy()),
		// 				StrategyVariant:         uint64(tradeConfig.GetStrategyVariant()),
		// 				Parity:                  uint64(tradeConfig.GetParity()),
		// 				Exchange:                uint64(tradeConfig.GetExchange()),
		// 				OperationQuantity:       uint64(tradeConfig.GetOperationQuantity()),
		// 				OperationAmount:         tradeConfig.GetOperationAmount(),
		// 				DefaultProfitPercentage: tradeConfig.GetDefaultProfitPercentage(),
		// 				WalletValueLimit:        tradeConfig.GetWalletValueLimit(),
		// 				Enabled:                 true,
		// 			})

		// 			discord := discordService.NewDiscordWebhook()
		// 			var notify discordstructs.Notification
		// 			notify.ChannelUrl = "/1127722138414092431/h13r0Wy77BUwmrvZtWdtvTRlmMjVetXDMVI1VDaJF13zOg6iZUh888FGk7vrR2fwTOa-"
		// 			notify.Content = fmt.Sprintf("(%v) Strategy %v(%v) is Enabled.", time.Now().Format("2006-01-02 15:04:05"), tradeConfig.StrategyName, tradeConfig.StrategyVariantName)
		// 			discord.SendNotification(&notify)
		// 		}
		// 	}
		// }
	}

	*loopChannel <- true
}
