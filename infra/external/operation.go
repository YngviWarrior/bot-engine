package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/bot-engine/infra/external/proto/pb"
)

func (i *external) ListAllOperation(in *pb.ListAllOperationRequest) *pb.ListAllOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.ListAllOperation(context.Background(), in)
	if err != nil {
		log.Fatalln("E-LAO", err)
	}
	return resp
}

func (i *external) GetOperation(in *pb.GetOperationRequest) *pb.GetOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.GetOperation(context.Background(), in)
	if err != nil {
		log.Fatalln("E-LAO", err)
	}
	return resp
}

func (i *external) ListOperationByPeriod(in *pb.ListOperationByPeriodRequest) *pb.ListOperationByPeriodResponse {
	client := pb.NewExchangeServiceClient(i.Conn)
	resp, err := client.ListOperationByPeriod(context.Background(), &pb.ListOperationByPeriodRequest{
		MtsStart: in.GetMtsStart(),
		MtsEnd:   in.GetMtsEnd(),
	})
	if err != nil {
		log.Fatalln("E-LOBP", err)
	}
	return resp
}

func (i *external) ListOperation(in *pb.ListOperationRequest) *pb.ListOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.ListOperation(context.Background(), &pb.ListOperationRequest{
		User:            in.GetUser(),
		Strategy:        in.GetStrategy(),
		StrategyVariant: in.GetStrategyVariant(),
		Parity:          in.GetParity(),
		Exchange:        in.GetExchange(),
		Closed:          in.GetClosed(),
		Enabled:         in.GetEnabled(),
	})

	if err != nil {
		log.Fatalln("E-LO", err)
	}

	return resp
}

func (i *external) UpdateOperation(in *pb.UpdateOperationRequest) *pb.UpdateOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.UpdateOperation(context.Background(), in)

	if err != nil {
		log.Fatalln("E-UO", err)
	}

	return resp
}

func (i *external) CreateOperation(in *pb.CreateOperationRequest) *pb.CreateOperationResponse {
	client := pb.NewExchangeServiceClient(i.Conn)

	resp, err := client.CreateOperation(context.Background(), in)

	if err != nil {
		log.Fatalln("E-CO", err)
	}

	return resp
}
