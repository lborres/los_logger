package main

import (
	"context"
	"log"

	"github.com/lborres/los_logger/internal/config"
	"github.com/lborres/los_logger/internal/db"
	"github.com/lborres/los_logger/logger"
)

func run(_ context.Context, cfg config.Config) error {
	logWriter := logger.InitLogWriter(cfg.LogFileLoc)
	defer logWriter.CloseWriter()

	log.Println("Starting LOS Logger")

	db, err := db.InitDB(*cfg.PGConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger.StartLOSLogger(db)

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, config.InitConfig()); err != nil {
		log.Fatal(err)
	}
}
