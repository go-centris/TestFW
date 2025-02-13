package dbRepository

import (
	"errors"
	

	modulesDTO "stncCms/app/modules/dto"
	authEntity "stncCms/app/auth/entity"

	"github.com/jinzhu/gorm"
	Iauth "stncCms/app/auth/services"
)

type PermissionRepo struct {
	db *gorm.DB
}

func PermissionRepositoryInit(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db}
}

// PermissionRepo implements the repository.PermissionRepository interface
var _ Iauth.PermissionAppInterface = &PermissionRepo{}

// GetAll all data
func (r *PermissionRepo) GetAll() ([]authEntity.Permission, error) {
	var datas []authEntity.Permission
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *PermissionRepo) GetAllPaginationermissionForModulID(modulId int) ([]authEntity.Permission, error) {
	var datas []authEntity.Permission
	var err error
	err = r.db.Debug().Where("modul_id = ?", modulId).Order("id desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *PermissionRepo) GetUserPermission(roleID int) ([]modulesDTO.RbcaCheck, error) {
	var data []modulesDTO.RbcaCheck
	var err error
	query := r.db.Table(authEntity.RolePermissonTableName + " AS role_permission")
	query = query.Select(`permission.modul_id AS modul_id,
	         role_permission.role_id,role_permission.permission_id,
		     role_permission.active AS role_permission_active ,
			 permission.title AS permission_Title,
			 permission.controller AS controller,
			 permission.func_name AS functionName ,
		   CONCAT (permission.controller,'-',permission.func_name) As permission_name`)
	query = query.Joins(" INNER JOIN rbca_permission AS  permission ON permission.id=role_permission.permission_id   ")
	query = query.Where("role_permission.role_id=? ", roleID)
	query = query.Where("role_permission.active=1 ")
	err = query.Find(&data).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return data, nil
}

func (r *PermissionRepo) GetUserPermissionForComponent(roleID int, componentBaseName string) ([]modulesDTO.RbcaCheck, error) {
	var data []modulesDTO.RbcaCheck
	var err error
	query := r.db.Table(authEntity.RolePermissonTableName + " AS role_permission")
	query = query.Select(`permission.modul_id AS modul_id,
	         role_permission.role_id,role_permission.permission_id,
		     role_permission.active AS role_permission_active ,
			 permission.title AS permission_Title,
			 permission.controller AS controller,
			 permission.func_name AS function ,
		   CONCAT (permission.controller,'-',permission.func_name) As permission_name`)
	query = query.Joins(" INNER JOIN rbca_permission AS  permission ON permission.id=role_permission.permission_id   ")
	query = query.Where("permission.controller=? ", componentBaseName)
	query = query.Where("role_permission.role_id=? ", roleID)
	query = query.Where("role_permission.active=1 ")
	err = query.Find(&data).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return data, nil
}
