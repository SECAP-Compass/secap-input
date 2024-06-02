package application

import (
	"context"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/core/events"
	"secap-input/internal/domain/goal/domain/aggregate"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/port"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	jsoniter "github.com/json-iterator/go"
)

type MeasurementCalculatedConsumer struct {
	ec      *esdb.Client
	handler port.GoalUpdater
	*repository.EventRepository
	*repository.AggregateRepository
}

//const buildingMeasurementCalculatedGoal = "goal.building.measurement.calculated.consumer.group"

type dto struct {
	streamId string
	goal     *event.GoalCreatedEvent
}

func NewMeasurementCalculatedConsumer(ec *esdb.Client, handler port.GoalUpdater, eventRepository *repository.EventRepository, aggregateRepository *repository.AggregateRepository) *MeasurementCalculatedConsumer {
	c := &MeasurementCalculatedConsumer{
		ec:                  ec,
		handler:             handler,
		EventRepository:     eventRepository,
		AggregateRepository: aggregateRepository,
	}

	go c.getGoals()
	return c
}

func (c *MeasurementCalculatedConsumer) getGoals() {
	stream, err := c.ec.SubscribeToStream(context.Background(), "$et-goal.created", esdb.SubscribeToStreamOptions{
		From:           esdb.Start{},
		ResolveLinkTos: true,
	})
	if err != nil {
		panic(err)
	}

	defer stream.Close()

	for {
		e := stream.Recv()

		if e.EventAppeared != nil {
			gc := &event.GoalCreatedEvent{}
			if err := jsoniter.Unmarshal(e.EventAppeared.Event.Data, gc); err != nil {
				panic(err)
			}

			c.subscribe(e.EventAppeared.Event.StreamID)
			go c.consume(context.Background(), e.EventAppeared.Event.StreamID, uint(gc.CityId), uint(gc.DistrictId))
		}

		if e.SubscriptionDropped != nil {
			break
		}
	}
}

func (c *MeasurementCalculatedConsumer) subscribe(goalId string) {

	sOpts := esdb.SubscriptionSettingsDefault()
	sOpts.ResolveLinkTos = true

	options := esdb.PersistentStreamSubscriptionOptions{
		StartFrom: esdb.Start{},
		Settings:  &sOpts,
	}

	err := c.ec.CreatePersistentSubscription(context.Background(), "$et-building.measurement.calculated", goalId, options)
	if err != nil && !strings.Contains(err.Error(), "exists") {
		panic(err)
	}
}

func (c *MeasurementCalculatedConsumer) consume(ctx context.Context, goalId string, cityId, districtId uint) {
	sub, err := c.ec.SubscribeToPersistentSubscription(
		context.Background(), "$et-building.measurement.calculated", goalId, esdb.SubscribeToPersistentSubscriptionOptions{},
	)

	if err != nil {
		panic(err)
	}

	// TODO: Here may consume concurrently?
	for {
		esdbEvent := sub.Recv()
		if esdbEvent.EventAppeared != nil {
			re, err := c.GetFirstEvent(ctx, esdbEvent.EventAppeared.Event.Event.StreamID)
			if err != nil {
				panic(err)
			}

			bce := &events.BuildingCreatedEvent{}
			if err := jsoniter.Unmarshal(re.Data, bce); err != nil {
				panic(err)
			}

			a := aggregate.NewGoalAggregateWithId(goalId)
			if bce.Address.Province.Id == cityId {
				if err := c.Load(context.Background(), a); err != nil {
					panic(err)
				}

				e := &event.MeasurementCalculatedEvent{}
				if err := jsoniter.Unmarshal(esdbEvent.EventAppeared.Event.Event.Data, e); err != nil {
					panic(err)
				}

				if err := c.handler.Handle(ctx, e, a); err != nil {
					sub.Nack("failed to process goal event", esdb.NackActionRetry, esdbEvent.EventAppeared.Event)
				}

				if err := c.AggregateRepository.Save(ctx, a); err != nil {
					sub.Nack("failed to save aggregate", esdb.NackActionRetry, esdbEvent.EventAppeared.Event)
				}

				sub.Ack(esdbEvent.EventAppeared.Event)
			}

			//// Province ve district icin ayri ayri ilerletmeliyiz.
			//if bce.Address.Province.Id == districtId {
			//	var e *event.MeasurementCalculatedEvent
			//	if err := jsoniter.Unmarshal(esdbEvent.EventAppeared.Event.OriginalEvent().Data, e); err != nil {
			//		panic(err)
			//	}
			//
			//	c.handler.Handle(context.Background(), e, a)
			//}

			sub.Ack(esdbEvent.EventAppeared.Event)
		}

		if esdbEvent.SubscriptionDropped != nil {
			break
		}
	}

	sub.Close()
}
