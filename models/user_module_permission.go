package models

import (
	"github.com/gofrs/uuid"
	master "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"gorm.io/gorm"
)

type UserModulePermission struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Permission string    `json:"-"`

	ModuleID uint
	Module   master.Module `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ModuleID"`
	UserID   uint
	User     master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	gorm.Model
}
