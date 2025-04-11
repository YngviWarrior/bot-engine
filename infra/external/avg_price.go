package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) GetAvgPriceByParityExchange(in *pb.GetAvgPriceByParityExchangeRequest) (out *pb.GetAvgPriceByParityExchangeResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetAvgPriceByParityExchange(context.Background(), in)

	if err != nil {
		log.Fatalln("E-LP", err)
	}

	return
}

func (e *external) ListAvgPrices(in *pb.ListAvgPricesRequest) (out *pb.ListAvgPricesResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListAvgPrices(context.Background(), in)
	if err != nil {
		log.Fatalln("E-LAP", err)
	}

	return
}
