package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"math"
	"secap-input/internal/common/eventsourcing"

	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
)

const readCount = math.MaxUint64

type AggregateRepository struct {
	db *esdb.Client
}

type metadata struct {
	Authority string `json:"authority"`
}

func NewAggregateRepository(db *esdb.Client) *AggregateRepository {
	return &AggregateRepository{db: db}
}

func (r *AggregateRepository) Load(ctx context.Context, a eventsourcing.AggregateRoot) error {
	stream, err := r.db.ReadStream(ctx, a.GetAggregateId().String(), esdb.ReadStreamOptions{}, readCount)
	if err != nil {
		slog.Error("error reading stream", err)
		return err
	}
	defer stream.Close()

	for {
		re, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		var esdbErr *esdb.Error
		errors.As(err, &esdbErr)

		if esdbErr != nil && esdbErr.IsErrorCode(esdb.ErrorCodeResourceNotFound) {
			slog.Error("stream not found", err)
			return err
		}
		//if errors.Is(err, esdb.ErrorCodeResourceNotFound) {
		//	slog.Error("stream not found", err)
		//	return err
		//}

		e := eventsourcing.NewEventFromRecordedEvent(re.Event)

		if err := a.RaiseEvent(e); err != nil {
			slog.Error("error raising event", err)
			return err
		}
	}

	return nil
}

func (r *AggregateRepository) Save(ctx context.Context, a eventsourcing.AggregateRoot) error {
	if len(a.GetUncommittedEvents()) == 0 {
		return nil
	}

	eventDataList := make([]esdb.EventData, 0, len(a.GetUncommittedEvents()))
	for _, e := range a.GetUncommittedEvents() {
		eventDataList = append(eventDataList, e.ToEventData())
	}

	var er esdb.ExpectedRevision
	if a.GetVersion() == 0 {
		er = esdb.NoStream{}
	} else {
		ev := uint64(a.GetVersion()) - uint64(len(a.GetUncommittedEvents()))
		er = esdb.Revision(ev)
	}

	_, err := r.db.AppendToStream(
		ctx,
		a.GetAggregateId().String(),
		esdb.AppendToStreamOptions{ExpectedRevision: er},
		eventDataList...,
	)
	if err != nil {
		slog.Error("error appending to stream", err)
		return err
	}

	a.ClearUncommittedEvents()
	return nil
}

func (r *AggregateRepository) Exists(ctx context.Context, aggregateId uuid.UUID) bool {
	readStreamOptions := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.Revision(1)}
	// error nil means stream exists?
	stream, err := r.db.ReadStream(ctx, aggregateId.String(), readStreamOptions, 1)
	defer stream.Close()

	if err == nil {
		return false
	}
	return true
}
