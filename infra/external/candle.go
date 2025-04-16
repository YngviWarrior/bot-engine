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
		log.Fatalln("E-LCL", err)
	}

	return
}

func (e *external) GetCandleFirstMts(in *pb.GetCandleFirstMtsRequest) (out *pb.GetCandleFirstMtsResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetCandleFirstMts(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GCFM", err)
	}

	return
}

func (e *external) CreateCandles(in *pb.CreateCandlesRequest) (out *pb.CreateCandlesResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.CreateCandles(context.Background(), in)
	if err != nil {
		log.Fatalln("E-CC", err)
	}

	return
}

func (e *external) GetFirstPrice(in *pb.GetFirstPriceRequest) (out *pb.GetFirstPriceResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetFirstPrice(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GFP", err)
	}

	return
}

func (e *external) GetLastPrice(in *pb.GetLastPriceRequest) (out *pb.GetLastPriceResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.GetLastPrice(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GLP", err)
	}

	return
}
