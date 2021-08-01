package bank

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	bankM "example.com/creditcard/app/view_card/models/bank"
	pb "example.com/creditcard/app/view_card/protos/bank"
	"example.com/creditcard/app/view_card/services/bank"
)

func NewRoute(
	server *grpc.Server,
	bankService bank.Service,
) {

	pb.RegisterBankServer(server, &BankRoute{
		bankService: bankService,
	})
}

// server is used to implement helloworld.GreeterServer.
type BankRoute struct {
	pb.UnimplementedBankServer

	bankService bank.Service
}

func (r *BankRoute) Create(ctx context.Context, in *pb.BankCreateRequest) (*pb.BankCreateReply, error) {

	bank := &bankM.Repr{
		Name: in.GetPayload().GetName(),
		Icon: in.GetPayload().GetIcon(),
	}
	if err := r.bankService.Create(ctx, bank); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.BankCreateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.BankCreateReply{
		Success: true,
		Msg:     "",
	}
	return resp, nil
}

func (r *BankRoute) Update(ctx context.Context, in *pb.BankUpdateRequest) (*pb.BankUpdateReply, error) {

	bank := &bankM.Repr{
		ID:   in.GetPayload().GetId(),
		Name: in.GetPayload().GetName(),
		Icon: in.GetPayload().GetIcon(),
	}

	if err := r.bankService.UpdateByID(ctx, bank); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.BankUpdateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.BankUpdateReply{
		Success: true,
		Msg:     "",
	}

	return resp, nil
}
func (r *BankRoute) GetAll(ctx context.Context, in *pb.BankGetAllRequest) (*pb.BankGetAllReply, error) {

	banks, err := r.bankService.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.BankGetAllReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	bankPayloads := []*pb.BankPayload{}

	for _, b := range banks {
		paylaod := &pb.BankPayload{
			Id:   b.ID,
			Name: b.Name,
			Icon: b.Icon,
		}
		bankPayloads = append(bankPayloads, paylaod)
	}

	reply := &pb.BankGetAllReply{
		Success: true,
		Msg:     "",
		Banks:   bankPayloads,
	}

	return reply, nil
}
func (r *BankRoute) GetByID(ctx context.Context, in *pb.BankGetByIDRequest) (*pb.BankGetByIDReply, error) {

	id := in.GetId()

	bank, err := r.bankService.GetByID(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.BankGetByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	payload := &pb.BankPayload{
		Id:   bank.ID,
		Name: bank.Name,
		Icon: bank.Icon,
	}

	reply := &pb.BankGetByIDReply{
		Success: true,
		Msg:     "",
		Bank:    payload,
	}

	return reply, nil
}
