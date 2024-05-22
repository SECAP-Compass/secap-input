package port

import "secap-input/internal/domain/calculation/domain/event"

type BuildingMeasuredHandler interface {
	Handle(aggregateId string, event *event.BuildingMeasured)
}

type BuildingMeasuredConsumer interface {
	Consume()
}
