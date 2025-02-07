package authServices_mod

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

// PermissionAppInterface interface
type PermissionAppInterface interface {
	GetAll() ([]entity.Permission, error)
	GetAllPaginationermissionForModulID(int) ([]entity.Permission, error)
	GetUserPermission(int) ([]dto.RbcaCheck, error)
	GetUserPermissionForComponent(int, string) ([]dto.RbcaCheck, error)
}

type permissionApp struct {
	request PermissionAppInterface
}

// UserApp implements the UserAppInterface
var _ PermissionAppInterface = &permissionApp{}

func (f *permissionApp) GetAll() ([]entity.Permission, error) {
	return f.request.GetAll()
}

func (f *permissionApp) GetAllPaginationermissionForModulID(modulId int) ([]entity.Permission, error) {
	return f.request.GetAllPaginationermissionForModulID(modulId)
}

func (f *permissionApp) GetUserPermission(roleID int) ([]dto.RbcaCheck, error) {
	return f.request.GetUserPermission(roleID)
}
func (f *permissionApp) GetUserPermissionForComponent(roleID int, componentBaseName string) ([]dto.RbcaCheck, error) {
	return f.request.GetUserPermissionForComponent(roleID, componentBaseName)
}
