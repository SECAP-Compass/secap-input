package esdb

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	"github.com/gofiber/fiber/v2/log"
)

type SubscriptionWorker func(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error

func ConnectESDB() *esdb.Client {
	cfg, err := esdb.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		log.Fatal(err)
	}

	client, err := esdb.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
