package jobtest_test

import (
	"testing"
	"time"
)

func TestSyncAssets(t *testing.T) {
	loopChannel := make(chan bool)
	go jobs.SyncAssets(&loopChannel)

	time.Sleep(time.Second * 5)
	t.Context().Done()
}
