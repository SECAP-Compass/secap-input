package server

import (
	eventstore "github.com/EventStore/EventStore-Client-Go/v4/esdb"
	"secap-input/internal/common/esdb"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/application"
	"secap-input/internal/domain/building/core/ports"
	"secap-input/internal/domain/building/infrastructure"
	"secap-input/internal/domain/calculation/consumer"
	"secap-input/internal/domain/calculation/domain/port"
	"secap-input/internal/domain/calculation/domain/use_case"
	infrastructure2 "secap-input/internal/domain/calculation/infrastructure"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type FiberServer struct {
	*fiber.App
	esdbClient *eventstore.Client

	ports.MeasurementTypeProvider
	port.CalculationRepository
	port.BuildingMeasuredConsumer
	port.BuildingMeasuredHandler

	application.CreateBuildingCommandHandler
	application.MeasureBuildingCommandHandler
}

func New() *FiberServer {

	esdbClient := esdb.ConnectESDB()

	aggregateRepository := repository.NewAggregateRepository(esdbClient)
	eventRepository := repository.NewEventRepository(esdbClient)
	mtp := infrastructure.NewMeasurementTypeProvider()

	// Building
	// CommandHandlers
	createBuildingCommandHandler := application.NewCreateBuildingCommandHandler(aggregateRepository)
	measureBuildingCommandHandler := application.NewMeasureBuildingCommandHandler(aggregateRepository, mtp)

	// Calculation
	calculationRepository := infrastructure2.NewCalculationRepository(eventRepository)
	buildingMeasuredHandler := use_case.NewBuildingMeasuredHandler(calculationRepository)
	buildingMeasuredConsumer := consumer.NewBuildingMeasuredConsumer(esdbClient, buildingMeasuredHandler)

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			DisableKeepalive: true,
			ServerHeader:     "secap-input",
			AppName:          "secap-input",
			JSONEncoder:      jsoniter.Marshal,
			JSONDecoder:      jsoniter.Unmarshal,
		}),

		esdbClient: esdbClient,

		MeasurementTypeProvider: mtp,

		CreateBuildingCommandHandler:  createBuildingCommandHandler,
		MeasureBuildingCommandHandler: measureBuildingCommandHandler,

		BuildingMeasuredHandler:  buildingMeasuredHandler,
		BuildingMeasuredConsumer: buildingMeasuredConsumer,
	}

	return server
}
