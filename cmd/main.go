package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lborres/los_logger/api"
	"github.com/lborres/los_logger/config"
	"github.com/lborres/los_logger/db"
)

func run(_ context.Context, cfg config.Config) error {

	addr := fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.ServerPort)
	log.Println(addr)

	db, err := db.InitDB(*cfg.PGConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.StartAPIServer(addr, db); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	log.Println("Starting HTTP Server")
	ctx := context.Background()
	if err := run(ctx, config.InitConfig()); err != nil {
		log.Fatal(err)
	}
}
