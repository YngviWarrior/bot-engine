package job

import (
	"os"
	"strconv"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
)

func (j *job) SyncKlines() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	parities := j.ExchangeMS.ListParity(&pb.ListParityRequest{})

	for _, parity := range parities.Parities {
		if !parity.Active {
			continue
		}

		fMts := j.ExchangeMS.GetCandleFirstMts(&pb.GetCandleFirstMtsRequest{
			Parity:   parity.Parity,
			Exchange: 2,
		})

		var list bybitstructs.GetKlinesResponse
		if len(fMts.Candles) == 0 {
			list = bybit.GetKlines(&bybitstructs.GetKlinesParams{
				Symbol:    parity.Symbol,
				Interval:  "1",
				StartTime: int64(fMts.Candles[0].Mts) / 1000,
				EndTime:   time.Now().UnixMilli(),
				Limit:     200,
			})
		} else {
			list = bybit.GetKlines(&bybitstructs.GetKlinesParams{
				Symbol:    parity.Symbol,
				Interval:  "1",
				StartTime: int64(fMts.Candles[0].Mts) / 1000,
				EndTime:   time.Now().UnixMilli(),
			})
		}

		create := pb.CreateCandlesRequest{
			Candles: []*pb.Candle{},
		}

		for _, kline := range list.Result.List {
			var c pb.Candle

			c.Parity = parity.Parity
			c.Exchange = 2

			mts, _ := strconv.Atoi(kline[0])
			c.Mts = uint64(mts)

			c.Open = kline[1]
			c.Close = kline[4]
			c.High = kline[2]
			c.Low = kline[3]
			c.Volume = kline[5]

			create.Candles = append(create.Candles, &c)
		}

		j.ExchangeMS.CreateCandles(&create)
	}

}
