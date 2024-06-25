package dashboard

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/lborres/los_logger/types"
	"github.com/lborres/los_logger/views/dashboard_views"
)

type Handler struct {
	storage types.DashboardStorage
}

func NewHandler(storage types.DashboardStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Static("/static", "static")
	e.GET("/", h.handleHome)
}

func (h *Handler) handleHome(c echo.Context) error {
	return renderView(c, dashboard_views.HomeIndex("LOS Logger | Home", dashboard_views.Home()))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
