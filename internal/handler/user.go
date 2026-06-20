package handler

import (
	"github.com/jvbenetti/go-financial-core/internal/service"
	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) Register(c echo.Context) error {
	
}
