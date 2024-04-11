package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"log/slog"
	"secap-input/internal/common/esdb"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/application"
	"secap-input/internal/projection/building"
)

type FiberServer struct {
	*fiber.App

	*application.CreateBuildingCommandHandler
	*application.MeasureBuildingCommandHandler
}

func New() *FiberServer {

	esdbClient := esdb.ConnectESDB()

	aggregateRepository := repository.NewAggregateRepository(esdbClient)

	// CommandHandlers
	createBuildingCommandHandler := application.NewCreateBuildingCommandHandler(aggregateRepository)
	measureBuildingCommandHandler := application.NewMeasureBuildingCommandHandler(aggregateRepository)

	// Projections
	buildingProjection := building.NewBuildingProjection(esdbClient)
	go func() {
		if err := buildingProjection.Subscribe(context.Background()); err != nil {
			slog.Error("error on building projection", err)
			// Should there be a cancel?
		}
	}()

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "secap-input",
			AppName:      "secap-input",
			JSONEncoder:  jsoniter.Marshal,
			JSONDecoder:  jsoniter.Unmarshal,
		}),

		CreateBuildingCommandHandler:  createBuildingCommandHandler,
		MeasureBuildingCommandHandler: measureBuildingCommandHandler,
	}

	return server
}
