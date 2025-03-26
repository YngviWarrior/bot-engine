package jobtest_test

import (
	"testing"
	"time"
)

func TestOpenLiveKline(t *testing.T) {
	go jobs.OpenLiveKline()

	time.Sleep(time.Second * 5)
	t.Context().Done()
}
