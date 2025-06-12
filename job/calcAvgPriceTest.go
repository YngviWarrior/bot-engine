package job

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) CalcAvgPrice(loopChannel *chan bool) {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	parities := j.ExchangeMS.ListParity(&pb.ListParityRequest{})

	for _, parity := range parities.Parities {
		if !parity.Active {
			continue
		}

		list := bybit.GetKlines(&bybitstructs.GetKlinesParams{
			Symbol:   parity.GetSymbol(),
			Interval: "D",
			Limit:    2,
		}).Result.List

		var dayRoc float64
		currentPrice, _ := strconv.ParseFloat(list[0][2], 64)
		lastPeriodPrice, _ := strconv.ParseFloat(list[1][3], 64)

		dayRoc = ((currentPrice - lastPeriodPrice) / lastPeriodPrice) * 100
		dayAvg := (currentPrice + lastPeriodPrice) / 2

		list = bybit.GetKlines(&bybitstructs.GetKlinesParams{
			Symbol:   parity.GetSymbol(),
			Interval: "W",
			Limit:    2,
		}).Result.List

		var weekRoc float64
		currentPrice, _ = strconv.ParseFloat(list[0][2], 64)
		lastPeriodPrice, _ = strconv.ParseFloat(list[1][3], 64)

		weekRoc = ((currentPrice - lastPeriodPrice) / lastPeriodPrice) * 100
		weekAvg := (currentPrice + lastPeriodPrice) / 2

		list = bybit.GetKlines(&bybitstructs.GetKlinesParams{
			Symbol:   parity.GetSymbol(),
			Interval: "M",
			Limit:    2,
		}).Result.List

		var monthRoc float64
		currentPrice, _ = strconv.ParseFloat(list[0][2], 64)
		lastPeriodPrice, _ = strconv.ParseFloat(list[1][3], 64)

		monthRoc = ((currentPrice - lastPeriodPrice) / lastPeriodPrice) * 100
		monthAvg := (currentPrice + lastPeriodPrice) / 2

		j.ExchangeMS.UpdateAveragePrice(&pb.UpdateAveragePriceRequest{
			Parity:               parity.GetParity(),
			Exchange:             2,
			Day:                  fmt.Sprintf("%0.2f", dayAvg),
			DayRoc:               fmt.Sprintf("%0.2f", dayRoc),
			DayUpdateTimestamp:   uint64(time.Now().UnixMicro()),
			Week:                 fmt.Sprintf("%0.2f", weekAvg),
			WeekRoc:              fmt.Sprintf("%0.2f", weekRoc),
			WeekUpdateTimestamp:  uint64(time.Now().UnixMicro()),
			Month:                fmt.Sprintf("%0.2f", monthAvg),
			MonthRoc:             fmt.Sprintf("%0.2f", monthRoc),
			MonthUpdateTimestamp: uint64(time.Now().UnixMicro()),
		})
	}

	*loopChannel <- true
}
