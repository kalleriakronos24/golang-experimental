package models

import (
	"time"

	"gorm.io/gorm"
)

type userLogOrm struct {
	db *gorm.DB
}

type UserLog struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Action    string `json:"-"`
	UserID    uint
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
}
