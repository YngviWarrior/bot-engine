package job

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j job) CalculateProfit(loopChannel *chan bool) {
	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from := time.Date(to.Year(), to.Month(), to.Day()-1, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list24H := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(from.UnixMicro()),
		MtsEnd:   uint64(to.UnixMicro()),
	})

	var profit24H float64 = 0
	for _, operation := range list24H.GetOperations() {
		operationHistoryList := j.ExchangeMS.ListOperationHistoryByOperation(&pb.ListOperationHistoryByOperationRequest{
			Operation: operation.GetOperation(),
		})

		if len(operationHistoryList.GetOperationHistory()) == 2 {
			var operationProfit float64 = 0
			for _, oph := range operationHistoryList.GetOperationHistory() {
				stablePrice, err := strconv.ParseFloat(oph.StablePrice, 64)
				if err != nil {
					log.Panic("CP 01: ", err)
				}

				if oph.TransactionType == 1 { // Buy
					operationProfit -= stablePrice
				} else if oph.TransactionType == 2 { // Sell
					operationProfit += stablePrice
				}
			}
			profit24H += operationProfit
		}
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-7, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list7D := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(from.UnixMicro()),
		MtsEnd:   uint64(to.UnixMicro()),
	})

	var profit7D float64 = 0
	for _, operation := range list7D.GetOperations() {
		operationHistoryList := j.ExchangeMS.ListOperationHistoryByOperation(&pb.ListOperationHistoryByOperationRequest{
			Operation: operation.GetOperation(),
		})

		if len(operationHistoryList.GetOperationHistory()) == 2 {
			var operationProfit float64 = 0
			for _, oph := range operationHistoryList.GetOperationHistory() {
				stablePrice, err := strconv.ParseFloat(oph.StablePrice, 64)
				if err != nil {
					log.Panic("CP 01: ", err)
				}

				if oph.TransactionType == 1 { // Buy
					operationProfit -= stablePrice
				} else if oph.TransactionType == 2 { // Sell
					operationProfit += stablePrice
				}
			}
			profit7D += operationProfit
		}
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-30, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list1M := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(from.UnixMicro()),
		MtsEnd:   uint64(to.UnixMicro()),
	})

	var profit1M float64 = 0
	for _, operation := range list1M.GetOperations() {
		operationHistoryList := j.ExchangeMS.ListOperationHistoryByOperation(&pb.ListOperationHistoryByOperationRequest{
			Operation: operation.GetOperation(),
		})

		if len(operationHistoryList.GetOperationHistory()) == 2 {
			var operationProfit float64 = 0
			for _, oph := range operationHistoryList.GetOperationHistory() {
				stablePrice, err := strconv.ParseFloat(oph.StablePrice, 64)
				if err != nil {
					log.Panic("CP 01: ", err)
				}

				if oph.TransactionType == 1 { // Buy
					operationProfit -= stablePrice
				} else if oph.TransactionType == 2 { // Sell
					operationProfit += stablePrice
				}
			}
			profit1M += operationProfit
		}
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year()-1, to.Month(), to.Day(), to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	listAllTime := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(from.UnixMicro()),
		MtsEnd:   uint64(to.UnixMicro()),
	})

	var profitAllTime float64 = 0
	for _, operation := range listAllTime.GetOperations() {
		operationHistoryList := j.ExchangeMS.ListOperationHistoryByOperation(&pb.ListOperationHistoryByOperationRequest{
			Operation: operation.GetOperation(),
		})

		if len(operationHistoryList.GetOperationHistory()) == 2 {
			var operationProfit float64 = 0
			for _, oph := range operationHistoryList.GetOperationHistory() {
				stablePrice, err := strconv.ParseFloat(oph.StablePrice, 64)
				if err != nil {
					log.Panic("CP 01: ", err)
				}

				if oph.TransactionType == 1 { // Buy
					operationProfit -= stablePrice
				} else if oph.TransactionType == 2 { // Sell
					operationProfit += stablePrice
				}
			}
			profitAllTime += operationProfit
		}
	}

	discord := discordService.NewDiscordWebhook()
	var notify discordstructs.Notification
	notify.ChannelUrl = "/1127722252213960864/sAGIhO4C83MQp5hUALmFDUU2jYec1geVg1m2_-Bdscn3ABN2mZygp8B-1Ae-QwW8WuK8"
	notify.Content = fmt.Sprintf("(%v) Daily Profit: $ %v <----> Week Profit: $ %v <----> Month Profit: $ %v <----> Year Profit: $ %v ", time.Now().Format("2006-01-02 15:04:05"), profit24H, profit7D, profit1M, profitAllTime)
	discord.SendNotification(&notify)

	*loopChannel <- true
}
