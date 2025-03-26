package job

import (
	"os"
	"time"

	"github.com/YngviWarrior/bot-engine/botengine"
	"github.com/YngviWarrior/bot-engine/infra/external"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

type job struct {
	RabbitMQ   rabbitmq.RabbitMQInterface
	ExchangeMS external.ExternalInterface
}

type JobInterface interface {
	InitJobs()
	OpenLiveKline()
	AliveNotification(*chan bool)
	SyncAssets(*chan bool)
	SyncKlines()
	ManageTradeConfigStrategy(*chan bool)
	CalculateProfit()
}

func NewJobs() JobInterface {
	exchangeMS := external.NewExchangeExternal()
	rabbitmq := rabbitmq.NewRabbitMQConnection()

	return &job{
		RabbitMQ:   rabbitmq,
		ExchangeMS: exchangeMS,
	}
}

func (j *job) InitJobs() {
	go j.OpenLiveKline()
	time.Sleep(time.Second * 5)

	go func() {
		loopChannel := make(chan bool)
		j.OpenOperationManager(&loopChannel)

		for <-loopChannel {
			j.OpenOperationManager(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.ManageTradeConfigStrategy(&loopChannel)

		for <-loopChannel {
			j.ManageTradeConfigStrategy(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.CalculateProfit()

		for <-loopChannel {
			time.Sleep(time.Minute * 2)
			j.CalculateProfit()
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.ManageTradeConfigStrategy(&loopChannel)

		for <-loopChannel {
			j.ManageTradeConfigStrategy(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.SyncKlines()

		for <-loopChannel {
			time.Sleep(time.Minute * 2)
			j.SyncKlines()
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.AliveNotification(&loopChannel)

		for <-loopChannel {
			j.AliveNotification(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		j.SyncAssets(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 30)
			j.SyncAssets(&loopChannel)
		}
	}()

	func() {
		loopChannel := make(chan bool)
		bot := botengine.NewBotEngine(bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY")))
		bot.InitBotEngine(&loopChannel)

		for <-loopChannel {
			go bot.InitBotEngine(&loopChannel)
		}
	}()

}
