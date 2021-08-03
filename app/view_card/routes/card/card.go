package card

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/models/common"
	pb "example.com/creditcard/app/view_card/protos/card"
	"example.com/creditcard/app/view_card/services/card"
)

func NewRoute(
	server *grpc.Server,
	cardService card.Service,
) {

	pb.RegisterCardServer(server, &CardRoute{
		cardService: cardService,
	})
}

type CardRoute struct {
	pb.UnimplementedCardServer
	dig.In

	cardService card.Service
}

func (r *CardRoute) Create(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardCreateReply, error) {

	card := &cardM.Repr{
		Name:     in.GetPayload().GetName(),
		Icon:     in.GetPayload().GetIcon(),
		BankID:   in.Payload.GetBankID(),
		MaxPoint: float64(in.GetPayload().GetMaxPoint()),

		Features:    converFeratureType(in.GetPayload().GetFeatures()),
		FeatureDesc: in.GetPayload().GetFeatureDedsc(),

		StartTime: in.GetPayload().GetStartTime(),
		EndTime:   in.GetPayload().GetEndTime(),

		ApplicantQualifications: in.GetPayload().GetApplicationQualifications(),
	}
	if err := r.cardService.Create(ctx, card); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.CardCreateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.CardCreateReply{
		Success: true,
		Msg:     "",
	}
	return resp, nil
}

func (r *CardRoute) GetByID(ctx context.Context, in *pb.CardGetByIDRequest) (*pb.CardGetByIDReply, error) {

	id := in.GetId()
	card, err := r.cardService.GetByID(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.CardGetByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	payload := &pb.CardPayload{
		Id:                        card.ID,
		Name:                      card.Name,
		Icon:                      card.Icon,
		BankID:                    card.BankID,
		MaxPoint:                  card.MaxPoint,
		Features:                  convertFeatureInt(card.Features),
		FeatureDedsc:              card.FeatureDesc,
		StartTime:                 card.StartTime,
		EndTime:                   card.EndTime,
		ApplicationQualifications: card.ApplicantQualifications,
	}

	resp := &pb.CardGetByIDReply{
		Success: true,
		Msg:     "error",

		Card: payload,
	}
	return resp, nil
}

func (r *CardRoute) UpdateByID(ctx context.Context, in *pb.CardUpdateByIDRequest) (*pb.CardUpdateByIDReply, error) {

	card := &cardM.Repr{
		Name:     in.GetPayload().GetName(),
		Icon:     in.GetPayload().GetIcon(),
		BankID:   in.Payload.GetBankID(),
		MaxPoint: float64(in.GetPayload().GetMaxPoint()),

		Features:    converFeratureType(in.GetPayload().GetFeatures()),
		FeatureDesc: in.GetPayload().GetFeatureDedsc(),

		StartTime: in.GetPayload().GetStartTime(),
		EndTime:   in.GetPayload().GetEndTime(),

		ApplicantQualifications: in.GetPayload().GetApplicationQualifications(),
	}
	if err := r.cardService.Create(ctx, card); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.CardUpdateByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.CardUpdateByIDReply{
		Success: true,
		Msg:     "",
	}
	return resp, nil

}

func (r *CardRoute) GetAll(ctx context.Context, in *pb.CardGetAllRequest) (*pb.CardGetAllReply, error) {

	cards, err := r.cardService.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.CardGetAllReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	payloads := []*pb.CardPayload{}

	for _, c := range cards {
		payload := &pb.CardPayload{
			Id:                        c.ID,
			Name:                      c.Name,
			Icon:                      c.Icon,
			BankID:                    c.BankID,
			MaxPoint:                  c.MaxPoint,
			Features:                  convertFeatureInt(c.Features),
			FeatureDedsc:              c.FeatureDesc,
			StartTime:                 c.StartTime,
			EndTime:                   c.EndTime,
			ApplicationQualifications: c.ApplicantQualifications,
		}
		payloads = append(payloads, payload)
	}

	resp := &pb.CardGetAllReply{
		Success: true,
		Msg:     "",
		Card:    payloads,
	}

	return resp, nil
}

func (r *CardRoute) GetByBankID(ctx context.Context, in *pb.CardGetByBankIDRequest) (*pb.CardGetByBankIDReply, error) {

	id := in.GetId()
	cards, err := r.cardService.GetByBankID(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.CardGetByBankIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	payloads := []*pb.CardPayload{}

	for _, c := range cards {
		payload := &pb.CardPayload{
			Id:                        c.ID,
			Name:                      c.Name,
			Icon:                      c.Icon,
			BankID:                    c.BankID,
			MaxPoint:                  c.MaxPoint,
			Features:                  convertFeatureInt(c.Features),
			FeatureDedsc:              c.FeatureDesc,
			StartTime:                 c.StartTime,
			EndTime:                   c.EndTime,
			ApplicationQualifications: c.ApplicantQualifications,
		}
		payloads = append(payloads, payload)
	}

	resp := &pb.CardGetByBankIDReply{
		Success: true,
		Msg:     "",
		Card:    payloads,
	}

	return resp, nil
}

func converFeratureType(features []int32) []common.FeatureType {
	featureTypes := []common.FeatureType{}

	for _, f := range features {
		common.ConvertFeature(f)
		featureTypes = append(featureTypes, common.FeatureType(f))
	}
	return featureTypes
}

func convertFeatureInt(featureTypes []common.FeatureType) []int32 {

	var features []int32
	for _, f := range featureTypes {
		features = append(features, int32(f))
	}

	return features
}
