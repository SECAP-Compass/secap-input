package building

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log/slog"
)

type BuildingProjection struct {
	esdb *esdb.Client
}

const (
	BuildingPosgresProjectionGroupName = "building_postgres_projection_start"
	BuildingKafkaProjectionGroupName   = "building_kafka_projection"

	WorkerPoolSize = 1
)

func NewBuildingProjection(esdb *esdb.Client) *BuildingProjection {
	return &BuildingProjection{esdb: esdb}
}

func (b *BuildingProjection) Subscribe(ctx context.Context) error {

	// Create if not exists

	err := b.esdb.CreatePersistentSubscriptionToAll(ctx, BuildingPosgresProjectionGroupName, esdb.PersistentAllSubscriptionOptions{
		Filter: &esdb.SubscriptionFilter{Type: esdb.EventFilterType, Prefixes: []string{"building"}}},
	)

	if err != nil {
		var esdbErr *esdb.Error
		errors.As(err, &esdbErr)
		if esdbErr != nil && esdbErr.IsErrorCode(esdb.ErrorCodeResourceAlreadyExists) {
			slog.Error("Subscription already exists", err)
		} else {
			slog.Error("Error creating subscription", err)
			return err
		}
	}

	stream, err := b.esdb.SubscribeToPersistentSubscriptionToAll(ctx, BuildingPosgresProjectionGroupName, esdb.SubscribeToPersistentSubscriptionOptions{})
	if err != nil {
		return err
	}
	defer stream.Close()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= WorkerPoolSize; i++ {
		g.Go(func() error {
			return b.worker(ctx, stream, i)
		})
	}
	return g.Wait()
}

func (b *BuildingProjection) worker(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			e := stream.Recv()

			if e.SubscriptionDropped != nil {
				slog.Info("SubscriptionDropped")
				return errors.Errorf("SubscriptionDropped: %v", e.SubscriptionDropped)
			}

			if e.EventAppeared != nil {
				slog.Info("EventAppeared", slog.Int("WorkerID", workerID))
				slog.Info("", slog.String("Type", e.EventAppeared.Event.Event.EventType),
					slog.String("Data", string(e.EventAppeared.Event.Event.Data)),
					slog.String("AggregateId", e.EventAppeared.Event.Event.StreamID),
				)
			}

			if e.CheckPointReached != nil {
				slog.Info("CheckPointReached",
					slog.Uint64("Checkpoint", e.CheckPointReached.Commit),
					slog.Uint64("Checkpoint", e.CheckPointReached.Prepare),
				)
			}
		}
	}
}
