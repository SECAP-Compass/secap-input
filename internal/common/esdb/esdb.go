package esdb

import (
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gofiber/fiber/v2/log"
)

func ConnectESDB() *esdb.Client {
	cfg, err := esdb.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		log.Fatal(err)
	}

	client, err := esdb.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client.Config.Address: %v\n", client.Config.Address)

	return client
}
