package models

import (
	masterModel "github.com/kalleriakronos24/mygoapp2nd/models/master"
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
	User   masterModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	gorm.Model
}
