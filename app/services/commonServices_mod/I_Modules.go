package commonServices_mod

import (
	// "stncCms/app/domain/entity"
		modulesEntity "stncCms/app/modules/entity"
		modulesDTO "stncCms/app/modules/dto"
)

// ModuleAppInterface interface
type ModulesAppInterface interface {
	GetAll() ([]modulesEntity.Modules, error)
	GetAllModulesMerge() ([]modulesEntity.ModulesAndPermissionDTO, error)
	GetAllModulesMergePermission() ([]modulesDTO.ModulesAndPermissionRoleDTO, error)

	Save(*modulesEntity.Modules) (*modulesEntity.Modules, map[string]string)
	GetByID(uint64) (*modulesEntity.Modules, error)
	Update(*modulesEntity.Modules) (*modulesEntity.Modules, map[string]string)
	Delete(uint64) error
	GetAllPaginate(postsPerPage int, offset int) ([]modulesEntity.Modules, error)
	GetAllPaginateCount(returnValue *int64)
}

type ModuleApp struct {
	request ModulesAppInterface
}

// UserApp implements the UserAppInterface
var _ ModulesAppInterface = &ModuleApp{}

func (f *ModuleApp) GetAll() ([]modulesEntity.Modules, error) {
	return f.request.GetAll()
}

func (f *ModuleApp) GetAllModulesMerge() ([]modulesEntity.ModulesAndPermissionDTO, error) {
	return f.request.GetAllModulesMerge()
}
func (f *ModuleApp) GetAllModulesMergePermission() ([]modulesDTO.ModulesAndPermissionRoleDTO, error) {
	return f.request.GetAllModulesMergePermission()
}

// Save service init
func (f *ModuleApp) Save(Cat *modulesEntity.Modules) (*modulesEntity.Modules, map[string]string) {
	return f.request.Save(Cat)
}

// GetByID service init
func (f *ModuleApp) GetByID(catID uint64) (*modulesEntity.Modules, error) {
	return f.request.GetByID(catID)
}

// Update service init
func (f *ModuleApp) Update(cat *modulesEntity.Modules) (*modulesEntity.Modules, map[string]string) {
	return f.request.Update(cat)
}

// Delete service init
func (f *ModuleApp) Delete(catID uint64) error {
	return f.request.Delete(catID)
}

// GetAllPaginate list
func (f *ModuleApp) GetAllPaginate(postsPerPage int, offset int) ([]modulesEntity.Modules, error) {
	return f.request.GetAllPaginate(postsPerPage, offset)
}

// GetAllPaginateCount toplam adet
func (f *ModuleApp) GetAllPaginateCount(returnValue *int64) {
	f.request.GetAllPaginateCount(returnValue)
}
