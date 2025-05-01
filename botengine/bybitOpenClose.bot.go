package botengine

import (
	"sync"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
)

func (b *botengine) ByBitOpenClose(configs *pb.TradeConfig, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	b.OperateOpenClose(configs, kline, wg)
}

func (b *botengine) OperateOpenClose(tradeConfig *pb.TradeConfig, kline *rabbitmq.CombinedData, wg *sync.WaitGroup) {
	b.External.Strategy(&pb.StrategyRequest{
		Kline: &pb.KlineData{
			Start:     kline.Data[0].Start,
			End:       kline.Data[0].End,
			Interval:  kline.Data[0].Interval,
			Open:      kline.Data[0].Open,
			Close:     kline.Data[0].Close,
			High:      kline.Data[0].High,
			Low:       kline.Data[0].Low,
			Volume:    kline.Data[0].Volume,
			Turnover:  kline.Data[0].Turnover,
			Confirm:   kline.Data[0].Confirm,
			Timestamp: kline.Data[0].Timestamp,
		},
		TradeConfig: tradeConfig,
		Strategy:    "OpenClose",
	})
}
