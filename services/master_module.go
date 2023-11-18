package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	masterModels "github.com/kalleriakronos24/mygoapp2nd/models/master"
)

func (module *module) RetrieveMasterModule(p dto.RetrieveOneMasterModule) (m masterModels.Module, err error) {
	if m, err = module.db.masterModuleModel.GetOneModule(p); err != nil {
		return m, errors.New(err.Error())
	}
	return m, err
}

func (module *module) CreateMasterModule(p dto.CreateMasterModule) (m masterModels.Module, err error) {
	if m, err = module.db.masterModuleModel.InsertModule(p); err != nil {
		return m, errors.New(err.Error())
	}
	return
}

func (module *module) UpdateMasterModule(id uuid.UUID, pMasterModule dto.UpdateMasterModule) (err error) {
	if err = module.db.masterModuleModel.UpdateModule(id, masterModels.Module{
		ModuleName:  pMasterModule.ModuleName,
		Description: pMasterModule.Description,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) DeleteMasterModule(id uuid.UUID) (err error) {
	if err = module.db.masterModuleModel.DeleteModule(id); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingMasterModule(moduleName string, id string) (err error) {

	if moduleName != "" {
		if _, dbErr := module.db.masterModuleModel.GetOneByModuleName(moduleName); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.masterModuleModel.GetOneByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
