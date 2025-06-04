package jobtest_test

import (
	"testing"
	"time"
)

func TestManageOperationCreationByStrategy(t *testing.T) {
	loopChannel := make(chan bool)
	go jobs.ManageOperationCreationByStrategy(&loopChannel)

	time.Sleep(time.Second * 2)
	t.Context().Done()
}
