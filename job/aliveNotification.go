package job

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

type Response []struct {
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	Turnover  string `json:"turnover"`
	Confirm   bool   `json:"confirm"`
	Timestamp int64  `json:"timestamp"`
}

func (j *job) AliveNotification(loopChannel *chan bool) {
	service := discordService.NewDiscordWebhook()

	var notify discordstructs.Notification

	lastBtcJson, err := j.Redis.Get(context.Background(), "kline.1.BTCUSDT").Result()
	if err != nil {
		log.Panicln("JAN 00: ", err)
	}
	lastEthJson, err := j.Redis.Get(context.Background(), "kline.1.ETHUSDT").Result()
	if err != nil {
		log.Panicln("JAN 01: ", err)
	}

	var responseBtc Response
	var responseEth Response

	err = json.Unmarshal([]byte(lastBtcJson), &responseBtc)
	if err != nil {
		log.Panicln("JAN 02: ", err)
	}

	err = json.Unmarshal([]byte(lastEthJson), &responseEth)
	if err != nil {
		log.Panicln("JAN 03: ", err)
	}

	notify.ChannelUrl = "/1127722818566623243/gQakWasCCQvo8lpsTZqxWFvc6WVWNez3ZwNiqrUr5jqUBguqIk0cDzhe6sUiNMZRrAjj"
	notify.Content = fmt.Sprintf("%v Greetings! Bot Container is up! (BTCUSDT: %v, ETHUSDT: %v)", time.Now().Format("2006-01-02 15:04:05"), responseBtc[0].Close, responseEth[0].Close)

	service.SendNotification(&notify)
	time.Sleep(time.Minute * 5)
	*loopChannel <- true
}
