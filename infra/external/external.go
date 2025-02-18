package external

import (
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"google.golang.org/grpc"
)

type external struct {
	Conn *grpc.ClientConn
}

type ExternalInterface interface {
	ListUserStrategy() *pb.UserStrategyResponse
	ListTradeConfig() *pb.TradeConfigResponse
	GetWalletWithCoin(*pb.GetWalletWithCoinRequest) *pb.GetWalletWithCoinResponse
	ListWalletWithCoin(*pb.ListWalletWithCoinRequest) *pb.ListWalletWithCoinResponse
	ListCoin(*pb.ListCoinRequest) *pb.ListCoinResponse
	CreateWallet(*pb.CreateWalletRequest) *pb.CreateWalletResponse
	UpdateWallet(*pb.UpdateWalletRequest) *pb.UpdateWalletResponse
}

func NewUserExternal() ExternalInterface {
	conn, err := grpc.Dial("ms-user:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CMU: %v", err)
	}

	return &external{
		Conn: conn,
	}
}

func NewExchangeExternal() ExternalInterface {
	conn, err := grpc.Dial("127.0.0.1:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CME: %v", err)
	}

	return &external{
		Conn: conn,
	}
}
