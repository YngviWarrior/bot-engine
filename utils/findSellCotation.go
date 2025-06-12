package utils

import (
	"fmt"
	"log"
	"strconv"
)

func FindSellCotation(investedAmount, coinQuantity, fee, openPrice, expectedProfit string) (closePrice float64) {
	if investedAmount == "" || coinQuantity == "" || fee == "" || openPrice == "" || expectedProfit == "" {
		closePrice = 0
		return
	}

	investedAmountFloat, err := strconv.ParseFloat(investedAmount, 64)
	if err != nil {
		log.Panic("FSC 01: ", err)
	}
	fmt.Println(coinQuantity)
	coinQuantityFloat, err := strconv.ParseFloat(coinQuantity, 64)
	if err != nil {
		log.Panic("FSC 02: ", err)
	}
	feeFloat, err := strconv.ParseFloat(fee, 64)
	if err != nil {
		log.Panic("FSC 03: ", err)
	}
	// openPriceFloat, err := strconv.ParseFloat(openPrice, 64)
	// if err != nil {
	// 	log.Panic("FSC 04: ", err)
	// }
	expectedProfitFloat, err := strconv.ParseFloat(expectedProfit, 64)
	if err != nil {
		log.Panic("FSC 05: ", err)
	}

	var buyfeeAmount, sellFeeAmount, amountToClose, wantedProfit float64
	if feeFloat > 0 {
		buyfeeAmount = investedAmountFloat * feeFloat // taxa q foi subtraida na compra
		wantedProfit = ((investedAmountFloat + buyfeeAmount) * (1 + expectedProfitFloat))
		sellFeeAmount = wantedProfit * feeFloat        // taxa no momento da venda
		amountToClose = (wantedProfit + sellFeeAmount) // acrescentar taxa que virá
	} else {
		wantedProfit = (investedAmountFloat * (1 + expectedProfitFloat)) //profit desejado
		amountToClose = wantedProfit
	}

	if coinQuantityFloat > 0 {
		closePrice = amountToClose / coinQuantityFloat
	} else {
		closePrice = 0
	}

	return
}
