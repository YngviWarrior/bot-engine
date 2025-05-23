package botengine

import (
	"fmt"
	"time"

	external "github.com/YngviWarrior/bot-engine/infra/external"
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
	tradeConfigList := b.External.ListTradeConfig()
	var start time.Time
	start = time.Now()

	for _, configs := range tradeConfigList.GetTradeConfig() {
		if configs.Enabled {
			switch configs.Exchange {
			case BYBIT_EXCH:
				if configs.Enabled {
					switch configs.Modality {
					case SPOT_MODALITY:
						if configs.StrategyEnabled {
							switch configs.Strategy {
							case AVERAGE_PRICE:
								switch configs.StrategyVariant {
								case AVERAGE_PRICE_DAY:
									if configs.StrategyVariantEnabled {
										if kline.Topic == "kline.1.BTCUSDT" && configs.ParitySymbol == "BTCUSDT" {
											b.ByBitAvgPriceDay(configs, &kline, nil)
											fmt.Println("ByBitAvgPriceDay BTCUSDT")
										}
										if kline.Topic == "kline.1.ETHUSDT" && configs.ParitySymbol == "ETHUSDT" {
											b.ByBitAvgPriceDay(configs, &kline, nil)
										}
									}
								case AVERAGE_PRICE_WEEK:
									if configs.StrategyVariantEnabled {
										//  b.ByBitAvgPriceWeek(configs, kline, &nil)
									}
								}
							case CLOSE_OPEN:
								if configs.StrategyEnabled {
									if kline.Topic == "kline.1.BTCUSDT" && configs.ParitySymbol == "BTCUSDT" {
										b.ByBitOpenClose(configs, &kline, nil)
										fmt.Println("BybitOpenClose BTCUSDT")
									}
								}
							case FAST_TRADE:
								if configs.StrategyEnabled {
									//  b.ByBitFastTrade(configs, kline, &wg)
								}
							}
						}
					}
				}
			}
		}
	}

	if time.Since(start) > (time.Second * 1) {
		ds := discordService.NewDiscordWebhook()
		ds.SendNotification(&discordstructs.Notification{
			ChannelUrl: "/1127722138414092431/h13r0Wy77BUwmrvZtWdtvTRlmMjVetXDMVI1VDaJF13zOg6iZUh888FGk7vrR2fwTOa-",
			Content:    fmt.Sprintf("(%v) Loop estretégias executadas em : %v \n", time.Now().Format("2006-01-02 15:04:05"), time.Since(start)),
		})
	}
}
