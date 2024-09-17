package model

import (
	orgmodel "ekolo/account/model"
	"ekolo/pkg/storage"

	"github.com/google/uuid"
)

type Tag struct {
	storage.BaseModel
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	Description *string               `json:"description"`
	OrgUUID     uuid.UUID             `json:"org"`
	Org         orgmodel.Organization `json:"-"`
}
