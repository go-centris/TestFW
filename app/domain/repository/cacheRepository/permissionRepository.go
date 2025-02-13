package cacheRepository

import (
	"encoding/json"
	"fmt"
	

		repository "stncCms/app/domain/repository/dbRepository"
		authRepository "stncCms/app/auth/repository/dbRepository"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	"time"
	authEntity "stncCms/app/auth/entity"
	modulesDto "stncCms/app/modules/dto"
	"github.com/jinzhu/gorm"
)

type PermissionRepo struct {
	db *gorm.DB
}

func PermissionRepositoryInit(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db}
}

// PermissionRepo implements the repository.PermissionRepository interface
//var _ services.PermissionAppInterface = &PermissionRepo{}

// GetAll all data
func (r *PermissionRepo) GetAll() ([]authEntity.Permission, error) {
	var data []authEntity.Permission
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationermission(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationermission"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationermission(r.db)
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
func GetAllPaginationermission(db *gorm.DB) ([]authEntity.Permission, error) {
	repo := authRepository.PermissionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// getAllPaginationermissionForModulID all data
func (r *PermissionRepo) GetAllPaginationermissionForModulID(modulId int) ([]authEntity.Permission, error) {
	var data []authEntity.Permission
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllPaginationermissionForModulID(modulId, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllPaginationermissionForModulID" + stnccollection.IntToString(modulId)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllPaginationermissionForModulID(modulId, r.db)
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
func getAllPaginationermissionForModulID(modulId int, db *gorm.DB) ([]authEntity.Permission, error) {
	repo := authRepository.PermissionRepositoryInit(db)
	data, _ := repo.GetAllPaginationermissionForModulID(modulId)
	return data, nil
}

// GetUserPermission permissinon listesi
func (r *PermissionRepo) GetUserPermission(roleID int) ([]modulesDto.RbcaCheck, error) {
	var data []modulesDto.RbcaCheck
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getUserPermission(roleID, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetUserPermission" + stnccollection.IntToString(roleID)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getUserPermission(roleID, r.db)
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
func getUserPermission(roleID int, db *gorm.DB) ([]modulesDto.RbcaCheck, error) {
	repo := authRepository.PermissionRepositoryInit(db)
	data, _ := repo.GetUserPermission(roleID)
	return data, nil
}

// GetUserPermission permissinon listesi
func (r *PermissionRepo) GetUserPermissionForComponent(roleID int, componentBaseName string) ([]modulesDto.RbcaCheck, error) {
	var data []modulesDto.RbcaCheck
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getUserPermissionForComponent(roleID, componentBaseName, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetUserPermissionForComponent" + stnccollection.IntToString(roleID) + componentBaseName
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getUserPermissionForComponent(roleID, componentBaseName, r.db)
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
func getUserPermissionForComponent(roleID int, componentBaseName string, db *gorm.DB) ([]modulesDto.RbcaCheck, error) {
	repo := authRepository.PermissionRepositoryInit(db)
	data, _ := repo.GetUserPermissionForComponent(roleID, componentBaseName)
	return data, nil
}
