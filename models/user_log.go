package models

import (
	masterModels "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"github.com/kalleriakronos24/mygoapp2nd/types"
	"gorm.io/gorm"
)

type userLogOrm struct {
	db *gorm.DB
}

type UserLog struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	Action      string `json:"-"`
	Description string
	Kind        string // common / error

	UserID uint
	User   masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	types.DefaultModelProperty
}
