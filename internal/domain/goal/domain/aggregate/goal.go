package aggregate

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

type Goal struct {
	*eventsourcing.AggregateBase
	target     *vo.Emission
	current    *vo.Emission // Current can be reached from read api too but for audit purpose we can keep it here
	limit      float64
	cityId     uint
	districtId uint
	percentage *vo.Emission

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
	case "building.measurement.calculated":
		return g.OnGoalUpdatedEvent(e)
	default:
		return nil
	}
}

func (g *Goal) OnGoalCreatedEvent(e *eventsourcing.Event) error {
	ev := &event.GoalCreatedEvent{}

	err := e.GetEventData(ev)
	if err != nil {
		return err
	}

	g.target = &ev.Target
	g.start = ev.Start
	g.end = ev.End
	g.cityId = uint(ev.CityId)
	g.districtId = uint(ev.DistrictId)
	g.current = &vo.Emission{}

	return nil
}

func (g *Goal) OnGoalUpdatedEvent(e *eventsourcing.Event) error {
	ev := &event.GoalUpdatedEvent{}

	err := e.GetEventData(ev)
	if err != nil {
		return err
	}

	g.current.Add(ev.Emission)
	g.percentage.Add(ev.Delta)

	return nil
}

func (g *Goal) GetCurrent() *vo.Emission {
	return g.current
}

func (g *Goal) GetTarget() *vo.Emission {
	return g.target
}
