package botengine

// userTradeConfig "traderbot/core/entities/userTradeConfig"

// func (b *botengine) ByBitAvgPriceWeek(db database.DatabaseInterface, configs *userTradeConfig.UserTradeConfig, wg *sync.WaitGroup) {
// 	loopChannel := make(chan bool)

// 	go b.OperateAvgPriceWeek(db, configs, &loopChannel, wg)

// 	// for <-loopChannel {
// 	// 	go OperateAvgPriceWeek(db, configs, &loopChannel)
// 	// }
// }

// func (b *botengine) OperateAvgPriceWeek(db database.DatabaseInterface, configs *userTradeConfig.UserTradeConfig, loopChannel *chan bool, wg *sync.WaitGroup) {

// var newRegister bool
// var stableCoin uint64
// switch configs.Parity {
// case 1, 4:
// 	stableCoin = 1
// case 3:
// 	stableCoin = 2
// }

// ctx := context.TODO()
// conn := db.CreateConnection()
// tx := db.CreateTransaction(&ctx, conn)

// wc := walletRepo.FindWithCoin(tx, configs.User, configs.Exchange, stableCoin)
// list := operationRepo.List(tx, configs.User, int64(configs.Strategy), int64(configs.StrategyVariant), int64(configs.Parity), int64(configs.Exchange), false, true)
// currentPrice := candleRepo.FindLastPrice(tx, configs.Parity, configs.Exchange)
// avg := avgPriceRepo.FindByParityExchange(tx, configs.Parity, configs.Exchange)
// // log.Println("Open Operations: ", len(list))

// if len(list) < int(configs.OperationQuantity) && wc.Amount > (configs.OperationAmount+configs.WalletValueLimit) {
// 	opId := operationRepo.Create(tx, configs.User, int64(configs.Parity), int64(configs.Exchange), 1, 2, configs.OperationAmount, true)
// 	if opId == 0 {
// 		tx.Rollback()
// 		conn.Close()
// 		log.Panicln("BITB 01: ")
// 	}

// 	n.ChannelUrl = constants.DISCORD_NEW_OPERATION_TEST
// 	n.Content = fmt.Sprintf("(%v) <---> Registering New Order! %v Strategy %v (%v)", time.Now().Format("2006-01-02 15:04:05"), configs.ParitySymbol, configs.StrategyName, configs.StrategyVariantName)
// 	b.Discord.SendNotification(&n)
// 	newRegister = true
// }

// tx.Commit()
// conn.Close()

// for _, operation := range list {
// 	conn := db.CreateConnection()
// 	tx := db.CreateTransaction(&ctx, conn)

// 	coinQuantity, opFee := operationHistoryRepo.GetLastBuyRegisterByOperation(tx, operation.Operation)
// 	sellCotation := utils.FindSellCotation(operation.InvestedAmount, coinQuantity, opFee, operation.OpenPrice, configs.ProfitPercentage)

// 	var isProfitable bool = (currentPrice >= sellCotation)

// 	// fmt.Printf("%v AVG Day Op: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.Operation, sellCotation)
// 	if wc.Amount > (configs.OperationAmount+configs.WalletValueLimit) && coinQuantity == 0 {
// 		if currentPrice > avg.Week {
// 			tx.Rollback()
// 			conn.Close()
// 			continue
// 		}

// 		var bc BuyCoinParams
// 		bc.Operation = operation.Operation
// 		bc.ClosePrice = currentPrice
// 		bc.Exchange = int64(configs.Exchange)
// 		bc.Symbol = configs.ParitySymbol
// 		bc.OrderType = "LIMIT"
// 		bc.OpFee = 0.001

// 		switch operation.Parity {
// 		case constants.BTCUSDT:
// 			bc.OpAmount = fmt.Sprintf("%.6f", utils.Truncate(configs.OperationAmount/currentPrice, 5))
// 		case constants.ETHUSDT:
// 			bc.OpAmount = fmt.Sprintf("%.5f", utils.Truncate(configs.OperationAmount/currentPrice, 4))
// 		}

// 		switch os.Getenv("ENVIROMENT") {
// 		case "server", "local", "testnet":
// 			if b.BuyCoin(tx, &bc) {
// 				if !walletRepo.UpdateAmount(tx, fmt.Sprintf("(%v, %v, %v, %v, %v)", wc.Wallet, wc.Exchange, configs.User, wc.Coin, (-1*configs.OperationAmount))) {
// 					tx.Rollback()
// 					conn.Close()
// 					log.Panicln("BITB 02: ")
// 				}

// 				var x repository.OperationDto

// 				x.Operation = operation.Operation
// 				x.MtsStart = operation.MtsStart
// 				x.InvestedAmount = configs.OperationAmount
// 				x.OpenPrice = currentPrice
// 				x.ClosePrice = operation.ClosePrice
// 				x.Profit = configs.ProfitPercentage
// 				x.Closed = false
// 				x.Enabled = true

// 				if !operationRepo.Update(tx, &x) {
// 					tx.Rollback()
// 					conn.Close()
// 					log.Panicln("BITB 03: ")
// 				}

// 				newRegister = true
// 			}
// 		}
// 	} else if isProfitable && coinQuantity != 0 {
// 		var p SellCoinParams

// 		p.Operation = operation.Operation
// 		p.ClosePrice = currentPrice

// 		p.Exchange = int64(configs.Exchange)
// 		p.OpAmount = operation.InvestedAmount
// 		p.Symbol = configs.ParitySymbol
// 		p.OrderType = "LIMIT"
// 		p.OpFee = 0.001

// 		switch operation.Parity {
// 		case constants.BTCUSDT:
// 			p.CoinQuantity = utils.Truncate(coinQuantity, 5)
// 		case constants.ETHUSDT:
// 			p.CoinQuantity = utils.Truncate(coinQuantity, 4)
// 		}

// 		switch os.Getenv("ENVIROMENT") {
// 		case "server", "local", "testnet":
// 			if b.SellCoin(tx, &p) {
// 				if !walletRepo.UpdateAmount(tx, fmt.Sprintf("(%v, %v, %v, %v, %v)", wc.Wallet, wc.Exchange, configs.User, wc.Coin, (operation.InvestedAmount*(1+configs.ProfitPercentage)))) {
// 					tx.Rollback()
// 					conn.Close()
// 					log.Panicln("BITB 04: ")
// 				}

// 				var x repository.OperationDto

// 				x.Operation = operation.Operation
// 				x.MtsStart = operation.MtsStart
// 				x.MtsFinish = time.Now().UnixMicro()
// 				x.InvestedAmount = operation.InvestedAmount
// 				x.OpenPrice = operation.OpenPrice
// 				x.ClosePrice = currentPrice
// 				x.ProfitAmount = (currentPrice * coinQuantity) - operation.InvestedAmount
// 				x.Profit = operation.Profit
// 				x.Closed = false
// 				x.Enabled = true

// 				if !operationRepo.Update(tx, &x) {
// 					tx.Rollback()
// 					conn.Close()
// 					log.Panicln("BITB 05: ")
// 				}

// 				newRegister = true
// 			}
// 		}
// 	}

// 	tx.Commit()
// 	conn.Close()
// }

// if newRegister {
// 	time.Sleep(time.Millisecond * 500)
// }

// wg.Done()
// *loopChannel <- true
// }
