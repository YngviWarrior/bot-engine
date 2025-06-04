package job

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/bot-engine/botengine"
	"github.com/YngviWarrior/bot-engine/infra/external"
	"github.com/YngviWarrior/bot-engine/infra/rabbitmq"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
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
	OpenLiveExec()
	OpenLiveOrder()
	AliveNotification(*chan bool)
	SyncAssets(*chan bool)
	SyncKlines()
	ManageOperationCreationByStrategy(*chan bool)
	CalculateProfit(*chan bool)
	OpenOperationNotification(*chan bool)
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
	orderChannel := make(chan *bybitstructs.OrderRequest)

	go j.OpenLiveKline()
	go j.OpenLiveOrder()
	go j.OpenLiveTrade(orderChannel)

	fmt.Println("Opening live streams")
	j.SyncKlines()
	time.Sleep(time.Second * 5)

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				time.Sleep(time.Second * 10)
				go j.CalculateAveragePrices(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				go j.ManageOperationCreationByStrategy(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				time.Sleep(time.Second * 60)
				go j.CalculateProfit(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				time.Sleep(time.Minute * 5)
				go j.AliveNotification(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				time.Sleep(time.Second * 15)
				go j.SyncAssets(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	go func() {
		loopChannel := make(chan bool)
		go func(loopChannel *chan bool) {
			for <-*loopChannel {
				time.Sleep(time.Second * 60)
				go j.OpenOperationNotification(loopChannel)
			}
		}(&loopChannel)

		loopChannel <- true
	}()

	func(orderChannel chan *bybitstructs.OrderRequest) {
		bot := botengine.NewBotEngine(bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY")), external.NewExchangeExternal(), orderChannel)
		rabbit := rabbitmq.NewRabbitMQConnection()

		messages := rabbit.Listen("klines", "")
		for msg := range messages {
			var kline rabbitmq.CombinedData
			err := json.Unmarshal(msg.Body, &kline)
			if err != nil {
				log.Println("Erro ao parsear mensagem:", err)
				continue
			}
			bot.InitBotEngine(kline)
		}

		// msg, ok := <-messages
		// if !ok {
		// 	log.Println("Canal de mensagens fechado")
		// 	return
		// }

		// var kline rabbitmq.CombinedData
		// err := json.Unmarshal(msg.Body, &kline)
		// if err != nil {
		// 	log.Println("Erro ao parsear mensagem:", err)
		// 	return
		// }

		// go bot.InitBotEngine(kline)

	}(orderChannel)
}
