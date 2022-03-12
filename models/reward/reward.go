package reward

import (
	"time"

	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/payload"
)

type RewardType int32

const (
	InCash RewardType = iota
	OutCash
	Point
)

type PayloadOperator int32

const (
	AddPayloadOperator PayloadOperator = iota + 1
	XORHighPayloadOperator
)

type Reward struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardID"`

	Order int32 `json:"order"`

	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`
	// RewardType RewardType `json:"rewardType"`

	PayloadOperator PayloadOperator    `json:"payloadOperator"`
	Payloads        []*payload.Payload `json:"payloads"`
}

type RewardResp struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardID"`
	Order        int32  `json:"order"`

	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`

	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	UpdateDate string `json:"updateDate"`

	// RewardType RewardType `json:"rewardType"`

	PayloadOperator PayloadOperator        `json:"payloadOperator"`
	PayloadResps    []*payload.PayloadResp `json:"payloadResps"`
	FeedReturn      *feedback.FeedReturn   `json:"feedReturn"`

	FeedbackResp *feedback.FeedbackResp `json:"feedbackResp"`
}

const DATE_FORMAT = "2006/01/02"

func TransferRewardResp(rewardType RewardType, reward *Reward) *RewardResp {
	rewardResp := &RewardResp{}

	rewardResp.ID = reward.ID
	rewardResp.CardRewardID = reward.CardRewardID
	rewardResp.Order = reward.Order
	rewardResp.Title = reward.Title
	rewardResp.SubTitle = reward.SubTitle
	rewardResp.StartDate = time.Unix(reward.StartDate, 0).Format(DATE_FORMAT)
	rewardResp.EndDate = time.Unix(reward.EndDate, 0).Format(DATE_FORMAT)
	rewardResp.UpdateDate = time.Unix(reward.UpdateDate, 0).Format(DATE_FORMAT)
	rewardResp.PayloadOperator = reward.PayloadOperator

	payloadResps := []*payload.PayloadResp{}
	for _, p := range reward.Payloads {
		payloadResp := payload.TransferPayloadResp(p)
		payloadResps = append(payloadResps, payloadResp)
	}
	rewardResp.PayloadResps = payloadResps

	switch rewardType {
	case InCash:
		cashbackResp := getOptimalCashbackResp(reward.PayloadOperator, payloadResps)
		rewardResp.FeedbackResp = &feedback.FeedbackResp{
			CashbackResp: cashbackResp,
		}
	case OutCash:
		cashbackResp := getOptimalCashbackResp(reward.PayloadOperator, payloadResps)
		rewardResp.FeedbackResp = &feedback.FeedbackResp{
			CashbackResp: cashbackResp,
		}
	case Point:
	}

	return rewardResp
}

func getOptimalCashbackResp(payloadOperator PayloadOperator, payloadResps []*payload.PayloadResp) *feedback.CashbackResp {
	cashbackResp := &feedback.CashbackResp{}
	var min int64 = 0
	var max int64 = 99999999
	var bonus float64 = 0.0
	var cashbackType feedback.CashbackType

	for _, p := range payloadResps {
		cashbackType = p.FeedbackResp.CashbackResp.CashbackType
		switch payloadOperator {
		case XORHighPayloadOperator:
			if p.FeedbackResp.CashbackResp.Bonus < bonus {
				min = p.FeedbackResp.CashbackResp.Min
				max = p.FeedbackResp.CashbackResp.Max
				bonus = p.FeedbackResp.CashbackResp.Bonus
			}

		case AddPayloadOperator:

			bonus += p.FeedbackResp.CashbackResp.Bonus
			if min < p.FeedbackResp.CashbackResp.Min {
				min = p.FeedbackResp.CashbackResp.Min
			}

			if max > p.FeedbackResp.CashbackResp.Max {
				max = p.FeedbackResp.CashbackResp.Max
			}

		}
	}

	cashbackResp.CashbackType = cashbackType
	cashbackResp.Min = min
	cashbackResp.Max = max
	cashbackResp.Bonus = bonus

	return cashbackResp
}

type OutCashRewardResp struct {
	FeedReturn *feedback.FeedReturn `json:"feedReturn"`

	FeedbackResp *feedback.FeedbackResp `json:"feedbackResp"`

	RewardResps []*RewardResp `json:"rewardResps"`
}

type InCashRewardResp struct {
	FeedReturn *feedback.FeedReturn `json:"feedReturn"`

	FeedbackResp *feedback.FeedbackResp `json:"feedbackResp"`

	RewardResps []*RewardResp `json:"rewardResps"`
}
