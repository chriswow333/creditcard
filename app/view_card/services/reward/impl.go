package reward

import (
	"context"
	"time"

	rewardM "example.com/creditcard/app/view_card/models/reward"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"example.com/creditcard/app/view_card/models/common"
	taskM "example.com/creditcard/app/view_card/models/task"
	"example.com/creditcard/app/view_card/stores/reward"
	"example.com/creditcard/app/view_card/stores/task"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In
	rewardStore reward.Store
	taskStore   task.Store
}

func New(
	rewardStore reward.Store,
	taskStore task.Store,
) Service {
	return &impl{
		rewardStore: rewardStore,
		taskStore:   taskStore,
	}
}

func (im *impl) Create(ctx context.Context, rewardRepr *rewardM.Repr) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	validateTime := common.ValidateTime{
		StartTime: rewardRepr.StartTime,
		EndTime:   rewardRepr.EndTime,
	}
	reward := &rewardM.Reward{
		ID:           id.String(),
		Name:         rewardRepr.Name,
		CardID:       rewardRepr.CardID,
		Desc:         rewardRepr.Desc,
		RewardType:   rewardRepr.RewardType,
		OperatorType: rewardRepr.OperatorType,
		ValidateTime: validateTime,
		TotalPoint:   rewardRepr.TotalPoint,
		UpdateDate:   timeNow().Unix(),
	}

	if err := im.rewardStore.Create(ctx, reward); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	tasks := []*taskM.Task{}
	for _, t := range rewardRepr.TaskReprs {

		id, err := uuid.NewV4()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return err
		}

		task := &taskM.Task{
			ID:         id.String(),
			Name:       t.Name,
			Desc:       t.Desc,
			RewardID:   reward.ID,
			Point:      t.Point,
			UpdateDate: timeNow().Unix(),
		}

		tasks = append(tasks, task)

	}

	if err := im.taskStore.CreateTasks(ctx, tasks); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Repr, error) {

	reward, err := im.rewardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	rewardRepr := &rewardM.Repr{
		ID:           reward.ID,
		Name:         reward.Name,
		CardID:       reward.CardID,
		Desc:         reward.Desc,
		RewardType:   reward.RewardType,
		OperatorType: reward.OperatorType,
		StartTime:    reward.ValidateTime.StartTime,
		EndTime:      reward.ValidateTime.EndTime,
		TotalPoint:   reward.TotalPoint,
	}

	teakReprs, err := im.getTaskByRewardID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}
	rewardRepr.TaskReprs = teakReprs

	return rewardRepr, nil
}

func (im *impl) UpdateByID(ctx context.Context, rewardRepr *rewardM.Repr) error {

	return nil
}

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Repr, error) {

	rewards, err := im.rewardStore.GetByCardID(ctx, cardID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	rewardReprs := []*rewardM.Repr{}
	for _, r := range rewards {
		rewardRepr := &rewardM.Repr{
			ID:           r.ID,
			Name:         r.Name,
			CardID:       r.CardID,
			Desc:         r.Desc,
			RewardType:   r.RewardType,
			OperatorType: r.OperatorType,
			StartTime:    r.ValidateTime.StartTime,
			EndTime:      r.ValidateTime.EndTime,
			TotalPoint:   r.TotalPoint,
		}

		rewardReprs = append(rewardReprs, rewardRepr)
	}

	for _, r := range rewardReprs {
		taskReprs, err := im.getTaskByRewardID(ctx, r.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return nil, err
		}
		r.TaskReprs = taskReprs
	}

	return rewardReprs, nil
}

func (im *impl) getTaskByRewardID(ctx context.Context, rewardID string) ([]*taskM.Repr, error) {
	tasks, err := im.taskStore.GetByRewardID(ctx, rewardID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	tasksReprs := []*taskM.Repr{}

	for _, t := range tasks {
		task := &taskM.Repr{
			ID:       t.ID,
			Name:     t.Name,
			Desc:     t.Desc,
			RewardID: rewardID,
			Point:    t.Point,
		}

		tasksReprs = append(tasksReprs, task)
	}
	return tasksReprs, nil
}
