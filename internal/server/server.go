package server

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"secap-input/internal/common/esdb"
	"secap-input/internal/common/infrastructure/repository"
	"secap-input/internal/domain/building/application"
	"secap-input/internal/domain/building/infrastructure"
)

type FiberServer struct {
	*fiber.App

	*application.CreateBuildingCommandHandler
	*application.MeasureBuildingCommandHandler
}

func New() *FiberServer {

	esdbClient := esdb.ConnectESDB()

	aggregateRepository := repository.NewAggregateRepository(esdbClient)
	mtp := infrastructure.NewMeasurementTypeProvider()

	// CommandHandlers
	createBuildingCommandHandler := application.NewCreateBuildingCommandHandler(aggregateRepository)
	measureBuildingCommandHandler := application.NewMeasureBuildingCommandHandler(aggregateRepository, mtp)

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
