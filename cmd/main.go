package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lborres/los_logger/api"
	"github.com/lborres/los_logger/config"
	"github.com/lborres/los_logger/db"
	"github.com/lborres/los_logger/logger"
)

func run(_ context.Context, cfg config.Config) error {
	addr := fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.ServerPort)

	chExit := make(chan struct{})
	defer close(chExit)

	db, err := db.InitDB(*cfg.PGConfig)
	if err != nil {
		log.Fatal(err)
	}

	go logger.InitLogger(chExit, cfg.LogFileLoc)

	if err := api.StartAPIServer(addr, db); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("Starting LOS Logger")
	ctx := context.Background()
	if err := run(ctx, config.InitConfig()); err != nil {
		log.Fatal(err)
	}
}
