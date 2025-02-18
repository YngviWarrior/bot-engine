package botengine

// func (b *botengine) ByBitFastTrade(db database.DatabaseInterface, configs *userTradeConfig.UserTradeConfig, wg *sync.WaitGroup) {
// 	// loopChannel := make(chan bool)

// 	// go OperateFastTrade(db, configs, &loopChannel, wg)

// 	// for <-loopChannel {
// 	// 	go OperateFastTrade(db, configs, &loopChannel)
// 	// }
// }

// func (b *botengine) OperateFastTrade(db database.DatabaseInterface, configs *userTradeConfig.UserTradeConfig, loopChannel *chan bool, wg *sync.WaitGroup) {
// 	var newRegister bool
// 	var stableCoin uint64
// 	switch configs.Parity {
// 	case 1, 4:
// 		stableCoin = 1
// 	case 3:
// 		stableCoin = 2
// 	}

// 	ctx := context.TODO()
// 	conn := db.CreateConnection()
// 	tx := db.CreateTransaction(&ctx, conn)

// 	wc := walletRepo.FindWithCoin(tx, configs.User, configs.Exchange, stableCoin)
// 	list := operationRepo.List(tx, configs.User, int64(configs.Strategy), int64(configs.StrategyVariant), int64(configs.Parity), int64(configs.Exchange), false, true)
// 	currentPrice := candleRepo.FindLastPrice(tx, configs.Parity, configs.Exchange)

// 	// fmt.Println("Open Operations: ", len(list))
// 	if len(list) < int(configs.OperationQuantity) && wc.Amount > (configs.OperationAmount+configs.WalletValueLimit) {
// 		opId := operationRepo.Create(tx, configs.User, int64(configs.Parity), int64(configs.Exchange), 3, 1, configs.OperationAmount, true)
// 		if opId == 0 {
// 			tx.Rollback()
// 			conn.Close()
// 			log.Panicln("BITB 01: ")
// 		}

// 		if !operationMetaDataFastTradeRepo.Create(tx, opId, currentPrice) {
// 			tx.Rollback()
// 			conn.Close()
// 			log.Panicln("BITB 01: ")
// 		}

// 		n.ChannelUrl = constants.DISCORD_NEW_OPERATION_TEST
// 		n.Content = fmt.Sprintf("(%v) <---> Registering New Order! %v Strategy %v (%v)", time.Now().Format("2006-01-02 15:04:05"), configs.ParitySymbol, configs.StrategyName, configs.StrategyVariantName)
// 		b.Discord.SendNotification(&n)
// 		newRegister = true
// 	}

// 	tx.Commit()
// 	conn.Close()

// 	for _, operation := range list {
// 		// mtsS := time.UnixMicro(operation.MtsStart)
// 		// mtsStart := time.Date(mtsS.Year(), mtsS.Month(), mtsS.Day(), mtsS.Hour(), mtsS.Minute(), mtsS.Second()+1, mtsS.Nanosecond(), time.Local)

// 		var buy repository.OperationDto
// 		var sell repository.OperationDto

// 		conn := db.CreateConnection()
// 		tx := db.CreateTransaction(&ctx, conn)

// 		var isReciding bool = false
// 		updateFields := []string{"is_receding", "last_price"}
// 		updateValues := []any{0, currentPrice}
// 		if currentPrice > operation.LastPrice.Float64 {
// 			if currentPrice > operation.MaximumPrice.Float64 {
// 				updateFields = append(updateFields, "maximum_price")
// 				updateValues = append(updateValues, currentPrice)
// 				operation.MaximumPrice.Float64 = currentPrice
// 			}

// 			isReciding = false
// 			operationMetaDataFastTradeRepo.UpdateDynamically(tx, updateFields, updateValues, []string{"operation_meta_data_fast_trade"}, []any{operation.OperationMetaDataFastTrade}, []any{}, "")
// 			operation.LastPrice.Float64 = currentPrice
// 		} else if currentPrice < operation.LastPrice.Float64 {
// 			if currentPrice < operation.MinimumPrice.Float64 {
// 				updateFields = append(updateFields, "minimum_price")
// 				updateValues = append(updateValues, currentPrice)
// 				operation.MinimumPrice.Float64 = currentPrice
// 			}

// 			isReciding = true
// 			operationMetaDataFastTradeRepo.UpdateDynamically(tx, updateFields, updateValues, []string{"operation_meta_data_fast_trade"}, []any{operation.OperationMetaDataFastTrade}, []any{}, "")
// 			operation.LastPrice.Float64 = currentPrice
// 		}

// 		coinQuantity, opFee := operationHistoryRepo.GetLastBuyRegisterByOperation(tx, operation.Operation)
// 		sellCotation := utils.FindSellCotation(operation.InvestedAmount, coinQuantity, opFee, operation.OpenPrice, configs.ProfitPercentage)

// 		var movedAwayFromTheMaximum bool = (isReciding && (currentPrice <= operation.MaximumPrice.Float64*0.99995) && (operation.MaximumPrice.Float64*0.99995 < operation.OpenPrice))
// 		var lowerThanOpening bool = (isReciding && currentPrice <= operation.OpenPrice)
// 		var isProfitable bool = (currentPrice >= sellCotation)

// 		// fmt.Printf("%v OP: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.Operation, sellCotation)
// 		if wc.Amount > (configs.OperationAmount+configs.WalletValueLimit) && coinQuantity == 0 {
// 			var bc BuyCoinParams

// 			bc.Operation = operation.Operation
// 			bc.ClosePrice = currentPrice
// 			bc.Exchange = int64(configs.Exchange)
// 			bc.Symbol = configs.ParitySymbol
// 			bc.OrderType = "MARKET"
// 			bc.OpFee = 0.001

// 			switch operation.Parity {
// 			case constants.BTCUSDT:
// 				bc.OpAmount = fmt.Sprintf("%.6f", configs.OperationAmount)
// 			case constants.ETHUSDT:
// 				bc.OpAmount = fmt.Sprintf("%.6f", configs.OperationAmount)
// 			}

// 			coinQty := (configs.OperationAmount / currentPrice) * (1 - bc.OpFee)

// 			switch os.Getenv("ENVIROMENT") {
// 			case "server", "local", "testnet":
// 				if b.BuyCoin(tx, &bc) {
// 					if !walletRepo.UpdateAmount(tx, fmt.Sprintf("(%v, %v, %v, %v, %v)", wc.Wallet, wc.Exchange, configs.User, wc.Coin, (configs.OperationAmount*-1))) {
// 						tx.Rollback()
// 						conn.Close()
// 						log.Panicln("BITB 02: ")
// 					}

// 					buy.Operation = operation.Operation
// 					buy.MtsStart = time.Now().UnixMicro()
// 					buy.InvestedAmount = configs.OperationAmount
// 					buy.OpenPrice = currentPrice
// 					buy.ClosePrice = operation.ClosePrice
// 					buy.ProfitAmount = 0
// 					buy.Profit = 0
// 					buy.Closed = false
// 					buy.Enabled = true

// 					if !operationRepo.Update(tx, &buy) {
// 						tx.Rollback()
// 						conn.Close()
// 						log.Panicln("BITB 03: ")
// 					}

// 					n.ChannelUrl = constants.DISCORD_OPEN_OPERATION_TEST
// 					n.Content = fmt.Sprintf("(%v) <---> Bought %.7f %v at %v cotation !", time.Now().Format("2006-01-02 15:04:05"), coinQty, configs.ParitySymbol, currentPrice)
// 					b.Discord.SendNotification(&n)

// 					newRegister = true
// 				}
// 			}
// 		} else if (isProfitable || lowerThanOpening || movedAwayFromTheMaximum) && coinQuantity != 0 {
// 			var p SellCoinParams

// 			p.Operation = operation.Operation
// 			p.ClosePrice = currentPrice
// 			p.Exchange = int64(configs.Exchange)
// 			p.OpAmount = operation.InvestedAmount
// 			p.Symbol = configs.ParitySymbol
// 			p.OrderType = "MARKET"
// 			p.OpFee = 0.001

// 			switch operation.Parity {
// 			case constants.BTCUSDT:
// 				p.CoinQuantity = utils.Truncate(coinQuantity, 5)
// 			case constants.ETHUSDT:
// 				p.CoinQuantity = utils.Truncate(coinQuantity, 4)
// 			}

// 			switch os.Getenv("ENVIROMENT") {
// 			case "server", "local", "testnet":
// 				if b.SellCoin(tx, &p) {
// 					if lowerThanOpening || movedAwayFromTheMaximum {
// 						fmt.Println(configs.ParitySymbol, configs.StrategyName, "Deactivate")
// 						if !userTraderConfigRepo.UpdateDynamically(tx, []string{"enabled"}, []any{0}, []string{"strategy", "parity"}, []any{configs.Strategy, configs.Parity}, []any{}, "") {
// 							tx.Rollback()
// 							conn.Close()
// 							log.Panicln("BITB 04: ")
// 						}

// 						n.ChannelUrl = constants.DISCORD_DEBUG_TEST
// 						n.Content = fmt.Sprintf("(%v) Strategy %v(%v) is Desabled.", time.Now().Format("2006-01-02 15:04:05"), configs.StrategyName, configs.StrategyVariantName)
// 						b.Discord.SendNotification(&n)
// 					}

// 					if !walletRepo.UpdateAmount(tx, fmt.Sprintf("(%v, %v, %v, %v, %v)", wc.Wallet, wc.Exchange, configs.User, wc.Coin, (p.CoinQuantity*currentPrice))) {
// 						tx.Rollback()
// 						conn.Close()
// 						log.Panicln("BITB 05: ")
// 					}

// 					sell.Operation = operation.Operation
// 					sell.MtsStart = operation.MtsStart
// 					sell.MtsFinish = time.Now().UnixMicro()
// 					sell.InvestedAmount = operation.InvestedAmount
// 					sell.OpenPrice = operation.OpenPrice
// 					sell.ClosePrice = currentPrice
// 					sell.ProfitAmount = 0
// 					sell.Profit = operation.Profit
// 					sell.Closed = true
// 					sell.Enabled = false

// 					if !operationRepo.Update(tx, &sell) {
// 						tx.Rollback()
// 						conn.Close()
// 						log.Panicln("BITB 06: ")
// 					}

// 					n.ChannelUrl = constants.DISCORD_OPEN_OPERATION_TEST
// 					n.Content = fmt.Sprintf("(%v) <---> Sold %.7f %v at %v cotation !", time.Now().Format("2006-01-02 15:04:05"), coinQuantity, configs.ParitySymbol, currentPrice)
// 					b.Discord.SendNotification(&n)

// 					newRegister = true
// 				}
// 			}
// 		}

// 		tx.Commit()
// 		conn.Close()
// 	}

// 	if newRegister {
// 		time.Sleep(time.Millisecond * 500)
// 	}

// 	wg.Done()
// 	*loopChannel <- true
// }
