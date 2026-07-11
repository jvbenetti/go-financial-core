package route

import (
	"github.com/jvbenetti/go-financial-core/internal/handler"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	api := e.Group("/api/v1")

	publicGroup := api.Group("")

	SetupUserRoutes(publicGroup, userHandler)
}
