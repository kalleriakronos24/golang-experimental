package dto

import "github.com/gofrs/uuid"

type CreateMasterModule struct {
	ID          uuid.UUID `json:"-"`
	ModuleName  string    `json:"moduleName,omitempty" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

type UpdateMasterModule struct {
	ID          uuid.UUID `json:"-"`
	ModuleName  string    `json:"moduleName,omitempty" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

type RetrieveOneMasterModule struct {
	ModuleName  string `json:"moduleName" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}
