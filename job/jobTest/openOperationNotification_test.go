package jobtest_test

import (
	"testing"
	"time"
)

func TestOpenOperationNotification(t *testing.T) {
	loopChan := make(chan bool)
	go jobs.OpenOperationNotification(&loopChan)

	time.Sleep(time.Second * 2)
	t.Context().Done()
}
