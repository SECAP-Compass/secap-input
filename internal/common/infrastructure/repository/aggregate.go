package repository

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"math"
	"secap-input/internal/common/eventsourcing"
	"time"

	"github.com/google/uuid"

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
	stream, err := r.db.ReadStream(ctx, a.GetAggregateId(), esdb.ReadStreamOptions{}, readCount)
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

func (r *AggregateRepository) Create(ctx context.Context, a eventsourcing.AggregateRoot) error {
	er := esdb.StreamExists{}
	return r.save(ctx, a, er)
}

func (r *AggregateRepository) Save(ctx context.Context, a eventsourcing.AggregateRoot) error {
	var er esdb.ExpectedRevision
	if a.GetVersion() == 0 {
		er = esdb.NoStream{}
	} else {
		ev := uint64(a.GetVersion()) - uint64(len(a.GetUncommittedEvents()))
		er = esdb.Revision(ev)
	}

	return r.save(ctx, a, er)
}

func (r *AggregateRepository) SaveWithoutVersionCheck(ctx context.Context, a eventsourcing.AggregateRoot) error {
	er := esdb.Any{}
	return r.save(ctx, a, er)
}

func (r *AggregateRepository) Exists(ctx context.Context, aggregateId uuid.UUID) bool {
	readStreamOptions := esdb.ReadStreamOptions{
		Direction: esdb.Backwards,
		From:      esdb.Revision(1),
	}
	// error nil means stream exists?
	stream, err := r.db.ReadStream(ctx, aggregateId.String(), readStreamOptions, 1)
	stream.Close()

	return err != nil
}

func (r *AggregateRepository) save(ctx context.Context, a eventsourcing.AggregateRoot, er esdb.ExpectedRevision) error {
	if len(a.GetUncommittedEvents()) == 0 {
		return nil
	}

	eventDataList := make([]esdb.EventData, 0, len(a.GetUncommittedEvents()))
	for _, e := range a.GetUncommittedEvents() {
		eventDataList = append(eventDataList, e.ToEventData())
	}
	deadline := 250 * time.Second
	t := time.Now()
	_, err := r.db.AppendToStream(
		ctx,
		a.GetAggregateId(),
		esdb.AppendToStreamOptions{ExpectedRevision: er, Deadline: &deadline},
		eventDataList...,
	)
	if err != nil {
		slog.Error("error appending to stream", err)
		return err
	}

	slog.Info("eventSaved In", slog.String("time", time.Since(t).String()))
	a.ClearUncommittedEvents()
	return nil
}
