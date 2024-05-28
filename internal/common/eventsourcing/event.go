package eventsourcing

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	jsoniter "github.com/json-iterator/go"
)

type Event struct {
	EventID       string
	EventType     string
	Data          []byte
	Timestamp     time.Time
	AggregateType string
	AggregateID   string
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
		EventID:       uuid.NewString(),
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
	return &Event{
		EventID:     re.EventID.String(),
		EventType:   re.EventType,
		Data:        re.Data,
		Timestamp:   re.CreatedDate,
		AggregateID: re.StreamID,
		Version:     int64(re.EventNumber),
		Metadata:    re.UserMetadata,
	}
}

func (e *Event) GetEventId() string {
	return e.EventID
}

func (e *Event) GetAggregateId() string {
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
		EventID:     uuid.MustParse(e.EventID),
		EventType:   e.EventType,
		ContentType: esdb.ContentTypeJson,
		Data:        e.Data,
		Metadata:    e.Metadata,
	}
}
