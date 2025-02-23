package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) ListParity(in *pb.ListParityRequest) (out *pb.ListParityResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListParity(context.Background(), &pb.ListParityRequest{})

	if err != nil {
		log.Fatalln("E-LP", err)
	}

	return
}
