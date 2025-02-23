package job

import (
	"os"

	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

var topics = []string{"kline.1.BTCUSDT", "kline.1.ETHUSDT"}

func (j *job) OpenLiveKline() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LivePublic(topics, make(<-chan struct{}))
}
