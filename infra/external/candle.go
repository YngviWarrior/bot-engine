package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) ListCandleLimit(in *pb.ListCandleLimitRequest) (out *pb.ListCandleLimitResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListCandleLimit(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GKFM", err)
	}

	return
}

func (e *external) GetCandleFirstMts(in *pb.GetCandleFirstMtsRequest) (out *pb.GetCandleFirstMtsResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetCandleFirstMts(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GKFM", err)
	}

	return
}

func (e *external) CreateCandles(in *pb.CreateCandlesRequest) (out *pb.CreateCandlesResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.CreateCandles(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GKFM", err)
	}

	return
}
