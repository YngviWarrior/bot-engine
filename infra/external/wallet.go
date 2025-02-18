package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) ListWalletWithCoin(in *pb.ListWalletWithCoinRequest) (out *pb.ListWalletWithCoinResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListWalletWithCoin(context.Background(), &pb.ListWalletWithCoinRequest{
		Exchange: in.GetExchange(),
	})

	if err != nil {
		log.Fatalln("E-LWEC", err)
	}

	return
}

func (e *external) GetWalletWithCoin(in *pb.GetWalletWithCoinRequest) (out *pb.GetWalletWithCoinResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetWalletWithCoin(context.Background(), &pb.GetWalletWithCoinRequest{
		Exchange: in.GetExchange(),
		Coin:     in.GetCoin(),
	})

	if err != nil {
		log.Fatalln("E-GWEC", err)
	}

	return
}

func (e *external) CreateWallet(in *pb.CreateWalletRequest) (out *pb.CreateWalletResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.CreateWallet(context.Background(), &pb.CreateWalletRequest{
		Coin:     in.GetCoin(),
		Exchange: in.GetExchange(),
		Amount:   in.GetAmount(),
	})

	if err != nil {
		log.Fatalln("E-CW", err)
	}

	return
}

func (e *external) UpdateWallet(in *pb.UpdateWalletRequest) (out *pb.UpdateWalletResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.UpdateWallet(context.Background(), &pb.UpdateWalletRequest{
		Wallet: in.GetWallet(),
	})

	if err != nil {
		log.Fatalln("E-UW", err)
	}

	return
}
