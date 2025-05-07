package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (e *external) ListTransactionType(in *pb.ListTransactionTypeRequest) (out *pb.ListTransactionTypeResponse) {
	client := pb.NewExchangeServiceClient(e.Conn)

	out, err := client.ListTransactionType(context.Background(), in)

	if err != nil {
		log.Fatalln("E-LP", err)
	}

	return
}
