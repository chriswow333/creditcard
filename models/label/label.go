package label

type LabelType int32

const (
	ALL     = iota // 不分通路
	Channel        // by 分通路別, 不適用貼label, 直接看channel type
// 上述兩個不須向各通路做查詢

// 實體通路
// MICRO_PAYMENT              // 小額支付
// OVERSEA                    // 海外消費
// GENERAL_CONSUMPTION        // 一般消費
// TW_RESTAURANT              // 全臺餐廳

)

type Label struct {
	Match bool `json:"match"`

	LabelType LabelType `json:"labelType"` // 根據 channelLabeltype 跟通路去做 match or mismatch

	// if label type is Channel, use this
	ChannelTypes []int32 `json:"channelTypes"` // 根據 label type is channel type && channelType is match or not

}
