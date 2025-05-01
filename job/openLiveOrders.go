package job

import (
	"os"

	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

func (j *job) OpenLiveOrders() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveExec(make(<-chan struct{}))
}
