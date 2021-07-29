package common

import (
	"errors"
	"strconv"
)

type ValidateTime struct {
	StartTime int64 `json:"startDate"`
	EndTime   int64 `json:"endDate"`
}

type OperatorType int32

const (
	AND OperatorType = iota
	OR
	Union // 全部都要符合
	Cross // 其中一個符合
)

type FeatureType int32

const (
	ECommerce FeatureType = iota + 1
	Supremarket
	Delivery
	Fee
	Transport
)

func ConvertFeature(featureType int) (FeatureType, error) {
	switch featureType {
	case int(ECommerce):
		return ECommerce, nil
	case int(Supremarket):
		return Supremarket, nil
	case int(Delivery):
		return Delivery, nil
	case int(Fee):
		return Fee, nil
	case int(Transport):
		return Transport, nil
	default:
		return 0, errors.New("No match featureType: " + strconv.Itoa(featureType))
	}
}

type RewardType int32

const (
	DomesticCache RewardType = iota // 國內現金回饋
	AbroadCache                     // 國外現金回饋
	Point                           // 點數回饋
	Limited                         // 限量優惠
	FirstGift                       // 首刷禮
)
