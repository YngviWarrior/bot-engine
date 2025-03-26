package jobtest_test

import (
	"testing"
	"time"
)

func TestAliveNotification(t *testing.T) {
	loopChan := make(chan bool)
	go jobs.AliveNotification(&loopChan)

	time.Sleep(time.Second * 2)
	t.Context().Done()
}
