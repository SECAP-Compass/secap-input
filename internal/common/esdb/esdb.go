package esdb

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v4/esdb"
	"github.com/gofiber/fiber/v2/log"
	"log/slog"
	"os"
	"os/exec"
)

type SubscriptionWorker func(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error

func ConnectESDB() *esdb.Client {

	out, _ := exec.Command("ls").Output()
	slog.Info(string(out))
	out, _ = exec.Command("pwd").Output()
	slog.Info(string(out))

	connectionString := os.Getenv("EVENTSTORE_CONNECTION_STRING")
	slog.Info("", slog.String("Eventstoredb connection string", connectionString))

	cfg, err := esdb.ParseConnectionString(connectionString + "?tls=false")
	if err != nil {
		log.Fatal(err)
	}

	client, err := esdb.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
