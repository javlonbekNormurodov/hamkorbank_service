package main

import (
	"context"
	"hamkorbank/api"
	"hamkorbank/api/handlers"
	"hamkorbank/config"
	"hamkorbank/pkg/logger"
	"hamkorbank/storage/postgres"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.Debug:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.Test:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	log := logger.NewLogger("hamkorbank_servise", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, pgStore)

	api.SetUpAPI(r, h, cfg)

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
