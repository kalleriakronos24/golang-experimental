package models

import (
	"github.com/gofrs/uuid"
	masterModels "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"github.com/kalleriakronos24/mygoapp2nd/types"
)

type UserModule struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ModuleName string    `json:"-"`

	UserID uint
	User   masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	types.DefaultModelProperty
}
