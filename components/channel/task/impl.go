package customization

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/task"
	channelSvc "example.com/creditcard/service/channel"
	"github.com/sirupsen/logrus"
)

type impl struct {
	channel        *channelM.Channel
	tasks          []*task.Task
	channelService channelSvc.Service
}

func New(
	channel *channelM.Channel,
	tasks []*task.Task,
	channelService channelSvc.Service,

) channel.Component {

	return &impl{
		channel:        channel,
		tasks:          tasks,
		channelService: channelService,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.TaskType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}

	eventTaskMap := make(map[string]bool)

	for _, t := range e.Tasks {
		eventTaskMap[t] = true
	}

	for _, t := range im.tasks {

		taskType := t.TaskType

		pass := false

		switch taskType {

		case task.NONE:
			pass = im.processNoneType(t, eventTaskMap)
			break
		case task.WEEKDAY:
			pass = im.processWeekDayType(e, t, eventTaskMap)
			break
		case task.CHANNEL_LABEL:
			pass = im.processChannelLabel(ctx, e, t, eventTaskMap)
			break
		case task.CHANNEL:
			pass = im.processChannel(e, t, eventTaskMap)
			break
		default:
			return nil, errors.New("not found task type")
		}

		if pass {
			matches = append(matches, t.ID)
		} else {
			misses = append(misses, t.ID)
		}
	}

	channelEventResp.Matches = matches
	channelEventResp.Misses = misses

	switch im.channel.ChannelOperatorType {
	case channelM.OR:
		if len(matches) > 0 {
			channelEventResp.Pass = true
		} else {
			channelEventResp.Pass = false
		}
		break
	case channelM.AND:
		if len(misses) > 0 || len(matches) == 0 {
			channelEventResp.Pass = false
		} else {
			channelEventResp.Pass = true
		}
		break
	default:
		return nil, errors.New("not found channel operatortype")
	}

	if im.channel.ChannelMappingType == channelM.MISMATCH {
		channelEventResp.Pass = !channelEventResp.Pass
	}

	return channelEventResp, nil
}

func (im *impl) processNoneType(t *task.Task, eventTaskMap map[string]bool) bool {
	if _, ok := eventTaskMap[t.ID]; ok {
		return true
	} else if t.DefaultPass {
		return true
	} else {
		return false
	}

}

func (im *impl) processWeekDayType(e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	if _, ok := taskMap[t.ID]; ok {
		return true
	} else {
		weekDay := time.Unix(e.EffectiveTime, 0).Weekday()

		weekdayLimit := t.TaskTypeModel.WeekDayLimit

		for _, d := range weekdayLimit.WeekDays {
			if d == int(weekDay) {
				return true
			}
		}
		return false
	}

}

func (im *impl) processChannel(e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	if _, ok := taskMap[t.ID]; ok {
		return true
	} else {

		channelLimit := t.TaskTypeModel.ChannelLimit
		for _, ch := range channelLimit.Channels {
			switch ch {
			case int32(channelM.MobilepayType):
				if len(e.Mobilepays) > 0 {
					return true
				}
				break
			case int32(channelM.EcommerceType):
				if len(e.Ecommerces) > 0 {
					return true
				}
				break

			case int32(channelM.SupermarketType):
				if len(e.Supermarkets) > 0 {
					return true
				}
				break
			case int32(channelM.OnlinegameType):
				if len(e.Onlinegames) > 0 {
					return true
				}
				break
			case int32(channelM.StreamingType):
				if len(e.Streamings) > 0 {
					return true
				}
				break

			case int32(channelM.FoodType):
				if len(e.Foods) > 0 {
					return true
				}
				break
			case int32(channelM.TransportationType):
				if len(e.Transportations) > 0 {
					return true
				}
				break
			case int32(channelM.TravelType):
				if len(e.Travels) > 0 {
					return true
				}
				break

			case int32(channelM.DeliveryType):
				if len(e.Deliveries) > 0 {
					return true
				}
				break
			case int32(channelM.InsuranceType):
				if len(e.Insurances) > 0 {
					return true
				}
				break
			case int32(channelM.MallType):
				if len(e.Malls) > 0 {
					return true
				}
				break

			case int32(channelM.SportType):
				if len(e.Sports) > 0 {
					return true
				}
				break
			case int32(channelM.ConvenienceStoreType):
				if len(e.Conveniencestores) > 0 {
					return true
				}
				break
			case int32(channelM.AppStoreType):
				if len(e.AppStores) > 0 {
					return true
				}
				break

			case int32(channelM.HotelType):
				if len(e.Hotels) > 0 {
					return true
				}
				break
			case int32(channelM.AmusementType):
				if len(e.Amusements) > 0 {
					return true
				}
				break
			case int32(channelM.CinemaType):
				if len(e.Cinemas) > 0 {
					return true
				}
				break

			case int32(channelM.PublicUtilityType):
				if len(e.Publicutilities) > 0 {
					return true
				}
				break
			}

		}
		return false
	}

}

func (im *impl) processChannelLabel(ctx context.Context, e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	if _, ok := taskMap[t.ID]; ok {
		return true
	} else {

		channelLabelLimit := t.TaskTypeModel.ChannelLabelLimit

		channelLabels := channelLabelLimit.ChannelLabels
		channelLabelMap := map[channelM.ChannelLabel]bool{}

		for _, label := range channelLabels {
			switch label {
			case int32(channelM.MICRO_PAYMENT):
				channelLabelMap[channelM.MICRO_PAYMENT] = true
				break
			case int32(channelM.OVERSEA):
				channelLabelMap[channelM.OVERSEA] = true
				break
			case int32(channelM.GENERAL_CONSUMPTION):
				channelLabelMap[channelM.GENERAL_CONSUMPTION] = true
				break
			case int32(channelM.TW_RESTAURANT):
				channelLabelMap[channelM.TW_RESTAURANT] = true
			}
		}

		if len(e.Mobilepays) > 0 {
			for _, m := range e.Mobilepays {

				mobile, err := im.channelService.GetMobilepay(ctx, m)

				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range mobile.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Ecommerces) > 0 {

			for _, e := range e.Ecommerces {
				ecommerce, err := im.channelService.GetEcommerce(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range ecommerce.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Supermarkets) > 0 {
			for _, s := range e.Supermarkets {
				supermarket, err := im.channelService.GetSupermarket(ctx, s)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}
				for _, cl := range supermarket.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Onlinegames) > 0 {

			for _, e := range e.Onlinegames {
				onlinegame, err := im.channelService.GetOnlinegame(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range onlinegame.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Streamings) > 0 {

			for _, e := range e.Streamings {
				streaming, err := im.channelService.GetStreaming(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range streaming.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Foods) > 0 {

			for _, e := range e.Foods {
				food, err := im.channelService.GetFood(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range food.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Transportations) > 0 {

			for _, e := range e.Transportations {
				transportation, err := im.channelService.GetTransportation(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range transportation.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Travels) > 0 {

			for _, e := range e.Travels {
				travel, err := im.channelService.GetTravel(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range travel.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Deliveries) > 0 {

			for _, e := range e.Deliveries {
				delivery, err := im.channelService.GetDelivery(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range delivery.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Insurances) > 0 {

			for _, e := range e.Insurances {
				insurance, err := im.channelService.GetInsurance(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range insurance.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Malls) > 0 {

			for _, e := range e.Malls {
				mall, err := im.channelService.GetMall(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range mall.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Sports) > 0 {

			for _, e := range e.Sports {
				sport, err := im.channelService.GetSport(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range sport.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Conveniencestores) > 0 {

			for _, e := range e.Conveniencestores {
				convenienceStore, err := im.channelService.GetConvenienceStore(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range convenienceStore.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.AppStores) > 0 {

			for _, e := range e.AppStores {
				appstore, err := im.channelService.GetAppstore(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range appstore.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Hotels) > 0 {

			for _, e := range e.Hotels {
				hotel, err := im.channelService.GetHotel(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range hotel.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Amusements) > 0 {

			for _, e := range e.Amusements {
				amusement, err := im.channelService.GetAmusement(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range amusement.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Cinemas) > 0 {

			for _, e := range e.Cinemas {
				cinema, err := im.channelService.GetCinema(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range cinema.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

		if len(e.Publicutilities) > 0 {
			for _, e := range e.Publicutilities {
				publicutility, err := im.channelService.GetPublicUtility(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					continue
				}

				for _, cl := range publicutility.ChannelLabels {
					if _, ok := channelLabelMap[cl]; ok {
						return true
					}
				}
			}
		}

	}

	return false
}
