package cacheRepository

import (
	"encoding/json"
	"fmt"
	authRepository "stncCms/app/auth/repository/dbRepository"

	"stncCms/pkg/cache"
	"time"
	repository "stncCms/app/domain/repository/dbRepository"
	"github.com/jinzhu/gorm"
	
	authEntity "stncCms/app/auth/entity"
)

type RolePermissionRepo struct {
	db *gorm.DB
}

func RolePermissionRepositoryInit(db *gorm.DB) *RolePermissionRepo {
	return &RolePermissionRepo{db}
}

// PermissionRepo implements the authRepo.PermissionRepository interface
//var _ services.RolePermissionAppInterface = &RolePermissionRepo{}

// GetAll all data
func (r *RolePermissionRepo) GetAll() ([]authEntity.RolePermisson, error) {
	var data []authEntity.RolePermisson
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllRolePermission(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationermission"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRolePermission(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("Create Key Error")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("Redis Error")
		}
	}
	return data, nil
}
func getAllRolePermission(db *gorm.DB) ([]authEntity.RolePermisson, error) {
	repo := authRepository.RolePermissionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// Save data
func (r *RolePermissionRepo) Save(data *authEntity.RolePermisson) (*authEntity.RolePermisson, map[string]string) {
	repo := authRepository.RolePermissionRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *RolePermissionRepo) Update(data *authEntity.RolePermisson) (*authEntity.RolePermisson, map[string]string) {
	repo := authRepository.RolePermissionRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

func (r *RolePermissionRepo) UpdateActiveStatus(roleId int, permissionId int, active int) {
	repo := authRepository.RolePermissionRepositoryInit(r.db)
	repo.UpdateActiveStatus(roleId, permissionId, active)
}
