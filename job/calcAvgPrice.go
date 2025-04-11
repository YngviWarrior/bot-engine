package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (j *job) CalculateAveragePrices(loopChannel *chan bool) {
	// var roc float64
	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from := time.Date(to.Year(), to.Month(), to.Day()-1, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list := j.ExchangeMS.ListAvgPrices(&pb.ListAvgPricesRequest{
		To:   uint64(to.UnixMicro()),
		From: uint64(from.UnixMicro()),
	})

	fmt.Println(list)
	for _, parityInfo := range list.GetCandles() {
		firstPrice := j.ExchangeMS.GetFirstPrice(&pb.GetFirstPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
			From:     uint64(from.UnixMicro()),
		})
		fmt.Println("FirstPrice: ", parityInfo.GetParity(), firstPrice)
	}

	*loopChannel <- true
}
