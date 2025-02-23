package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j *job) ManageTradeConfigStrategy(loopChannel *chan bool) {
	tradeConfig := j.ExchangeMS.ListTradeConfig()

	for _, strategy := range tradeConfig.GetTradeConfig() {
		if strategy.GetStrategy() == 3 {
			lastCandles := j.ExchangeMS.ListCandleLimit(&pb.ListCandleLimitRequest{
				Exchange: uint64(strategy.GetExchange()),
				Parity:   uint64(strategy.GetParity()),
				Limit:    2,
			})

			if len(lastCandles.GetCandles()) > 0 {
				var currentPrice, closedPrice float64
				for i, v := range lastCandles.GetCandles() {
					if i == 0 {
						currentPrice = v.GetClose()
					} else {
						closedPrice = v.GetClose()
					}
				}

				if currentPrice > (closedPrice * 1.00010) {
					j.ExchangeMS.UpdateTradeConfig(&pb.UpdateTradeConfigRequest{
						Enabled:  true,
						Strategy: 3,
						Parity:   1,
					})

					discord := discordService.NewDiscordWebhook()
					var notify discordstructs.Notification
					notify.ChannelUrl = "/1127722138414092431/h13r0Wy77BUwmrvZtWdtvTRlmMjVetXDMVI1VDaJF13zOg6iZUh888FGk7vrR2fwTOa-"
					notify.Content = fmt.Sprintf("(%v) Strategy %v(%v) is Enabled.", time.Now().Format("2006-01-02 15:04:05"), strategy.StrategyName, strategy.StrategyVariantName)
					discord.SendNotification(&notify)
				}
			}
		}
		*loopChannel <- true
	}
}
