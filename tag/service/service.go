package service

import (
	"context"
	generic "ekolo/pkg/echogeneric"
	"ekolo/pkg/storage"
	"ekolo/tag/model"
	"errors"

	"github.com/google/uuid"
)

// GetModels returns service models
func GetModels() []any {
	return []interface{}{
		model.Tag{},
	}
}

// Tag is the service object
type Tag struct {
	repo storage.Storer
}

func (s Tag) GetName() string {
	return "organization/:org/tag"
}

// GetPathParams returns service' path params
func (s Tag) GetPathParams() []string {
	return []string{"tag"}
}

// GetRequest returns the request object for the service
func (s Tag) GetRequest(name string) generic.IRequest {
	switch name {
	case "create":
		return &RequestTagCreate{}
	case "get":
		return &RequestTagGet{}
	case "list":
		return &RequestTagList{}
	case "update":
		return &RequestTagUpdate{}
	case "delete":
		return &RequestTagDelete{}
	default:
		return RequestTag{}
	}
}

// New returns a new service
func New(repo storage.Storer) *Tag {
	return &Tag{
		repo: repo,
	}
}

// RequestTag is the request object for the service
type RequestTag struct{}

func (r RequestTag) GetID() string {
	return "tag"
}

// PayloadTagCreate is the struct representing the create request payload
type PayloadTag struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

// RequestTagCreate is the request object for the create method
type RequestTagCreate struct {
	RequestTag
	PayloadTag
	OrgParam string `param:"org"`
}

// RequestTagGet is the request object for the get method
type RequestTagGet struct {
	RequestTag
	OrgParam string `param:"org"`
	TagParam string `param:"tag"`
}

// RequestTagList is the request object for the list method
type RequestTagList struct {
	RequestTag
	OrgParam uuid.UUID `param:"org"`
}

// RequestTagUpdate is the request object for the update method
type RequestTagUpdate struct {
	RequestTag
	TagParam string `param:"tag"`
	OrgParam string `param:"org"`
	PayloadTag
}

// RequestTagDelete is the request object for the delete method
type RequestTagDelete struct {
	RequestTag
	TagParam string `param:"tag"`
	OrgParam string `param:"org"`
}

// Create creates a new tag
// @Summary Create an tag
// @Description Create an tag
// @ID tag-create
// @Tags tag
// @Accept json
// @Produce json
// @Param org path string true "Organization ID" Format(uuid)
// @Param tag body PayloadTag true "tag data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/tag [post]
func (s Tag) Create(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestTagCreate)
	tag := model.Tag{
		Name:        r.Name,
		Type:        r.Type,
		Description: &r.Description,
		OrgUUID:     uuid.MustParse(r.OrgParam),
	}

	_, err := s.repo.Create(&tag)
	if err != nil {
		return generic.NewResponse(500, []string{err.Error()}, nil), err
	}
	return generic.NewResponse(200, nil, tag), err
}

// Get gets an tag
// @Summary Get an tag
// @Description Get an tag
// @ID tag-get
// @Tags tag
// @Produce json
// @Param org path string true "organization ID"  Format(uuid)
// @Param tag path string true "tag ID" Format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/tag/{tag} [get]
func (s Tag) Get(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	var (
		r      = req.(*RequestTagGet)
		org    model.Tag
		filter = map[string]any{
			"uuid":     r.TagParam,
			"org_uuid": r.OrgParam,
		}
	)
	_, err := s.repo.Get(&org, filter)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return generic.NewResponse(404, []string{err.Error()}, nil), err
		}
		return generic.NewResponse(500, []string{err.Error()}, nil), err
	}
	return generic.NewResponse(200, nil, org), err
}

// List lists tags
// @Summary List tag tags
// @Description List tag tags
// @ID tags-get
// @Tags tag
// @Produce json
// @Param org path string true "organization ID" Format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/tag [get]
func (s Tag) List(ctx context.Context, req generic.IRequest, filter map[string]any) (generic.IResponse, error) {
	var (
		r  = req.(*RequestTagList)
		uu []model.Tag
	)
	filter["org_uuid"] = r.OrgParam
	_, err := s.repo.List(&uu, filter)
	if err != nil {
		return generic.NewResponse(500, []string{err.Error()}, nil), err
	}
	return generic.NewResponse(200, nil, uu), nil
}

// Update updates an tag
// @Summary Update an organization tag
// @Description Update an organization tag
// @ID tag-update
// @Tags tag
// @Accept json
// @Produce json
// @Param org path string true "organization ID" Format(uuid)
// @Param tag path string true "tag ID" Format(uuid)
// @Param tag body PayloadTag true "tag data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/tag/{tag} [patch]
func (s Tag) Update(ctx context.Context, req generic.IRequest) (generic.IResponse, error) {
	r := req.(*RequestTagUpdate)
	tag := model.Tag{
		BaseModel:   storage.BaseModel{UUID: uuid.MustParse(r.TagParam)},
		Name:        r.Name,
		Type:        r.Type,
		Description: &r.Description,
		OrgUUID:     uuid.MustParse(r.OrgParam),
	}
	_, err := s.repo.Update(&tag)
	if err != nil {
		return generic.NewResponse(500, []string{err.Error()}, nil), err
	}
	return generic.NewResponse(200, nil, tag), nil
}

// Delete deletes an tag
// @Summary Delete organization tag
// @Description Delete organization tag
// @ID tag-delete
// @Tags tag
// @Param org path string true "organization ID" Format(uuid)
// @Param tag path string true "tag ID" Format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /organization/{org}/tag/{tag} [delete]
func (s Tag) Delete(ctx context.Context, req generic.IRequest) error {
	var (
		org    model.Tag
		filter = map[string]any{
			"uuid": req.(*RequestTagDelete).TagParam,
		}
	)
	_, err := s.repo.Delete(&org, filter)
	if err != nil {
		return err
	}
	return nil
}

// Tag is the service interface
var _ generic.IService = new(Tag)
