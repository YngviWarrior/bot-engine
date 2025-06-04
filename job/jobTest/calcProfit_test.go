package jobtest_test

import "testing"

func TestCalcProfit(t *testing.T) {
	loopChannel := make(chan bool)
	jobs.CalculateProfit(&loopChannel)
}
