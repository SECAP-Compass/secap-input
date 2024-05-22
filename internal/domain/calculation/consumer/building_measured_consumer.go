package consumer

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	jsoniter "github.com/json-iterator/go"
	"secap-input/internal/domain/calculation/domain/event"
	"secap-input/internal/domain/calculation/domain/port"
	"strings"
)

const buildingMeasured = "$et-building.measured"
const buildingMeasuredConsumerGroup = "buildingMeasuredConsumerGroup"

type buildingMeasuredConsumer struct {
	ec      *esdb.Client
	handler port.BuildingMeasuredHandler
}

func NewBuildingMeasuredConsumer(ec *esdb.Client, handler port.BuildingMeasuredHandler) port.BuildingMeasuredConsumer {
	consumer := &buildingMeasuredConsumer{ec: ec, handler: handler}
	consumer.createSubscription()
	go consumer.Consume()

	return consumer
}

func (b *buildingMeasuredConsumer) createSubscription() {
	options := esdb.PersistentAllSubscriptionOptions{
		StartFrom: esdb.Start{},
		Filter: &esdb.SubscriptionFilter{
			Type:     esdb.EventFilterType,
			Prefixes: []string{"building.measured"},
		},
	}

	err := b.ec.CreatePersistentSubscriptionToAll(context.Background(), buildingMeasuredConsumerGroup, options)
	if err != nil && !strings.Contains(err.Error(), "exists") {
		panic(err)
	}
}

func (b *buildingMeasuredConsumer) Consume() {
	sub, err := b.ec.SubscribeToPersistentSubscriptionToAll(
		context.Background(), buildingMeasuredConsumerGroup, esdb.SubscribeToPersistentSubscriptionOptions{},
	)
	if err != nil {
		panic(err)
	}

	// TODO: Here may consume concurrently?
	for {
		esdbEvent := sub.Recv()
		if esdbEvent.EventAppeared != nil {
			var e event.BuildingMeasured
			if err := jsoniter.Unmarshal(esdbEvent.EventAppeared.Event.OriginalEvent().Data, &e); err != nil {
				panic(err)
			}

			b.handler.Handle(esdbEvent.EventAppeared.Event.OriginalEvent().StreamID, &e)
			sub.Ack(esdbEvent.EventAppeared.Event)
		}

		if esdbEvent.SubscriptionDropped != nil {
			break
		}
	}
}
