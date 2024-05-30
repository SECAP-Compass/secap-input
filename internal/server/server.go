package server

import (
	"secap-input/internal/common/esdb"
	"secap-input/internal/common/infrastructure/repository"
	building_application "secap-input/internal/domain/building/application"
	"secap-input/internal/domain/building/core/ports"
	"secap-input/internal/domain/building/infrastructure"
	"secap-input/internal/domain/calculation/consumer"
	building_port "secap-input/internal/domain/calculation/domain/port"
	"secap-input/internal/domain/calculation/domain/use_case"
	infrastructure2 "secap-input/internal/domain/calculation/infrastructure"
	goal_application "secap-input/internal/domain/goal/application"
	goal_port "secap-input/internal/domain/goal/domain/port"

	eventstore "github.com/EventStore/EventStore-Client-Go/v4/esdb"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type FiberServer struct {
	*fiber.App
	esdbClient *eventstore.Client

	ports.MeasurementTypeProvider
	building_port.CalculationRepository
	building_port.BuildingMeasuredConsumer
	building_port.BuildingMeasuredHandler

	building_application.CreateBuildingCommandHandler
	building_application.MeasureBuildingCommandHandler

	goal_port.GoalCreator
}

func New() *FiberServer {

	esdbClient := esdb.ConnectESDB()

	aggregateRepository := repository.NewAggregateRepository(esdbClient)
	eventRepository := repository.NewEventRepository(esdbClient)
	mtp := infrastructure.NewMeasurementTypeProvider()

	// Building
	// CommandHandlers
	createBuildingCommandHandler := building_application.NewCreateBuildingCommandHandler(aggregateRepository)
	measureBuildingCommandHandler := building_application.NewMeasureBuildingCommandHandler(aggregateRepository, mtp)

	// Calculation
	calculationRepository := infrastructure2.NewCalculationRepository(eventRepository)
	buildingMeasuredHandler := use_case.NewBuildingMeasuredHandler(calculationRepository)
	buildingMeasuredConsumer := consumer.NewBuildingMeasuredConsumer(esdbClient, buildingMeasuredHandler)

	// Goal
	goalCreator := goal_application.NewGoalCreatorAdapter(aggregateRepository)

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

		GoalCreator: goalCreator,
	}

	return server
}
