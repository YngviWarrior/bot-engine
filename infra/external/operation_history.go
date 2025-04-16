package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (i *external) GetLastBuyRegisterByOperation(in *pb.GetLastBuyRegisterByOperationRequest) *pb.GetLastBuyRegisterByOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.GetLastBuyRegisterByOperation(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GLBRBO", err)
	}
	return resp
}
