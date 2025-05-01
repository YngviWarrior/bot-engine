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

func (i *external) CreateOperationHistory(in *pb.CreateOperationHistoryRequest) *pb.CreateOperationHistoryResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.CreateOperationHistory(context.Background(), in)
	if err != nil {
		log.Fatalln("E-COH", err)
	}
	return resp
}
