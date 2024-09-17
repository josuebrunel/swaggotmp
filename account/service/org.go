package service

import (
	"context"
	"ekolo/account/model"
	generic "ekolo/pkg/echogeneric"
	"ekolo/pkg/storage"
	"errors"

	"github.com/google/uuid"
)

// GetModels returns the models used by the service
func GetModels() []any {
	return []any{
		model.Organization{},
		model.User{},
	}
}

// Service is the service object
type Service struct {
	repo storage.Storer
}

func (s Service) GetName() string {
	return "organization"
}

// GetPathParams returns service' path params
func (s Service) GetPathParams() []string {
	return []string{"org"}
}

// GetRequest returns the request object for the service
func (s Service) GetRequest(name string) generic.IRequest {
	switch name {
	case "create":
		return &RequestOrgCreate{}
	case "get":
		return &RequestOrgGet{}
	case "list":
		return &RequestOrgList{}
	case "update":
		return &RequestOrgUpdate{}
	case "delete":
		return &RequestOrgDelete{}
	default:
		return Request{}
	}
}

// New returns a new service
func New(repo storage.Storer) *Service {
	return &Service{
		repo: repo,
	}
}

// Request is the request object for the service
type Request struct{}

func (r Request) GetID() string {
	return "org"
}

// RequestOrgCreate is the request object for the create method
type RequestOrgCreate struct {
	Request
	model.Organization
}

// RequestOrgGet is the request object for the get method
type RequestOrgGet struct {
	Request
	OrgParam string `param:"org"`
}

// RequestOrgList is the request object for the list method
type RequestOrgList struct {
	Request
}

// RequestOrgUpdate is the request object for the update method
type RequestOrgUpdate struct {
	Request
	OrgParam string `param:"org"`
	model.Organization
}

// RequestOrgDelete is the request object for the delete method
type RequestOrgDelete struct {
	Request
	OrgParam string `param:"org"`
}

// Response is the response object for the service
type Response struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
	Data   any      `json:"data"`
}

// NewResponse returns a new response object
func NewResponse(status int, errors []string, data any) Response {
	return Response{
		Status: status,
		Errors: errors,
		Data:   data,
	}
}

// GetStatusCode returns the status code of the response
func (r Response) GetStatusCode() int {
	return r.Status
}

// Create creates a new organization
// @Summary Create an organization
// @Description Create an organization
// @ID org-create
// @Tags organization
// @Accept json
// @Produce json
// @Param organization body RequestOrgCreate true "Organization data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization [post]
func (s Service) Create(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestOrgCreate)
	_, err := s.repo.Create(&r.Organization)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, r.Organization), err
}

// Get gets an organization
// @Summary Get an organization
// @Description Get an organization
// @ID org-get
// @Tags organization
// @Produce json
// @Param uuid path string true "Organization ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{uuid} [get]
func (s Service) Get(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	var (
		r      = req.(*RequestOrgGet)
		org    model.Organization
		filter = map[string]any{
			"uuid": r.OrgParam,
		}
	)
	_, err := s.repo.Get(&org, filter)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return NewResponse(404, []string{err.Error()}, nil), err
		}
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, org), err

}

// List lists organizations
// @Summary List organizations
// @Description List organizations
// @ID orgs-get
// @Tags organization
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization [get]
func (s Service) List(ctx context.Context, req generic.IRequest, filter map[string]any) (generic.IResponse, error) {
	var (
		_    = req.(*RequestOrgList)
		orgs []model.Organization
	)
	_, err := s.repo.List(&orgs, filter)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, orgs), nil
}

// Update updates an organization
// @Summary Update an organization
// @Description Update an organization
// @ID org-update
// @Tags organization
// @Accept json
// @Produce json
// @Param uuid path string true "Organization ID"
// @Param organization body RequestOrgUpdate true "Organization data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{uuid} [patch]
func (s Service) Update(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestOrgUpdate)
	r.UUID = uuid.MustParse(r.OrgParam)
	_, err := s.repo.Update(r.Organization)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, r.Organization), nil
}

// Delete deletes an organization
// @Summary Delete an organization
// @Description Delete an organization
// @ID org-delete
// @Tags organization
// @Param uuid path string true "Organization ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{uuid} [delete]
func (s Service) Delete(ctx context.Context, req generic.IRequest) error {
	var (
		org    model.Organization
		filter = map[string]any{
			"uuid": req.(*RequestOrgDelete).OrgParam,
		}
	)
	_, err := s.repo.Delete(&org, filter)
	if err != nil {
		return err
	}
	return nil
}

// Service is the service interface
var _ generic.IService = new(Service)
