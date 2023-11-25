package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/golang-experimental/types"
	"gorm.io/gorm"
)

type sysOptionsOrm struct {
	db *gorm.DB
}

type SysOptions struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	OptionName string    `json:"-"`
	Value      string    `json:"-"`
	Status     bool      `gorm:"default:true" json:"false"`

	types.DefaultModelProperty
}

type SysOptionsModelAction interface {
	GetOneByID(id uuid.UUID) (sysOpt SysOptions, err error)
	GetOneSystemOptions(username string) (sysOpt SysOptions, err error)
	InsertSystemOptions(sysOpt SysOptions) (id uuid.UUID, err error)
	UpdateSystemOptions(sysOpt SysOptions) (err error)
}

func NewSysOptionsAction(db *gorm.DB) SysOptionsModelAction {
	return &sysOptionsOrm{db}
}

func (o *sysOptionsOrm) GetOneByID(id uuid.UUID) (sysOpt SysOptions, err error) {
	result := o.db.Model(&SysOptions{}).Where("id = ?", id).First(&sysOpt)
	return sysOpt, result.Error
}

func (o *sysOptionsOrm) GetOneSystemOptions(username string) (sysOpt SysOptions, err error) {
	result := o.db.Model(&SysOptions{}).Where("option_name = ?", username).First(&sysOpt)
	return sysOpt, result.Error
}

func (o *sysOptionsOrm) InsertSystemOptions(sysOpt SysOptions) (id uuid.UUID, err error) {
	result := o.db.Model(&SysOptions{}).Create(&sysOpt)
	return sysOpt.ID, result.Error
}

func (o *sysOptionsOrm) UpdateSystemOptions(sysOpt SysOptions) (err error) {
	result := o.db.Model(&SysOptions{}).Model(&sysOpt).Updates(&sysOpt)
	return result.Error
}
