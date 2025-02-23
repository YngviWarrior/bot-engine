package job

import (
	"fmt"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

func (j job) CalculateProfit() {
	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from := time.Date(to.Year(), to.Month(), to.Day()-1, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list24H := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(to.UnixMicro()),
		MtsEnd:   uint64(from.UnixMicro()),
	})

	var profit24H float64 = 0
	for _, op := range list24H.GetOperations() {
		profit24H += op.ProfitAmount
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-7, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list7D := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(to.UnixMicro()),
		MtsEnd:   uint64(from.UnixMicro()),
	})

	var profit7D float64 = 0
	for _, op := range list7D.GetOperations() {
		profit7D += op.ProfitAmount
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year(), to.Month(), to.Day()-30, to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	list1M := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(to.UnixMicro()),
		MtsEnd:   uint64(from.UnixMicro()),
	})

	var profit1M float64 = 0
	for _, op := range list1M.GetOperations() {
		profit1M += op.ProfitAmount
	}

	to = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	from = time.Date(to.Year()-1, to.Month(), to.Day(), to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.Local)
	listAllTime := j.ExchangeMS.ListOperationByPeriod(&pb.ListOperationByPeriodRequest{
		MtsStart: uint64(to.UnixMicro()),
		MtsEnd:   uint64(from.UnixMicro()),
	})

	var profitAllTime float64 = 0
	for _, op := range listAllTime.GetOperations() {
		profitAllTime += op.ProfitAmount
	}

	discord := discordService.NewDiscordWebhook()
	var notify discordstructs.Notification
	notify.ChannelUrl = "/1127722252213960864/sAGIhO4C83MQp5hUALmFDUU2jYec1geVg1m2_-Bdscn3ABN2mZygp8B-1Ae-QwW8WuK8"
	notify.Content = fmt.Sprintf("(%v) Daily Profit: $ %v <----> Week Profit: $ %v <----> Month Profit: $ %v <----> Year Profit: $ %v ", time.Now().Format("2006-01-02 15:04:05"), profit24H, profit7D, profit1M, profitAllTime)
	discord.SendNotification(&notify)
}
