package job

import (
	"os"

	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) OpenLiveTrade(orderChannel <-chan *bybitstructs.OrderRequest) {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LiveTrade(orderChannel, make(<-chan struct{}))

	// messages := j.RabbitMQ.Listen("livetrade", "order.create")
	// for order := range messages {
	// 	// fmt.Println(order)
	// 	var msg bybitstructs.OrderResponse
	// 	err := json.Unmarshal(order.Body, &msg)
	// 	if err != nil {
	// 		log.Println("Failed to parse msg", err)
	// 		continue
	// 	}
	// }
}
