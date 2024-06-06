package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e.Start(server.addr)
}
