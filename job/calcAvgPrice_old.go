package job

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func safeSet(m map[uint64]map[uint64]map[string]float64, exchange, parity uint64, key string, value float64) {
	if _, ok := m[exchange]; !ok {
		m[exchange] = make(map[uint64]map[string]float64)
	}
	if _, ok := m[exchange][parity]; !ok {
		m[exchange][parity] = make(map[string]float64)
	}
	m[exchange][parity][key] = value
}

func (j *job) CalculateAveragePrices_OLD(loopChannel *chan bool) {
	updateMapping := make(map[uint64]map[uint64]map[string]float64)

	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from := time.Date(to.Year(), to.Month(), to.Day()-1, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	dayList := j.ExchangeMS.ListAvgPrices(&pb.ListAvgPricesRequest{
		To:   uint64(to.UnixMicro()),
		From: uint64(from.UnixMicro()),
	})

	for _, parityInfo := range dayList.GetCandles() {
		dayAvg := j.ExchangeMS.GetFirstPrice(&pb.GetFirstPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
			From:     uint64(from.UnixMicro()),
		})

		currentPrice := j.ExchangeMS.GetLastPrice(&pb.GetLastPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
		})

		var dayRoc float64
		var dayAvgPrice float64
		if (dayAvg != &pb.GetFirstPriceResponse{}) || (currentPrice.GetPrice() != "0") {
			currentPrice, err := strconv.ParseFloat(currentPrice.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 01: ", err)
			}

			dayAvgPrice, err = strconv.ParseFloat(dayAvg.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 02: ", err)
			}

			dayRoc = (currentPrice/dayAvgPrice - 1) * 100
		}

		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "day_avg", dayAvgPrice)
		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "day_roc", dayRoc)
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-7, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	weekList := j.ExchangeMS.ListAvgPrices(&pb.ListAvgPricesRequest{
		To:   uint64(to.UnixMicro()),
		From: uint64(from.UnixMicro()),
	})

	for _, parityInfo := range weekList.GetCandles() {
		weekAvg := j.ExchangeMS.GetFirstPrice(&pb.GetFirstPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
			From:     uint64(from.UnixMicro()),
		})

		currentPrice := j.ExchangeMS.GetLastPrice(&pb.GetLastPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
		})

		var weekRoc float64
		var weekAvgPrice float64
		if (weekAvg != &pb.GetFirstPriceResponse{}) || (currentPrice.GetPrice() != "0") {
			currentPrice, err := strconv.ParseFloat(currentPrice.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 03: ", err)
			}

			weekAvgPrice, err = strconv.ParseFloat(weekAvg.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 04: ", err)
			}
			weekRoc = (currentPrice/weekAvgPrice - 1) * 100
		}

		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "week_avg", weekAvgPrice)
		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "week_roc", weekRoc)
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-30, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	monthList := j.ExchangeMS.ListAvgPrices(&pb.ListAvgPricesRequest{
		To:   uint64(to.UnixMicro()),
		From: uint64(from.UnixMicro()),
	})

	for _, parityInfo := range monthList.GetCandles() {
		monthAvg := j.ExchangeMS.GetFirstPrice(&pb.GetFirstPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
			From:     uint64(from.UnixMicro()),
		})

		currentPrice := j.ExchangeMS.GetLastPrice(&pb.GetLastPriceRequest{
			Parity:   parityInfo.Parity,
			Exchange: parityInfo.Exchange,
		})

		var monthRoc float64
		var monthAvgPrice float64
		if (monthAvg != &pb.GetFirstPriceResponse{}) || (currentPrice.GetPrice() != "0") {
			currentPrice, err := strconv.ParseFloat(currentPrice.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 03: ", err)
			}

			monthAvgPrice, err = strconv.ParseFloat(monthAvg.GetPrice(), 64)
			if err != nil {
				log.Panic("CAP 04: ", err)
			}
			monthRoc = (currentPrice/monthAvgPrice - 1) * 100
		}

		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "month_avg", monthAvgPrice)
		safeSet(updateMapping, parityInfo.Exchange, parityInfo.Parity, "month_roc", monthRoc)
	}

	for exchange, parities := range updateMapping {
		var dayAverage, dayRoc, weekAverage, weekRoc, monthAverage, monthRoc string
		for parity, values := range parities {
			for key, value := range values {
				switch key {
				case "day_avg":
					dayAverage = fmt.Sprintf("%.2f", value)
				case "week_avg":
					weekAverage = fmt.Sprintf("%.2f", value)
				case "month_avg":
					monthAverage = fmt.Sprintf("%.2f", value)
				case "day_roc":
					dayRoc = fmt.Sprintf("%.2f", value)
				case "week_roc":
					weekRoc = fmt.Sprintf("%.2f", value)
				case "month_roc":
					monthRoc = fmt.Sprintf("%.2f", value)
				}
			}

			j.ExchangeMS.UpdateAveragePrice(&pb.UpdateAveragePriceRequest{
				Parity:               parity,
				Exchange:             exchange,
				Day:                  dayAverage,
				DayRoc:               dayRoc,
				DayUpdateTimestamp:   uint64(time.Now().UnixMicro()),
				Week:                 weekAverage,
				WeekRoc:              weekRoc,
				WeekUpdateTimestamp:  uint64(time.Now().UnixMicro()),
				Month:                monthAverage,
				MonthRoc:             monthRoc,
				MonthUpdateTimestamp: uint64(time.Now().UnixMicro()),
			})
		}

	}

	*loopChannel <- true
}
