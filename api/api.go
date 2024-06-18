package api

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lborres/los_logger/service/dashboard"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func StartAPIServer(addr string, db *sql.DB) error {
	apiserver := &APIServer{
		addr: addr,
		db:   db,
	}

	if err := apiserver.initRoutes(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (server *APIServer) initRoutes() error {
	e := echo.New()

	log.Println("Initializing Server Routes")

	dashboardStorage := dashboard.NewStorage(server.db)
	dashboardHandler := dashboard.NewHandler(dashboardStorage)
	dashboardHandler.RegisterRoutes(e)

	log.Printf("HTTP Server listening at %s\n", server.addr)

	return e.Start(server.addr)
}
