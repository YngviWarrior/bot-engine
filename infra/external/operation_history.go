package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (i *external) ListOperationHistoryByOperation(in *pb.ListOperationHistoryByOperationRequest) *pb.ListOperationHistoryByOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.ListOperationHistoryByOperation(context.Background(), in)
	if err != nil {
		log.Fatalln("E-LOHBO", err)
	}
	return resp
}

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

func (i *external) GetOperationHistory(in *pb.GetOperationHistoryRequest) *pb.GetOperationHistoryResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.GetOperationHistory(context.Background(), in)
	if err != nil {
		log.Fatalln("E-GOH", err)
	}
	return resp
}

func (i *external) UpdateOperationHistory(in *pb.UpdateOperationHistoryRequest) *pb.UpdateOperationHistoryResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.UpdateOperationHistory(context.Background(), in)
	if err != nil {
		log.Fatalln("E-UOH", err)
	}
	return resp
}
