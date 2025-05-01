package job

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/bot-engine/botengine"
	"github.com/YngviWarrior/bot-engine/infra/external"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

type MarketData struct {
	Start     uint64 `json:"start"`
	End       uint64 `json:"end"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	Turnover  string `json:"turnover"`
	Confirm   bool   `json:"confirm"`
	Timestamp uint64 `json:"timestamp"`
}

type Wallet struct {
	Wallet   int     `json:"wallet"`
	Exchange int     `json:"exchange"`
	Coin     int     `json:"coin"`
	Amount   float64 `json:"amount"`
}

type FullData struct {
	Market MarketData `json:"market"`
	Wallet Wallet     `json:"wallet"`
	Coin   int        `json:"coin"`
}

type CombinedData struct {
	Type  string       `json:"type"`
	Topic string       `json:"topic"`
	Data  []MarketData `json:"data"`
	Ts    int64        `json:"ts"`
}

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
	go j.OpenLiveOrders()
	j.SyncKlines()

	go j.OpenLiveKline()

	go func() {
		loopChannel := make(chan bool)
		go j.CalculateAveragePrices(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 1)
			go j.CalculateAveragePrices(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.OpenOperationManager(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 1)
			go j.OpenOperationManager(&loopChannel)
		}
	}()

	go func() {
		loopChannel := make(chan bool)
		go j.ManageTradeConfigStrategy(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 1)
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
			time.Sleep(time.Second * 1)
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

	func() {
		bot := botengine.NewBotEngine(bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY")), external.NewExchangeExternal())
		rabbit := rabbitmq.NewRabbitMQConnection()

		messages := rabbit.Listen("klines", "")
		for msg := range messages {
			var kline rabbitmq.CombinedData
			err := json.Unmarshal(msg.Body, &kline)
			if err != nil {
				log.Println("Erro ao parsear mensagem:", err)
				continue
			}
			go bot.InitBotEngine(kline)
		}

	}()

}
