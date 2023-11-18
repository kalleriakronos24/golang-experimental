package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	"github.com/kalleriakronos24/mygoapp2nd/types"
	"gorm.io/gorm"
)

type moduleOrm struct {
	db *gorm.DB
}

type Module struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ModuleName  string    `json:"-"`
	Description string    `json:"-"`
	Status      bool      `gorm:"default:true" json:"false"`

	types.DefaultModelProperty
}

func (m *Module) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

type ModuleModelAction interface {
	GetOneByID(id uuid.UUID) (m Module, err error)
	GetOneModule(p dto.RetrieveOneMasterModule) (m Module, err error)
	GetOneByModuleName(moduleName string) (m Module, err error)
	InsertModule(p dto.CreateMasterModule) (m Module, err error)
	UpdateModule(id uuid.UUID, m Module) (err error)
	DeleteModule(id uuid.UUID) (err error)
}

func NewModuleAction(db *gorm.DB) ModuleModelAction {
	return &moduleOrm{db}
}

func (o *moduleOrm) GetOneByID(id uuid.UUID) (m Module, err error) {
	result := o.db.Model(&Module{}).First(&m, id)
	return m, result.Error
}

func (o *moduleOrm) GetOneByModuleName(moduleName string) (m Module, err error) {
	result := o.db.Model(&Module{}).Where("module_name = ?", moduleName).First(&m)
	return m, result.Error
}

func (o *moduleOrm) GetOneModule(p dto.RetrieveOneMasterModule) (m Module, err error) {
	result := o.db.Model(&Module{}).First(&m, p)
	return m, result.Error
}

func (o *moduleOrm) InsertModule(p dto.CreateMasterModule) (m Module, err error) {

	result := o.db.Model(&Module{}).Create(&Module{
		ModuleName:  p.ModuleName,
		Description: p.Description,
	})
	return m, result.Error
}

func (o *moduleOrm) UpdateModule(id uuid.UUID, m Module) (err error) {
	result := o.db.Model(&m).Where("id", id).Updates(&m)
	return result.Error
}

func (o *moduleOrm) DeleteModule(id uuid.UUID) (err error) {
	result := o.db.Model(&Module{}).Unscoped().Delete(&Module{}, id)
	return result.Error
}
