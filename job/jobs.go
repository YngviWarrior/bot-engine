package job

import (
	"fmt"
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
	OpenOperationManager(*chan bool)
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
	j.SyncKlines()

	go j.OpenLiveKline()
	time.Sleep(time.Second * 5)

	go func() {
		loopChannel := make(chan bool)
		go j.CalculateAveragePrices(&loopChannel)

		for <-loopChannel {
			go j.CalculateAveragePrices(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.OpenOperationManager(&loopChannel)

		for <-loopChannel {
			go j.OpenOperationManager(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.ManageTradeConfigStrategy(&loopChannel)

		for <-loopChannel {
			go j.ManageTradeConfigStrategy(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.CalculateProfit()

		for <-loopChannel {
			time.Sleep(time.Minute * 2)
			go j.CalculateProfit()
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.ManageTradeConfigStrategy(&loopChannel)

		for <-loopChannel {
			go j.ManageTradeConfigStrategy(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.AliveNotification(&loopChannel)

		for <-loopChannel {
			go j.AliveNotification(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.SyncAssets(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 30)
			go j.SyncAssets(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.OpenOperationManager(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 30)
			go j.OpenOperationManager(&loopChannel)
		}
	}()

	go func() {
		bot := botengine.NewBotEngine(bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY")))

		for {
			fmt.Println("AQUIi")
			bot.InitBotEngine()
		}
	}()

}
