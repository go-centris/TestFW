package commonServices_mod

import (
	"stncCms/app/domain/entity"
)

// ModuleAppInterface interface
type ModulesAppInterface interface {
	GetAll() ([]entity.Modules, error)
	GetAllModulesMerge() ([]entity.ModulesAndPermissionDTO, error)
	GetAllModulesMergePermission() ([]entity.ModulesAndPermissionRoleDTO, error)

	Save(*entity.Modules) (*entity.Modules, map[string]string)
	GetByID(uint64) (*entity.Modules, error)
	Update(*entity.Modules) (*entity.Modules, map[string]string)
	Delete(uint64) error
	GetAllPaginate(postsPerPage int, offset int) ([]entity.Modules, error)
	GetAllPaginateCount(returnValue *int64)
}

type ModuleApp struct {
	request ModulesAppInterface
}

// UserApp implements the UserAppInterface
var _ ModulesAppInterface = &ModuleApp{}

func (f *ModuleApp) GetAll() ([]entity.Modules, error) {
	return f.request.GetAll()
}

func (f *ModuleApp) GetAllModulesMerge() ([]entity.ModulesAndPermissionDTO, error) {
	return f.request.GetAllModulesMerge()
}
func (f *ModuleApp) GetAllModulesMergePermission() ([]entity.ModulesAndPermissionRoleDTO, error) {
	return f.request.GetAllModulesMergePermission()
}

// Save service init
func (f *ModuleApp) Save(Cat *entity.Modules) (*entity.Modules, map[string]string) {
	return f.request.Save(Cat)
}

// GetByID service init
func (f *ModuleApp) GetByID(catID uint64) (*entity.Modules, error) {
	return f.request.GetByID(catID)
}

// Update service init
func (f *ModuleApp) Update(cat *entity.Modules) (*entity.Modules, map[string]string) {
	return f.request.Update(cat)
}

// Delete service init
func (f *ModuleApp) Delete(catID uint64) error {
	return f.request.Delete(catID)
}

// GetAllPaginate list
func (f *ModuleApp) GetAllPaginate(postsPerPage int, offset int) ([]entity.Modules, error) {
	return f.request.GetAllPaginate(postsPerPage, offset)
}

// GetAllPaginateCount toplam adet
func (f *ModuleApp) GetAllPaginateCount(returnValue *int64) {
	f.request.GetAllPaginateCount(returnValue)
}
