package job

import (
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/bot-engine/botengine"
	"github.com/YngviWarrior/bot-engine/infra/external"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	"github.com/rabbitmq/amqp091-go"
)

type job struct {
	RabbitMQ   *amqp091.Connection
	ExchangeMS external.ExternalInterface
}

type JobInterface interface {
	InitJobs()
	OpenLiveKline()
	AliveNotification(*chan bool)
	SyncAssets(*chan bool)
	SyncKlines()
	ManageTradeConfigStrategy(*chan bool)
}

func NewJobs() JobInterface {
	exchangeMS := external.NewExchangeExternal()

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	defer conn.Close()

	return &job{
		RabbitMQ:   conn,
		ExchangeMS: exchangeMS,
	}
}

func (j *job) InitJobs() {
	go j.OpenLiveKline()
	time.Sleep(time.Second * 5)

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
