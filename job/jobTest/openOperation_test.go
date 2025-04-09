package jobtest_test

import (
	"testing"
	"time"
)

func TestOpenOperationManager(t *testing.T) {
	loopChan := make(chan bool)
	go jobs.OpenOperationManager(&loopChan)

	time.Sleep(time.Second * 2)
	t.Context().Done()
}
