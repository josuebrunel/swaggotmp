package generic

import (
	"context"
	"ekolo/pkg/xlog"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	OpCreate = "create"
	OpGet    = "get"
	OpList   = "list"
	OpUpdate = "update"
	OpDelete = "delete"
)

// IRequest represents the interface for request objects.
type IRequest interface {
	GetID() string // Get the ID from the request.
}

// IResponse represents the interface for response objects.
type IResponse interface {
	GetStatusCode() int // Get the HTTP status code for the response.
}

// Response is the response object for the service
type Response struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
	Data   any      `json:"data"`
}

func (r Response) GetStatusCode() int { return r.Status }

// NewResponse returns a new response object
func NewResponse(status int, errors []string, data any) Response {
	return Response{
		Status: status,
		Errors: errors,
		Data:   data,
	}
}

// Service is an interface representing a generic service with CRUD operations.
type IService interface {
	GetName() string                                                   // Get the service name.
	GetPathParams() []string                                           // Get path parameters
	GetRequest(string) IRequest                                        // Get an instance of the request object.
	Create(context.Context, IRequest) (IResponse, error)               // Create a resource.
	Get(context.Context, IRequest) (IResponse, error)                  // Get a resource.
	List(context.Context, IRequest, map[string]any) (IResponse, error) // Get list of resource
	Update(context.Context, IRequest) (IResponse, error)               // Update a resource.
	Delete(context.Context, IRequest) error                            // Delete a resource.
}

// GenericServiceHandler is a handler for generic service operations.
type GenericServiceHandler struct {
	svc IService
	e   *echo.Echo
}

// Create is a handler for the create operation.
func (s GenericServiceHandler) Create(context context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			err error
			req = s.svc.GetRequest(OpCreate)
		)
		// Try to bind payload.
		if err = ctx.Bind(req); err != nil {
			xlog.Error("create-bind-error", "err", err)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		// Let the target service process the request.
		resp, err := s.svc.Create(context, req)
		if err != nil {
			return ctx.JSON(resp.GetStatusCode(), resp)
		}
		return ctx.JSON(resp.GetStatusCode(), resp)
	}
}

// Get is a handler for the get operation.
func (s GenericServiceHandler) Get(ctx context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := s.svc.GetRequest(OpGet)
		if err := ctx.Bind(req); err != nil {
			xlog.Error("get-bind-error", "err", err)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		resp, err := s.svc.Get(ctx.Request().Context(), req)
		if err != nil {
			return ctx.JSON(resp.GetStatusCode(), resp)
		}
		return ctx.JSON(resp.GetStatusCode(), resp)
	}
}

// List is an handler for the list operation
func (s GenericServiceHandler) List(ctx context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := s.svc.GetRequest(OpList)
		if err := ctx.Bind(req); err != nil {
			xlog.Error("list-bind-error", "err", err)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		filter := map[string]any{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &filter); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		resp, err := s.svc.List(ctx.Request().Context(), req, filter)
		if err != nil {
			return ctx.JSON(resp.GetStatusCode(), resp)
		}

		return ctx.JSON(resp.GetStatusCode(), resp)
	}
}

// Update is a handler for the update operation.
func (s GenericServiceHandler) Update(ctx context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		req := s.svc.GetRequest(OpUpdate)
		// Try to bind payload.
		if err = ctx.Bind(req); err != nil {
			xlog.Error("updated-bind-error", "err", err)
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		resp, err := s.svc.Update(ctx.Request().Context(), req)
		if err != nil {
			return ctx.JSON(resp.GetStatusCode(), resp)
		}
		return ctx.JSON(resp.GetStatusCode(), resp)
	}
}

// Delete is a handler for the delete operation.
func (s GenericServiceHandler) Delete(ctx context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := s.svc.GetRequest(OpDelete)
		if err := ctx.Bind(req); err != nil {
			xlog.Error("delete-bind-error", "err", err)
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		if err := s.svc.Delete(ctx.Request().Context(), req); err != nil {
			return ctx.JSON(http.StatusInternalServerError, nil)
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}

// GetPathParamName returns the path parameter name used for routing.
func (s GenericServiceHandler) GetPathParamName() string {
	params := []string{""}
	params = append(params, s.svc.GetPathParams()...)
	return strings.Join(params, "/:")
}

// MountService creates and mounts a GenericServiceHandler for the provided service on the given Echo instance.
func MountService(e *echo.Echo, svc IService) {
	ctx := context.Background()
	h := GenericServiceHandler{svc: svc, e: e}
	g := h.e.Group(svc.GetName())
	paramPath := h.GetPathParamName()
	param := svc.GetPathParams()[0]
	g.POST("", h.Create(ctx)).Name = fmt.Sprintf("%s-create", param)
	g.GET("", h.List(ctx)).Name = fmt.Sprintf("%s-list", param)
	g.GET(paramPath, h.Get(ctx)).Name = fmt.Sprintf("%s-get", param)
	g.PATCH(paramPath, h.Update(ctx)).Name = fmt.Sprintf("%s-update", param)
	g.DELETE(paramPath, h.Delete(ctx)).Name = fmt.Sprintf("%s-delete", param)
}
