package models

import (
	"github.com/gofrs/uuid"
	masterModels "github.com/kalleriakronos24/golang-experimental/models/master"
	"github.com/kalleriakronos24/golang-experimental/types"
)

type UserModulePermission struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Permission  string    `json:"-"`
	Description string    `json:"-"`

	ModuleID uint
	Module   masterModels.Module `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ModuleID"`
	UserID   uint
	User     masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	types.DefaultModelProperty
}
