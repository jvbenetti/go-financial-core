package handler

import (
	"net/http"

	"github.com/jvbenetti/go-financial-core/internal/dto/request"
	"github.com/jvbenetti/go-financial-core/internal/service"
	"github.com/labstack/echo/v5"
)

type UserHandler struct { // Struct for receiver
	UserService *service.UserService
}

func (h *UserHandler) Register(c *echo.Context) error {
	req := new(request.UserRequest)

	// 1. Validate JSON
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Data!"})
	}

	resp, err := h.UserService.CreateUserWithAccount(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}
