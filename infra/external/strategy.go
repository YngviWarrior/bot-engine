package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) Strategy(in *pb.StrategyRequest) (out *pb.StrategyResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.Strategy(context.Background(), in)
	if err != nil {
		log.Fatalln("E-SAP", err)
	}

	return
}
