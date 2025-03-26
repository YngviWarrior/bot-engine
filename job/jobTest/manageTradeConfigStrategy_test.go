package jobtest_test

import (
	"testing"
	"time"
)

func TestManageTradeConfigStrategy(t *testing.T) {
	loopChannel := make(chan bool)
	go jobs.ManageTradeConfigStrategy(&loopChannel)

	time.Sleep(time.Second * 2)
	t.Context().Done()
}
