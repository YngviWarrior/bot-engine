package botengine

import (
	"fmt"
	"time"

	external "github.com/YngviWarrior/bot-engine/infra/external"
	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

type BotEngine interface {
	InitBotEngine(kline rabbitmq.CombinedData)
}

type botengine struct {
	Bybit    bybitSDK.BybitServiceInterface
	External external.ExternalInterface

	OrderChannel chan *bybitstructs.OrderRequest
}

func NewBotEngine(bybit bybitSDK.BybitServiceInterface, external external.ExternalInterface, orderChannel chan *bybitstructs.OrderRequest) BotEngine {
	return &botengine{
		Bybit:    bybit,
		External: external,

		OrderChannel: orderChannel,
	}
}

func (b *botengine) InitBotEngine(kline rabbitmq.CombinedData) {
	start := time.Now()

	operations := b.External.ListOperationEnabled(&pb.ListOperationEnabledRequest{})
	// fmt.Println("Operations: ", operations.Operations)
	for _, operation := range operations.GetOperations() {
		if operation.Strategy == 1 {
			if kline.Topic == "kline.1.BTCUSDT" {
				b.ByBitAvgPriceDay(operation, &kline, nil)
				fmt.Println("ByBitAvgPriceDay BTCUSDT")
			}
		} else if operation.Strategy == 2 {
			if kline.Topic == "kline.1.BTCUSDT" {
				b.ByBitOpenClose(operation, &kline, nil)
				fmt.Println("ByBitOpenClose BTCUSDT")
			}
		}
	}

	ds := discordService.NewDiscordWebhook()
	ds.SendNotification(&discordstructs.Notification{
		ChannelUrl: "/1127722138414092431/h13r0Wy77BUwmrvZtWdtvTRlmMjVetXDMVI1VDaJF13zOg6iZUh888FGk7vrR2fwTOa-",
		Content:    fmt.Sprintf("🕒 [%s] Tempo de execução do loop de estratégias: %v\n", time.Now().Format("2006-01-02 15:04:05"), time.Since(start)),
	})

}
