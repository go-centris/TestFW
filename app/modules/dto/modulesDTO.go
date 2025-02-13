package dto

import (
	authDto "stncCms/app/auth/dto"
	authEntity "stncCms/app/auth/entity"
	modulesEntity "stncCms/app/modules/entity"
)



type RbcaCheck struct {
	ModulID              int    ` json:"modulID"`
	RoleID               int    ` json:"roleID"`
	PermissionID         int    ` json:"permissionID"`
	RolePermissionActive int    ` json:"rolePermissionActive"`
	PermissionTitle      string ` json:"permissionTitle"`
	Controller           string ` json:"controller"`
	Function             string ` json:"funcName"`
	PermissionName       string ` json:"PermissionName"`
}

type ModulesAndPermissionDTO struct {
	modulesEntity.Modules
	Permissions []authEntity.Permission
}

type ModulesAndPermissionRoleDTO struct {
	modulesEntity.Modules
	RoleEditList []authDto.RoleEditList
}
