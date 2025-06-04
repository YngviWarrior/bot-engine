package job

import (
	"os"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

var topics = []string{}

func (j *job) OpenLiveKline() {
	parities := j.ExchangeMS.ListParity(&pb.ListParityRequest{})

	if len(parities.GetParities()) == 0 {
		topics = append(topics, "kline.1.BTCUSDT", "kline.1.ETHUSDT")
	} else {
		for _, parity := range parities.GetParities() {
			topics = append(topics, "kline.1."+parity.GetSymbol())
		}
	}

	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LivePublic(topics, make(<-chan struct{}))
}
