package botengine

import (
	"sync"
	"time"

	external "github.com/YngviWarrior/bot-engine/infra/external"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

type BotEngine interface {
	InitBotEngine()
}

type botengine struct {
	Bybit bybitSDK.BybitServiceInterface
}

func NewBotEngine(bybit bybitSDK.BybitServiceInterface) BotEngine {
	return &botengine{
		Bybit: bybit,
	}
}

func (b *botengine) InitBotEngine() {
	externalRequest := external.NewExchangeExternal()
	tradeConfigList := externalRequest.ListTradeConfig()

	var wg sync.WaitGroup

	var start time.Time
	start = time.Now()

	for _, configs := range tradeConfigList.GetTradeConfig() {
		if configs.Enabled {
			wg.Add(1)
			switch configs.Exchange {
			case BINANCE_EXCH:
				if configs.Enabled {
					switch configs.Modality {
					case SPOT_MODALITY:
						// go binancestrategies.BinanceOpenClose(*db, configs)
					}
				}
			case BYBIT_EXCH:
				if configs.Enabled {
					switch configs.Modality {
					case SPOT_MODALITY:
						switch configs.Strategy {
						case AVERAGE_PRICE:
							switch configs.StrategyVariant {
							case AVERAGE_PRICE_DAY:
								if configs.StrategyVariantEnabled {
									go b.ByBitAvgPriceDay(configs, &wg)
								}
							case AVERAGE_PRICE_WEEK:
								if configs.StrategyVariantEnabled {
									// go b.ByBitAvgPriceWeek(configs, &wg)
								}
							}
						case CLOSE_OPEN:
							if configs.StrategyEnabled {
								// go b.ByBitOpenClose(configs, &wg)
							}
						case FAST_TRADE:
							if configs.StrategyEnabled {
								// go b.ByBitFastTrade(configs, &wg)
							}
						}
					}
				}
			}
		}
	}

	wg.Wait()
	if time.Since(start) > (time.Second * 1) {
		// n.ChannelUrl = DISCORD_DEBUG_TEST
		// n.Content = fmt.Sprintf("(%v) Loop estretégias executadas em : %v \n", time.Now().Format("2006-01-02 15:04:05"), time.Since(start))
		// b.Discord.SendNotification(&n)
	}

	// *loopChannel <- true
}
