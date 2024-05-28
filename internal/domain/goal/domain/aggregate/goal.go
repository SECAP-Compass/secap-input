package aggregate

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

type Goal struct {
	*eventsourcing.AggregateBase
	target  *vo.Emission
	current *vo.Emission // Current can be reached from read api too but for audit purpose we can keep it here

	start time.Time
	end   time.Time
}

func NewGoalAggregate() *Goal {
	g := &Goal{}
	aggregateCfg := &eventsourcing.AggregateConfig{
		Type:         "Goal",
		EventHandler: g.EventHandler,
	}
	g.AggregateBase = eventsourcing.NewAggregateBase(aggregateCfg)

	return g
}

func NewGoalAggregateWithId(id string) *Goal {
	a := NewGoalAggregate()
	a.AggregateBase.ID = id

	return a
}

func (g *Goal) EventHandler(e *eventsourcing.Event) error {
	switch e.EventType {
	case "goal.created":
		return g.OnGoalCreatedEvent(e)
	case "goal.progressed":
		return g.OnGoalProgressedEvent(e)
	default:
		return nil
	}
}

func (g *Goal) OnGoalCreatedEvent(e *eventsourcing.Event) error {
	event := &event.GoalCreatedEvent{}

	err := e.GetEventData(event)
	if err != nil {
		return err
	}

	g.target = &event.Target
	g.start = event.Start
	g.end = event.End

	return nil
}

func (g *Goal) OnGoalProgressedEvent(e *eventsourcing.Event) error {
	event := &event.GoalProgressedEvent{}

	err := e.GetEventData(event)
	if err != nil {
		return err
	}

	return nil
}
