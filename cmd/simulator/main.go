package main

import (
	"fmt"
	"net/http"

	"github.com/acergiyen/simulator/infra/postgresql"
	"github.com/acergiyen/simulator/internal/app/config"
	"github.com/acergiyen/simulator/internal/app/handler"
	"github.com/acergiyen/simulator/internal/app/logger"
	"github.com/acergiyen/simulator/internal/app/repositories"
	"github.com/acergiyen/simulator/internal/app/services"
)

func main() {
	// Get application configuration
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Initialize logger
	logger := logger.AppLogger(config)
	logger.Printf("%v application started", config.App.Name)

	// Initialize database client
	db, err := postgresql.NewClient(config)
	if err != nil {
		logger.Fatalf("Client error:%v", err)
	}

	// Initialize repositories
	teamRepository := repositories.NewTeamRepository(db)
	matchRepoistory := repositories.NewMatchRepository(db)
	scoreRepository := repositories.NewScoreRepository(db)
	assistRepository := repositories.NewAssistRepository(db)

	// Initialize services
	teamService := services.NewTeamService(logger, teamRepository)
	matchService := services.NewMatchService(logger, matchRepoistory)
	scoreService := services.NewScoreService(logger, scoreRepository)
	assistService := services.NewAssistService(logger, assistRepository)

	// Initialize handler for simulator
	simulatorHandler := handler.NewSimulatorHandler(logger, config, teamService, matchService, scoreService, assistService)

	// Define HTTP endpoint for simulation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read and process HTML content
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		simulatorHandler.Simulate(w)
	})

	// Start HTTP server
	http.ListenAndServe(fmt.Sprintf(":%v", config.App.Port), nil)
}
