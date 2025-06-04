package job

import (
	"os"
	"strings"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"github.com/YngviWarrior/bot-engine/utils"
	bybitSDK "github.com/YngviWarrior/bybit-sdk"
)

func (j *job) SyncAssets(loopChannel *chan bool) {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))
	coins := j.ExchangeMS.ListCoin(&pb.ListCoinRequest{})

	if len(coins.Coin) == 0 {
		return
	}

	walletCoinList := j.ExchangeMS.ListWalletWithCoin(&pb.ListWalletWithCoinRequest{
		Exchange: 2,
	})

	if len(walletCoinList.GetWallet()) == 0 {
		for _, coin := range coins.GetCoin() {
			newWallet := j.ExchangeMS.CreateWallet(&pb.CreateWalletRequest{
				Amount:   0,
				Exchange: 2,
				Coin:     coin.GetCoin(),
			})
			if (newWallet == &pb.CreateWalletResponse{}) {
				return
			}
		}

		walletCoinList = j.ExchangeMS.ListWalletWithCoin(&pb.ListWalletWithCoinRequest{
			Exchange: 2,
		})
	}

	bybitWallets := bybit.GetWalletInfo()

	if len(bybitWallets.Account.Balance) == 0 {
		return
	}

	updateValues := []*pb.Wallet{}
	for _, walletCoin := range walletCoinList.Wallet {
		// fmt.Printf("%+v\n", bybitWallets)
		for _, bybitWallet := range bybitWallets.Account.Balance {
			if strings.ToUpper(walletCoin.Symbol) == bybitWallet.Coin {
				updateValues = append(updateValues, &pb.Wallet{
					Wallet:   walletCoin.GetWallet(),
					Exchange: walletCoin.GetExchange(),
					Coin:     walletCoin.GetCoin(),
					Amount:   utils.ParseFloat(bybitWallet.WalletBalance),
				})
			}
		}
	}

	_ = j.ExchangeMS.UpdateWallet(&pb.UpdateWalletRequest{
		Wallet: updateValues,
	})

	*loopChannel <- true
}
