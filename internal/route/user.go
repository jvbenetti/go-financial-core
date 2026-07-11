package route

import (
	"github.com/jvbenetti/go-financial-core/internal/handler"
	"github.com/labstack/echo/v5"
)

func SetupUserGroup(public *echo.Echo, h handler.UserHandler) {
	// Public Routes
	authGroup := public.Group("/auth")
	authGroup.POST("", h.Register)
}
