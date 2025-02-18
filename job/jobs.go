package job

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/bot-engine/infra/external"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	"github.com/redis/go-redis/v9"
)

type job struct {
	Redis      *redis.Client
	ExchangeMS external.ExternalInterface
}

type JobInterface interface {
	InitJobs()
	AliveNotification(*chan bool)
	SyncAssets(*chan bool)
}

func NewJobs() JobInterface {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Endereço do servidor Redis
		Password: "",               // Senha (se houver)
		DB:       0,                // Número do banco de dados
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}

	exchangeMS := external.NewExchangeExternal()

	return &job{
		Redis:      rdb,
		ExchangeMS: exchangeMS,
	}
}

func (j *job) InitJobs() {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	go bybit.LivePublic("kline.1.BTCUSDT,kline.1.ETHUSDT", make(<-chan struct{}))
	time.Sleep(time.Second * 5)

	go func() {
		loopChannel := make(chan bool)
		go j.AliveNotification(&loopChannel)

		for <-loopChannel {
			go j.AliveNotification(&loopChannel)
		}
	}()

	func() {
		loopChannel := make(chan bool)
		go j.SyncAssets(&loopChannel)

		for <-loopChannel {
			time.Sleep(time.Second * 30)
			go j.SyncAssets(&loopChannel)
		}
	}()

}
