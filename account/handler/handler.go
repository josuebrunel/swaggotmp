package handler

import (
	"context"
	"ekolo/account/service"
	"ekolo/pkg/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(store storage.Storer) *UserHandler {
	return &UserHandler{
		svc: service.NewUserService(store),
	}
}

func (h *UserHandler) GetUserTypes(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := h.svc.GetTypes(ctx)
		return c.JSON(http.StatusOK, resp)
	}
}
