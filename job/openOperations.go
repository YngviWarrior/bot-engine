package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j *job) OpenOperationManager(loopChannel *chan bool) {
	operations := j.ExchangeMS.ListAllOperation(&pb.ListAllOperationRequest{})

	for _, op := range operations.Operations {
		tradeConfig := j.ExchangeMS.GetTradeConfig(&pb.GetTradeConfigRequest{
			User:            op.GetUser(),
			Strategy:        op.GetStrategy(),
			StrategyVariant: op.GetStrategyVariant(),
			Parity:          op.GetParity(),
			Exchange:        op.GetExchange(),
			Modality:        1,
		})

		// fmt.Println("AQUI1")
		if op.Enabled && !op.Closed && op.OpenPrice > 0 {
			// fmt.Println("AQUI2")
			sellQuote := utils.FindSellCotation(op.InvestedAmount, (op.InvestedAmount / op.OpenPrice), 0.001, op.OpenPrice, tradeConfig.TradeConfig.DefaultProfitPercentage)

			discord := discordService.NewDiscordWebhook()
			discord.SendNotification(&discordstructs.Notification{
				ChannelUrl: "/1127722360087253084/8_w1-Ffb7gnbtREySUL2z0RDfwDQnVEIQCa-pUMouJQoib7FFiGOo3ieDjcyOZobPPNs",
				Content:    fmt.Sprintf("(%v) Op: %v with Strategy %v(%v) is opened. Purchase Price: %v // Profit Percentage: %v // Quote to Sale: %v ", time.Now().Format("2006-01-02 15:04:05"), op.Operation, op.StrategyName, op.StrategyVariantName, op.OpenPrice, tradeConfig.TradeConfig.DefaultProfitPercentage, sellQuote),
			})
		}
	}

	*loopChannel <- true
}
