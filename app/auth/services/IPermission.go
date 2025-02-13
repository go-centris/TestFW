package authServices_mod

import (


	modulesDto "stncCms/app/modules/dto"
authEntity "stncCms/app/auth/entity"
)

// PermissionAppInterface interface
type PermissionAppInterface interface {
	GetAll() ([]authEntity.Permission, error)
	GetAllPaginationermissionForModulID(int) ([]authEntity.Permission, error)
	GetUserPermission(int) ([]modulesDto.RbcaCheck, error)
	GetUserPermissionForComponent(int, string) ([]modulesDto.RbcaCheck, error)
}

type permissionApp struct {
	request PermissionAppInterface
}

// UserApp implements the UserAppInterface
var _ PermissionAppInterface = &permissionApp{}

func (f *permissionApp) GetAll() ([]authEntity.Permission, error) {
	return f.request.GetAll()
}

func (f *permissionApp) GetAllPaginationermissionForModulID(modulId int) ([]authEntity.Permission, error) {
	return f.request.GetAllPaginationermissionForModulID(modulId)
}

func (f *permissionApp) GetUserPermission(roleID int) ([]modulesDto.RbcaCheck, error) {
	return f.request.GetUserPermission(roleID)
}
func (f *permissionApp) GetUserPermissionForComponent(roleID int, componentBaseName string) ([]modulesDto.RbcaCheck, error) {
	return f.request.GetUserPermissionForComponent(roleID, componentBaseName)
}
