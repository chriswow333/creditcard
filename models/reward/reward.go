package reward

import (
	"fmt"
	"time"

	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/payload"
)

type RewardType int32

const (
	CASH_TWD RewardType = iota + 1
	POINT
)

type PayloadOperator int32

const (
	ADD PayloadOperator = iota + 1
	MAXONE
)

type Reward struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardID"`
	Order        int32  `json:"order"`

	Title    string `json:"title,omitempty"`
	SubTitle string `json:"subTitle,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	PayloadOperator PayloadOperator    `json:"payloadOperator,omitempty"`
	Payloads        []*payload.Payload `json:"payloads,omitempty"`
}

type RewardResp struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardID"`
	Order        int32  `json:"order"`

	Title    string `json:"title,omitempty"`
	SubTitle string `json:"subTitle,omitempty"`

	StartDate  string `json:"startDate,omitempty"`
	EndDate    string `json:"endDate,omitempty"`
	UpdateDate string `json:"updateDate,omitempty"`

	PayloadOperator PayloadOperator `json:"payloadOperator,omitempty"`

	PayloadResps []*payload.PayloadResp `json:"payloadResps,omitempty"`
	Feedback     *feedback.Feedback     `json:"feedback,omitempty"`
}

type RewardEventJudgeType int32

const (
	ALL RewardEventJudgeType = iota + 1
	SOME
	NONE
)

type RewardEventResp struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardRewardID"`

	Order int32 `json:"order"`

	RewardEventJudgeType RewardEventJudgeType `json:"rewardEventJudgeType,omitempty"`

	PayloadOperator PayloadOperator `json:"payloadOperator,omitempty"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	PayloadEventResps []*payload.PayloadEventResp `json:"payloadEventResps,omitempty"`
}

const DATE_FORMAT = "2006/01/02"

func TransferRewardResp(rewardType RewardType, reward *Reward) *RewardResp {

	rewardResp := &RewardResp{
		ID:              reward.ID,
		CardRewardID:    reward.CardRewardID,
		Order:           reward.Order,
		Title:           reward.Title,
		SubTitle:        reward.SubTitle,
		StartDate:       time.Unix(reward.StartDate, 0).Format(DATE_FORMAT),
		EndDate:         time.Unix(reward.EndDate, 0).Format(DATE_FORMAT),
		UpdateDate:      time.Unix(reward.UpdateDate, 0).Format(DATE_FORMAT),
		PayloadOperator: reward.PayloadOperator,
	}

	payloadResps := []*payload.PayloadResp{}

	for _, p := range reward.Payloads {
		// constraintResp, err := constraintM.TransferConstraintResp(ctx, p.Constraint, constraintSvc)
		// if err != nil {
		// 	logrus.Error("constraintM.TransferConstraintResp")
		// 	return nil, err
		// }
		payloadResp := payload.TransferPayloadResp(p)
		payloadResps = append(payloadResps, payloadResp)
	}

	rewardResp.PayloadResps = payloadResps

	switch rewardType {
	case CASH_TWD:
		cashback := getOptimalCashback(reward.PayloadOperator, payloadResps)
		rewardResp.Feedback = &feedback.Feedback{
			Cashback: cashback,
		}
	case POINT:
	}

	return rewardResp
}

func getOptimalCashback(payloadOperator PayloadOperator, payloadResps []*payload.PayloadResp) *feedback.Cashback {
	cashback := &feedback.Cashback{}
	var min int64 = 0
	var max int64 = 9999999
	var bonus float64 = 0.0
	var cashbackType feedback.CashbackType

	fmt.Println("payload operator ", payloadOperator)

	for _, p := range payloadResps {
		cashbackType = p.Feedback.Cashback.CashbackType
		fmt.Println(cashbackType)

		switch payloadOperator {

		case ADD:
			bonus += p.Feedback.Cashback.Bonus
			if min < p.Feedback.Cashback.Min {
				min = p.Feedback.Cashback.Min
			}

			if max > p.Feedback.Cashback.Max {
				max = p.Feedback.Cashback.Max
			}
			continue
		case MAXONE:
			if p.Feedback.Cashback.Bonus > bonus {
				min = p.Feedback.Cashback.Min
				max = p.Feedback.Cashback.Max
				bonus = p.Feedback.Cashback.Bonus
			}
			fmt.Println(bonus)
		default:

		}

	}

	cashback.CashbackType = cashbackType
	cashback.Min = min
	cashback.Max = max
	cashback.Bonus = bonus

	return cashback
}
