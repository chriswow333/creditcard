package reward

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	pb "example.com/creditcard/app/view_card/protos/reward"
	"example.com/creditcard/app/view_card/services/reward"

	"example.com/creditcard/app/view_card/models/common"
	rewardM "example.com/creditcard/app/view_card/models/reward"
	taskM "example.com/creditcard/app/view_card/models/task"
)

func NewRoute(
	server *grpc.Server,
	rewardService reward.Service,
) {
	pb.RegisterRewardServer(server, &RewardRoute{
		rewardService: rewardService,
	})
}

type RewardRoute struct {
	pb.UnimplementedRewardServer
	dig.In

	rewardService reward.Service
}

func (r *RewardRoute) Create(ctx context.Context, in *pb.RewardCreateRequest) (*pb.RewardCreateReply, error) {

	rewardType, err := common.ConvertReward(in.GetPayload().GetRewardType())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardCreateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	operatorType, err := common.ConvertOperator(in.GetPayload().GetOperatorType())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardCreateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	reward := &rewardM.Repr{
		ID:     in.GetPayload().GetId(),
		Name:   in.GetPayload().GetName(),
		CardID: in.GetPayload().GetCardID(),
		Desc:   in.GetPayload().GetDesc(),

		RewardType:   rewardType,
		OperatorType: operatorType,

		StartTime: in.GetPayload().GetStartTime(),
		EndTime:   in.GetPayload().GetEndTime(),

		TotalPoint: float64(in.GetPayload().GetTotalPoint()),
	}

	tasks := []*taskM.Repr{}

	for _, t := range in.GetPayload().GetTaskPayloads() {
		task := &taskM.Repr{
			Name:     t.GetName(),
			Desc:     t.GetDesc(),
			RewardID: t.GetRewardID(),
			Point:    t.GetPoint(),
		}

		tasks = append(tasks, task)

	}

	reward.TaskReprs = tasks

	if err := r.rewardService.Create(ctx, reward); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardCreateReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.RewardCreateReply{
		Success: true,
		Msg:     "",
	}
	return resp, nil
}

func (r *RewardRoute) GetByID(ctx context.Context, in *pb.RewardGetByIDRequest) (*pb.RewardGetByIDReply, error) {

	reward, err := r.rewardService.GetByID(ctx, in.GetId())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardGetByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	payload := &pb.RewardPayload{
		Id:     reward.ID,
		Name:   reward.Name,
		CardID: reward.CardID,
		Desc:   reward.Desc,

		RewardType:   int32(reward.RewardType),
		OperatorType: int32(reward.OperatorType),
		StartTime:    reward.StartTime,
		EndTime:      reward.EndTime,
		TotalPoint:   reward.TotalPoint,
	}

	taskPayloads := []*pb.TaskPayload{}

	for _, t := range reward.TaskReprs {

		taskPayload := &pb.TaskPayload{
			Id:       t.ID,
			Name:     t.Name,
			Desc:     t.Desc,
			RewardID: t.RewardID,
			Point:    t.Point,
		}
		taskPayloads = append(taskPayloads, taskPayload)

	}

	payload.TaskPayloads = taskPayloads

	resp := &pb.RewardGetByIDReply{
		Success: true,
		Msg:     "",
		Reward:  payload,
	}
	return resp, nil
}

func (r *RewardRoute) UpdateByID(ctx context.Context, in *pb.RewardUpdateByIDRequest) (*pb.RewardUpdateByIDReply, error) {

	rewardType, err := common.ConvertReward(in.GetPayload().GetRewardType())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardUpdateByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	operatorType, err := common.ConvertOperator(in.GetPayload().GetOperatorType())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardUpdateByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	reward := &rewardM.Repr{
		ID:     in.GetPayload().GetId(),
		Name:   in.GetPayload().GetName(),
		CardID: in.GetPayload().GetCardID(),
		Desc:   in.GetPayload().GetDesc(),

		RewardType:   rewardType,
		OperatorType: operatorType,

		StartTime: in.GetPayload().GetStartTime(),
		EndTime:   in.GetPayload().GetEndTime(),

		TotalPoint: float64(in.GetPayload().GetTotalPoint()),
	}

	tasks := []*taskM.Repr{}

	for _, t := range in.GetPayload().GetTaskPayloads() {
		task := &taskM.Repr{
			Name:     t.GetName(),
			Desc:     t.GetDesc(),
			RewardID: t.GetRewardID(),
			Point:    t.GetPoint(),
		}

		tasks = append(tasks, task)

	}

	reward.TaskReprs = tasks

	if err := r.rewardService.UpdateByID(ctx, reward); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardUpdateByIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}

	resp := &pb.RewardUpdateByIDReply{
		Success: true,
		Msg:     "",
	}
	return resp, nil
}

func (r *RewardRoute) GetByCardID(ctx context.Context, in *pb.RewardGetByCardIDRequest) (*pb.RewardGetByCardIDReply, error) {

	rewards, err := r.rewardService.GetByCardID(ctx, in.GetId())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		resp := &pb.RewardGetByCardIDReply{
			Success: false,
			Msg:     "error",
		}
		return resp, err
	}
	payloads := []*pb.RewardPayload{}

	for _, reward := range rewards {

		payload := &pb.RewardPayload{
			Id:     reward.ID,
			Name:   reward.Name,
			CardID: reward.CardID,
			Desc:   reward.Desc,

			RewardType:   int32(reward.RewardType),
			OperatorType: int32(reward.OperatorType),
			StartTime:    reward.StartTime,
			EndTime:      reward.EndTime,
			TotalPoint:   reward.TotalPoint,
		}

		taskPayloads := []*pb.TaskPayload{}

		for _, t := range reward.TaskReprs {

			taskPayload := &pb.TaskPayload{
				Id:       t.ID,
				Name:     t.Name,
				Desc:     t.Desc,
				RewardID: t.RewardID,
				Point:    t.Point,
			}
			taskPayloads = append(taskPayloads, taskPayload)

		}

		payload.TaskPayloads = taskPayloads

		payloads = append(payloads, payload)
	}

	resp := &pb.RewardGetByCardIDReply{
		Success: true,
		Msg:     "",
		Reward:  payloads,
	}
	return resp, nil
}
