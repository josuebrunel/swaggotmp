package service

import (
	"context"
	"ekolo/account/model"
	generic "ekolo/pkg/echogeneric"
	"ekolo/pkg/storage"
	"ekolo/pkg/xlog"
	"errors"

	"github.com/google/uuid"
)

const (
	TypeMANAGER = "MANAGER"
	TypeTEACHER = "TEACHER"
	TypeSTUDENT = "STUDENT"
)

// getUserTypes returns types of users
func getUserTypes() []string {
	return []string{TypeMANAGER, TypeTEACHER, TypeSTUDENT}
}

// UserService is the service object
type UserService struct {
	repo storage.Storer
}

func (s UserService) GetName() string {
	return "organization/:org/user"
}

// GetPathParams returns service' path params
func (s UserService) GetPathParams() []string {
	return []string{"user"}
}

// GetRequest returns the request object for the service
func (s UserService) GetRequest(name string) generic.IRequest {
	switch name {
	case "create":
		return &RequestUserCreate{}
	case "get":
		return &RequestUserGet{}
	case "list":
		return &RequestUserList{}
	case "update":
		return &RequestUserUpdate{}
	case "delete":
		return &RequestUserDelete{}
	default:
		return Request{}
	}
}

// New returns a new service
func NewUserService(repo storage.Storer) *UserService {
	return &UserService{
		repo: repo,
	}
}

// RequestUser is the request object for the service
type RequestUser struct{}

func (r RequestUser) GetID() string {
	return "user"
}

// RequestUserCreate is the request object for the create method
type RequestUserCreate struct {
	RequestUser
	model.User
}

// RequestUserGet is the request object for the get method
type RequestUserGet struct {
	RequestUser
	UserParam string `param:"user"`
	OrgParam  string `param:"org"`
}

// RequestUserList is the request object for the list method
type RequestUserList struct {
	RequestUser
	OrgParam string `param:"org"`
}

// RequestUserUpdate is the request object for the update method
type RequestUserUpdate struct {
	RequestUser
	UserParam string `param:"user"`
	OrgParam  string `param:"org"`
	model.User
}

// RequestUserDelete is the request object for the delete method
type RequestUserDelete struct {
	RequestUser
	UserParam string `param:"user"`
	OrgParam  string `param:"org"`
}

// Create creates a new user
// @Summary Create an user
// @Description Create an user
// @ID user-create
// @Tags user
// @Accept json
// @Produce json
// @Param org path string true "Organization ID"
// @Param user body RequestUserCreate true "user data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/user [post]
func (s UserService) Create(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestUserCreate)
	_, err := s.repo.Create(&r.User)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, r.User), err
}

// Get gets an user
// @Summary Get an user
// @Description Get an user
// @ID user-get
// @Tags user
// @Produce json
// @Param org path string true "organization ID"
// @Param uuid path string true "user ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/user/{uuid} [get]
func (s UserService) Get(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	var (
		r      = req.(*RequestUserGet)
		org    model.User
		filter = map[string]any{
			"uuid":     r.UserParam,
			"orgUUUID": r.OrgParam,
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

// List lists users
// @Summary List user users
// @Description List user users
// @ID users-get
// @Tags user
// @Produce json
// @Param org path string true "organization ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/user [get]
func (s UserService) List(ctx context.Context, req generic.IRequest, filter map[string]any) (generic.IResponse, error) {
	var (
		r  = req.(*RequestUserList)
		uu []model.User
	)
	filter["org_uuid"] = r.OrgParam
	xlog.Debug("params", "filter", filter, "req", r)
	_, err := s.repo.List(&uu, filter)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, uu), nil
}

// Update updates an user
// @Summary Update an organization user
// @Description Update an organization user
// @ID user-update
// @Tags user
// @Accept json
// @Produce json
// @Param org path string true "organization ID"
// @Param uuid path string true "user ID"
// @Param user body RequestUserUpdate true "user data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/user/{uuid} [patch]
func (s UserService) Update(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestUserUpdate)
	r.UUID = uuid.MustParse(r.UserParam)
	r.OrgUUID = uuid.MustParse(r.OrgParam)
	_, err := s.repo.Update(&r.User)
	if err != nil {
		return NewResponse(500, []string{err.Error()}, nil), err
	}
	return NewResponse(200, nil, r.User), nil
}

// Delete deletes an user
// @Summary Delete organization user
// @Description Delete organization user
// @ID user-delete
// @Tags user
// @Param org path string true "organization ID"
// @Param uuid path string true "user ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/user/{uuid} [delete]
func (s UserService) Delete(ctx context.Context, req generic.IRequest) error {
	var (
		org    model.User
		filter = map[string]any{
			"uuid": req.(*RequestUserDelete).UserParam,
		}
	)
	_, err := s.repo.Delete(&org, filter)
	if err != nil {
		return err
	}
	return nil
}

// Types returns users' types
// @Summary List users's type
// @Description List users' type
// @ID user-types
// @Tags user
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /user/types [get]
func (s UserService) GetTypes(ctx context.Context) generic.IResponse {
	return NewResponse(200, nil, getUserTypes())
}

// UserService is the service interface
var _ generic.IService = new(UserService)
