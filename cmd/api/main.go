package main

import (
	"fmt"
	"os"
	"secap-input/internal/server"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()
	server.Use(pprof.New())

	server.RegisterFiberRoutes()
	server.RegisterBuildingRoutes()
	server.RegisterGoalRoutes()

	server.App.Get("/metrics", monitor.New())
	p := os.Getenv("PORT")
	if p == "" {
		p = "8001"
	}

	port, _ := strconv.Atoi(p)
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
