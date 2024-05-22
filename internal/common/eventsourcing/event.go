package eventsourcing

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	jsoniter "github.com/json-iterator/go"
)

type Event struct {
	EventID       uuid.UUID
	EventType     string
	Data          []byte
	Timestamp     time.Time
	AggregateType string
	AggregateID   uuid.UUID
	Version       int64
	Metadata      []byte
}

type EventMetadata struct {
	Authority   string `json:"authority"`
	OccurredBy  string `json:"occurredBy"`
	OperationId string `json:"operationId"`
}

func NewEvent(a *AggregateBase, eventType string) *Event {
	return &Event{
		EventID:       uuid.New(),
		EventType:     eventType,
		AggregateID:   a.GetAggregateId(),
		AggregateType: a.GetType(),
		Version:       a.GetVersion(),
		Timestamp:     time.Now().UTC(),
	}
}

func NewEventMetadataFromContext(ctx context.Context) *EventMetadata {
	return &EventMetadata{
		Authority:   ctx.Value("authority").(string),
		OccurredBy:  ctx.Value("agent").(string),
		OperationId: ctx.Value("operationId").(string),
	}
}

// This may cause performance issues?
func NewEventFromRecordedEvent(re *esdb.RecordedEvent) *Event {
	aggrId, _ := uuid.Parse(re.StreamID)

	return &Event{
		EventID:     re.EventID,
		EventType:   re.EventType,
		Data:        re.Data,
		Timestamp:   re.CreatedDate,
		AggregateID: aggrId,
		Version:     int64(re.EventNumber),
		Metadata:    re.UserMetadata,
	}
}

func (e *Event) GetEventId() uuid.UUID {
	return e.EventID
}

func (e *Event) GetAggregateId() uuid.UUID {
	return e.AggregateID
}

func (e *Event) GetVersion() int64 {
	return e.Version
}

func (e *Event) SetAggregateType(_type string) {
	e.AggregateType = _type
}

func (e *Event) SetVersion(version int64) {
	e.Version = version
}

// May use generics here?
func (e *Event) GetEventData(toParse interface{}) error {
	err := jsoniter.Unmarshal(e.Data, toParse)
	fmt.Errorf("%v", err)

	return err
}

func (e *Event) SetEventData(data interface{}) error {
	dataByteArr, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	e.Data = dataByteArr
	return nil
}

func (e *Event) GetMetaData(toParse interface{}) error {
	return jsoniter.Unmarshal(e.Metadata, toParse)
}

func (e *Event) SetMetaData(eventMetadata *EventMetadata) error {
	metaDataByteArr, err := jsoniter.Marshal(eventMetadata)
	if err != nil {
		return err
	}

	e.Metadata = metaDataByteArr
	return nil
}

func (e *Event) ToEventData() esdb.EventData {
	return esdb.EventData{
		EventID:     e.EventID,
		EventType:   e.EventType,
		ContentType: esdb.ContentTypeJson,
		Data:        e.Data,
		Metadata:    e.Metadata,
	}
}
