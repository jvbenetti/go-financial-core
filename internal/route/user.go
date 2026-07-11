package route

import (
	"github.com/jvbenetti/go-financial-core/internal/handler"
	"github.com/labstack/echo/v5"
)

func SetupUserRoutes(public *echo.Group, h *handler.UserHandler) {
	// Public Routes
	authGroup := public.Group("/auth")
	authGroup.POST("", h.Register)
}
