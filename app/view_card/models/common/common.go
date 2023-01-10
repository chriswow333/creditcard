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

func ConvertOperator(operatorType int32) (OperatorType, error) {

	switch operatorType {
	case int32(AND):
		return AND, nil
	case int32(OR):
		return OR, nil
	case int32(Union):
		return Union, nil
	case int32(Cross):
		return Cross, nil
	default:
		return 0, errors.New("ConvertOperator No match operatorType: " + strconv.Itoa(int(operatorType)))
	}
}

type FeatureType int32

const (
	ECommerce FeatureType = iota + 1
	Supremarket
	Delivery
	Fee
	Transport
)

func ConvertFeature(featureType int32) (FeatureType, error) {

	switch featureType {
	case int32(ECommerce):
		return ECommerce, nil
	case int32(Supremarket):
		return Supremarket, nil
	case int32(Delivery):
		return Delivery, nil
	case int32(Fee):
		return Fee, nil
	case int32(Transport):
		return Transport, nil
	default:
		return 0, errors.New("ConverFeature No match featureType: " + strconv.Itoa(int(featureType)))
	}
}

type RewardType int32

const (
	DomesticCache RewardType = iota // 國內現金回饋
	AbroadCache                     // 國外現金回饋
	Point                           // 點數回饋
	// Limited                         // 限量優惠
	// FirstGift                       // 首刷禮
)

func ConvertReward(rewardType int32) (RewardType, error) {

	switch rewardType {
	case int32(DomesticCache):
		return DomesticCache, nil
	case int32(AbroadCache):
		return AbroadCache, nil
	case int32(Point):
		return Point, nil
	// case int32(Limited):
	// 	return Limited, nil
	// case int32(FirstGift):
	// 	return FirstGift, nil
	default:
		return 0, errors.New("ConvertReward No match rewardType: " + strconv.Itoa(int(rewardType)))
	}

}
