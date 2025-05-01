package external

import (
	"log"
	"os"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
	"google.golang.org/grpc"
)

type external struct {
	Conn *grpc.ClientConn
}

type ExternalInterface interface {
	ListUserStrategy() *pb.UserStrategyResponse
	ListTradeConfig() *pb.TradeConfigResponse
	UpdateTradeConfig(*pb.UpdateTradeConfigRequest) *pb.UpdateTradeConfigResponse
	GetTradeConfig(*pb.GetTradeConfigRequest) *pb.GetTradeConfigResponse

	GetWalletWithCoin(*pb.GetWalletWithCoinRequest) *pb.GetWalletWithCoinResponse
	ListWalletWithCoin(*pb.ListWalletWithCoinRequest) *pb.ListWalletWithCoinResponse

	ListParity(*pb.ListParityRequest) *pb.ListParityResponse
	GetAvgPriceByParityExchange(*pb.GetAvgPriceByParityExchangeRequest) *pb.GetAvgPriceByParityExchangeResponse
	ListCoin(*pb.ListCoinRequest) *pb.ListCoinResponse

	CreateWallet(*pb.CreateWalletRequest) *pb.CreateWalletResponse
	UpdateWallet(*pb.UpdateWalletRequest) *pb.UpdateWalletResponse

	GetCandleFirstMts(*pb.GetCandleFirstMtsRequest) *pb.GetCandleFirstMtsResponse
	CreateCandles(*pb.CreateCandlesRequest) *pb.CreateCandlesResponse
	ListCandleLimit(*pb.ListCandleLimitRequest) *pb.ListCandleLimitResponse
	GetFirstPrice(*pb.GetFirstPriceRequest) *pb.GetFirstPriceResponse
	GetLastPrice(*pb.GetLastPriceRequest) *pb.GetLastPriceResponse

	ListAvgPrices(*pb.ListAvgPricesRequest) *pb.ListAvgPricesResponse
	UpdateAveragePrice(*pb.UpdateAveragePriceRequest) *pb.UpdateAveragePriceResponse

	ListOperation(*pb.ListOperationRequest) *pb.ListOperationResponse
	ListOperationByPeriod(*pb.ListOperationByPeriodRequest) *pb.ListOperationByPeriodResponse
	ListAllOperation(*pb.ListAllOperationRequest) *pb.ListAllOperationResponse
	UpdateOperation(*pb.UpdateOperationRequest) *pb.UpdateOperationResponse
	CreateOperation(*pb.CreateOperationRequest) *pb.CreateOperationResponse

	GetLastBuyRegisterByOperation(*pb.GetLastBuyRegisterByOperationRequest) *pb.GetLastBuyRegisterByOperationResponse
	CreateOperationHistory(*pb.CreateOperationHistoryRequest) *pb.CreateOperationHistoryResponse

	Strategy(*pb.StrategyRequest) *pb.StrategyResponse
}

func NewUserExternal() ExternalInterface {
	conn, err := grpc.Dial(os.Getenv("MS_USER"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CMU: %v", err)
	}

	return &external{
		Conn: conn,
	}
}

func NewExchangeExternal() ExternalInterface {
	conn, err := grpc.Dial(os.Getenv("MS_EXCHANGE"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CME: %v", err)
	}

	return &external{
		Conn: conn,
	}
}
