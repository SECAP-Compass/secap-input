package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"os"
	"secap-input/internal/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()
	server.Use(pprof.New())

	server.RegisterFiberRoutes()
	server.RegisterBuildingRoutes()
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
