package repository

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"secap-input/internal/common/eventsourcing"
	"time"

	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
)

type EventRepository struct {
	db *esdb.Client
}

func NewEventRepository(db *esdb.Client) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) SaveEvents(ctx context.Context, streamId string, events []eventsourcing.Event) error {
	eventDataList := make([]esdb.EventData, 0, len(events))
	for _, e := range events {
		eventDataList = append(eventDataList, e.ToEventData())
	}

	t := time.Now()
	_, err := r.db.AppendToStream(
		ctx,
		streamId,
		esdb.AppendToStreamOptions{
			ExpectedRevision: esdb.Any{},
		},
		eventDataList...,
	)
	if err != nil {
		slog.Error("error appending to stream", err)
		return err
	}

	slog.Info("event saved in", slog.String("time", time.Since(t).String()))

	return nil
}

func (r *EventRepository) Save(ctx context.Context, streamId string, event eventsourcing.Event) error {
	_, err := r.db.AppendToStream(
		ctx, streamId, esdb.AppendToStreamOptions{
			ExpectedRevision: esdb.Any{},
		}, event.ToEventData())

	if err != nil {
		slog.Error("LAN", err)
	}

	return nil
}

func (r *EventRepository) LoadEvents(ctx context.Context, streamId string) ([]*eventsourcing.Event, error) {
	stream, err := r.db.ReadStream(ctx, streamId, esdb.ReadStreamOptions{}, readCount)
	if err != nil {
		slog.Error("error reading stream", err)
		return nil, err
	}
	defer stream.Close()

	events := make([]*eventsourcing.Event, 0, 100)
	for {
		re, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			slog.Error("error reading stream", err)
			return nil, err
		}

		events = append(events, eventsourcing.NewEventFromRecordedEvent(re.Event))
	}

	return events, nil
}

func (r *EventRepository) GetLatestEvent(ctx context.Context, streamId string) (*eventsourcing.Event, error) {
	stream, err := r.db.ReadStream(ctx, streamId, esdb.ReadStreamOptions{
		From:      esdb.End{},
		Direction: esdb.Backwards,
	}, 1)
	if err != nil {
		slog.Error("error reading stream", err)
		return nil, err
	}
	defer stream.Close()

	re, err := stream.Recv()
	if err != nil && !errors.Is(err, io.EOF) {
		slog.Error("error reading stream", err)
		return nil, err
	}

	return eventsourcing.NewEventFromRecordedEvent(re.Event), nil
}

func (r *EventRepository) GetFirstEvent(ctx context.Context, streamId string) (*eventsourcing.Event, error) {
	stream, err := r.db.ReadStream(ctx, streamId, esdb.ReadStreamOptions{
		From:      esdb.Start{},
		Direction: esdb.Forwards,
	}, 1)
	if err != nil {
		slog.Error("error reading stream", err)
		return nil, err
	}
	defer stream.Close()

	re, err := stream.Recv()
	if err != nil && !errors.Is(err, io.EOF) {
		slog.Error("error reading stream", err)
	}

	return eventsourcing.NewEventFromRecordedEvent(re.Event), nil
}
