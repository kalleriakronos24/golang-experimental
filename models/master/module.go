package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Module struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ModuleName  string    `json:"-"`
	Description string    `json:"-"`
	Status      bool      `gorm:"default:true" json:"false"`

	gorm.Model
}
