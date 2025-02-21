package cacheRepository

import (
	"encoding/json"
	"fmt"


	authRepository "stncCms/app/auth/repository/dbRepository"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	"time"
	// repository "stncCms/app/domain/repository/dbRepository"
	"github.com/jinzhu/gorm"
		authEntity "stncCms/app/auth/entity"
		authDto "stncCms/app/auth/dto"
		optionRepository "stncCms/app/options/repository/dbRepository"

)

type RoleRepo struct {
	db *gorm.DB
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

// RoleRepo implements the authRepository.RoleRepository interface
//var _ services.RoleAppInterface = &RoleRepo{}

// GetAll all data
func (r *RoleRepo) GetAll() ([]authEntity.Role, error) {
	var data []authEntity.Role
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllRole(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllRole"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRole(r.db)
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
func getAllRole(db *gorm.DB) ([]authEntity.Role, error) {
	repo := authRepository.RoleRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// Save data
func (r *RoleRepo) Save(data *authEntity.Role) (*authEntity.Role, map[string]string) {
	repo := authRepository.RoleRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *RoleRepo) EditList(modulID int, roleID int) ([]authDto.RoleEditList, error) {
	repo := authRepository.RoleRepositoryInit(r.db)
	datas, err := repo.EditList(modulID, roleID)
	return datas, err
}

// Count
func (r *RoleRepo) Count(totalCount *int64) {
	var count int64
	repo := authRepository.RoleRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count
}

// Delete data
func (r *RoleRepo) Delete(id uint64) error {
	repo := authRepository.RoleRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}

// GetAllPagination pagination all data
func GetAllPaginationrole(db *gorm.DB, perPage int, offset int) ([]authEntity.Role, error) {
	repo := authRepository.RoleRepositoryInit(db)
	data, _ := repo.GetAllPagination(perPage, offset)
	return data, nil
}

// GetAllPagination pagination all data
func (r *RoleRepo) GetAllPagination(perPage int, offset int) ([]authEntity.Role, error) {
	var data []authEntity.Role
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationrole(r.db, perPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationpost_" + stnccollection.IntToString(perPage) + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationrole(r.db, perPage, offset)
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

// GetByID get data
func (r *RoleRepo) GetByID(id int) (*authEntity.Role, error) {
	var data *authEntity.Role
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDRole(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDRole_" + stnccollection.IntToString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDRole(r.db, id)
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

func getByIDRole(db *gorm.DB, id int) (*authEntity.Role, error) {
	repo := authRepository.RoleRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

func (r *RoleRepo) Update(data *authEntity.Role) (*authEntity.Role, map[string]string) {
	repo := authRepository.RoleRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *RoleRepo) UpdateTitle(id int, title string) {
	repo := authRepository.RoleRepositoryInit(r.db)
	repo.UpdateTitle(id, title)
}
