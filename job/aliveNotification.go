package job

import (
	"fmt"
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

	notify.ChannelUrl = "/1127722818566623243/gQakWasCCQvo8lpsTZqxWFvc6WVWNez3ZwNiqrUr5jqUBguqIk0cDzhe6sUiNMZRrAjj"
	notify.Content = fmt.Sprintf("%v Greetings! Bot Container is up!", time.Now().Format("2006-01-02 15:04:05"))

	service.SendNotification(&notify)

	*loopChannel <- true
}
