package jobtest_test

import "testing"

func TestCalcAvgPriceTest(t *testing.T) {
	loopChannel := make(chan bool)
	jobs.CalcAvgPrice(&loopChannel)
}
