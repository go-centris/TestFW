package cacheRepository

import (
	"encoding/json"
	"fmt"
	// "stncCms/app/domain/entity"
	// repository "stncCms/app/domain/repository/dbRepository"
	branchRepository "stncCms/app/branch/repository/dbRepository"
	optionRepository "stncCms/app/options/repository/dbRepository"
	// modulesDTO "stncCms/app/modules/dto"
	branchEntity "stncCms/app/branch/entity"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	"time"

	"github.com/jinzhu/gorm"
)

// BranchRepo struct
type BranchRepo struct {
	db *gorm.DB
}

// BranchRepositoryInit initial
func BranchRepositoryInit(db *gorm.DB) *BranchRepo {
	return &BranchRepo{db}
}

func getByIDBranch(db *gorm.DB, id uint64) (*branchEntity.Branches, error) {
	repo := branchRepository.BranchRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// GetByID get data
func (r *BranchRepo) GetByID(id uint64) (*branchEntity.Branches, error) {

	var data *branchEntity.Branches
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDBranch(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "branchGetByID" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDBranch(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("Create Key")
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

// GetAll all data
func (r *BranchRepo) GetAll() ([]branchEntity.Branches, error) {
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []branchEntity.Branches
	if cacheControl == "false" {
		data, _ = getAllbranch(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "branchGetAll"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllbranch(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("Create Key")
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

func getAllbranch(db *gorm.DB) ([]branchEntity.Branches, error) {
	repo := branchRepository.BranchRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

func (r *BranchRepo) GetByRegionID(regionID uint64) ([]branchEntity.Branches, error) {
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []branchEntity.Branches
	if cacheControl == "false" {
		data, _ = getByRegionID(regionID, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetByRegionID" + stnccollection.Uint64toString(regionID)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByRegionID(regionID, r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("Create Key")
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

func getByRegionID(regionID uint64, db *gorm.DB) ([]branchEntity.Branches, error) {
	repo := branchRepository.BranchRepositoryInit(db)
	data, _ := repo.GetByRegionID(regionID)
	return data, nil
}

// GetAllPaginate
func (r *BranchRepo) GetAllPaginate(postsPerPage int, offset int) ([]branchEntity.Branches, error) {
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []branchEntity.Branches
	if cacheControl == "false" {
		data, _ = getAllPaginateRegion(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllPaginateRegion_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllPaginateRegion(r.db, postsPerPage, offset)
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

// getAllPaginateRegion
func getAllPaginateRegion(db *gorm.DB, postsPerPage int, offset int) ([]branchEntity.Branches, error) {
	repo := branchRepository.BranchRepositoryInit(db)
	data, _ := repo.GetAllPaginate(postsPerPage, offset)
	return data, nil
}

// GetAllPaginateForRegionCount
func (r *BranchRepo) GetAllPaginateCount(returnValue *int64) {
	var count int64
	repo := branchRepository.BranchRepositoryInit(r.db)
	repo.GetAllPaginateCount(&count)
	*returnValue = count
}
