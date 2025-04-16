package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (i *external) GetTradeConfig(in *pb.GetTradeConfigRequest) *pb.GetTradeConfigResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.GetTradeConfig(context.Background(), in)

	if err != nil {
		log.Fatalln("E-GTC", err)
	}

	return resp
}

func (i *external) ListTradeConfig() *pb.TradeConfigResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.ListTradeConfig(context.Background(), &pb.ListTradeConfigRequest{})

	if err != nil {
		log.Fatalln("E-LTC", err)
	}

	return resp
}

func (i *external) UpdateTradeConfig(in *pb.UpdateTradeConfigRequest) (out *pb.UpdateTradeConfigResponse) {
	client := pb.NewExchangeServiceClient(i.Conn)

	out, err := client.UpdateTradeConfig(context.Background(), in)

	if err != nil {
		log.Fatalln("E-UTC", err)
	}

	return
}
