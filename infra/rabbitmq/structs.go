package rabbitmq

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
