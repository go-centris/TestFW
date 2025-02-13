package cacheRepository

import (
	"encoding/json"
	"fmt"

	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	modulesEntity "stncCms/app/modules/entity"
	modulesDto "stncCms/app/modules/dto"
	"time"

	
	"github.com/jinzhu/gorm"
)

type ModulesRepo struct {
	db *gorm.DB
}

func ModulesRepositoryInit(db *gorm.DB) *ModulesRepo {
	return &ModulesRepo{db}
}

//var _ services.ModulesAppInterface = &ModulesRepo{}

// GetAll all data
func (r *ModulesRepo) GetAll() ([]modulesEntity.Modules, error) {
	var data []modulesEntity.Modules
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllModules(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllModules"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllModules(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
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
func getAllModules(db *gorm.DB) ([]modulesEntity.Modules, error) {
	repo := repository.ModulesRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// GetAll all data
func (r *ModulesRepo) GetAllModulesMerge() ([]modulesDto.ModulesAndPermissionDTO, error) {
	var data []modulesDto.ModulesAndPermissionDTO
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllModulesMergeModules(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllModules"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllModulesMergeModules(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
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
func getAllModulesMergeModules(db *gorm.DB) ([]modulesDto.ModulesAndPermissionDTO, error) {
	repo := repository.ModulesRepositoryInit(db)
	data, _ := repo.GetAllModulesMerge()
	return data, nil
}

// GetAll all data
func (r *ModulesRepo) GetAllModulesMergePermission() ([]modulesDto.ModulesAndPermissionRoleDTO, error) {
	var data []modulesDto.ModulesAndPermissionRoleDTO
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllModulesMergePermission(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllModules"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllModulesMergePermission(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
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
func getAllModulesMergePermission(db *gorm.DB) ([]modulesDto.ModulesAndPermissionRoleDTO, error) {
	repo := repository.ModulesRepositoryInit(db)
	data, _ := repo.GetAllModulesMergePermission()
	return data, nil
}

//*****///

func getByIDRegionModules(db *gorm.DB, id uint64) (*modulesEntity.Modules, error) {
	repo := repository.ModulesRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// GetByID get data
func (r *ModulesRepo) GetByID(id uint64) (*modulesEntity.Modules, error) {

	var data *modulesEntity.Modules
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDRegionModules(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByIDRegionModules" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDRegionModules(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("Create Key")
			if err != nil {
				fmt.Println("hata baş")
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

// GetAllPaginate pagination
func (r *ModulesRepo) GetAllPaginate(postsPerPage int, offset int) ([]modulesEntity.Modules, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []modulesEntity.Modules
	if cacheControl == "false" {
		data, _ = getAllPaginateForModules(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllPaginateForModules_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllPaginateForModules(r.db, postsPerPage, offset)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
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

// getAllPaginate
func getAllPaginateForModules(db *gorm.DB, postsPerPage int, offset int) ([]modulesEntity.Modules, error) {
	repo := repository.ModulesRepositoryInit(db)
	data, _ := repo.GetAllPaginate(postsPerPage, offset)
	return data, nil
}

// GetAllPaginateCount
func (r *ModulesRepo) GetAllPaginateCount(returnValue *int64) {
	var count int64
	repo := repository.ModulesRepositoryInit(r.db)
	repo.GetAllPaginateCount(&count)
	*returnValue = count
}

// Save data
func (r *ModulesRepo) Save(data *modulesEntity.Modules) (*modulesEntity.Modules, map[string]string) {
	repo := repository.ModulesRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *ModulesRepo) Update(data *modulesEntity.Modules) (*modulesEntity.Modules, map[string]string) {
	repo := repository.ModulesRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err

}

// Delete delete data
func (r *ModulesRepo) Delete(id uint64) error {
	repo := repository.ModulesRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
