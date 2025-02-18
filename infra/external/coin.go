package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) ListCoin(in *pb.ListCoinRequest) (out *pb.ListCoinResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListCoin(context.Background(), &pb.ListCoinRequest{})

	if err != nil {
		log.Fatalln("E-LC", err)
	}

	return
}
