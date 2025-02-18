package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (i *external) ListUserStrategy() *pb.UserStrategyResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.ListUserStrategy(context.Background(), &pb.ListUserStrategyRequest{})

	if err != nil {
		log.Fatalln("E-GUS", err)
	}

	return resp
}
