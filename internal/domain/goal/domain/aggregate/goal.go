package aggregate

import (
	"secap-input/internal/common/eventsourcing"
	"secap-input/internal/domain/goal/domain/event"
	"secap-input/internal/domain/goal/domain/vo"
	"time"
)

type Goal struct {
	*eventsourcing.AggregateBase
	Target     *vo.Emission `json:"Target"`
	Current    *vo.Emission `json:"current"` // Current can be reached from read api too but for audit purpose we can keep it here
	Limit      float64      `json:"limit"`
	CityId     uint         `json:"cityId"`
	DistrictId uint         `json:"districtId"`
	Percentage *vo.Emission `json:"percentage"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
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
	case "goal.updated":
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

	g.Target = &ev.Target
	g.Start = ev.Start
	g.End = ev.End
	g.CityId = uint(ev.CityId)
	g.DistrictId = uint(ev.DistrictId)
	g.Current = &vo.Emission{}
	g.Percentage = &vo.Emission{}

	return nil
}

func (g *Goal) OnGoalUpdatedEvent(e *eventsourcing.Event) error {
	ev := &event.GoalUpdatedEvent{}

	err := e.GetEventData(ev)
	if err != nil {
		return err
	}

	g.Current.Add(ev.Emission)
	g.Percentage.Add(ev.Delta)

	return nil
}

func (g *Goal) GetCurrent() *vo.Emission {
	return g.Current
}

func (g *Goal) GetTarget() *vo.Emission {
	return g.Target
}
