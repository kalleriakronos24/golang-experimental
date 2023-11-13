package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type SysOptions struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	OptionName string    `json:"-"`
	Value      string    `json:"-"`
	Status     bool      `gorm:"default:true" json:"false"`

	gorm.Model
}
