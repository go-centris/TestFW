package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	"time"
	optionRepository "stncCms/app/options/repository/dbRepository"

	"github.com/jinzhu/gorm"
)

// RegionRepo struct
type RegionRepo struct {
	db *gorm.DB
}

// RegionRepositoryInit initial
func RegionRepositoryInit(db *gorm.DB) *RegionRepo {
	return &RegionRepo{db}
}

func getByIDRegion(db *gorm.DB, id uint64) (*entity.Region, error) {
	repo := repository.RegionRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// GetByID get data
func (r *RegionRepo) GetByID(id uint64) (*entity.Region, error) {

	var data *entity.Region
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDRegion(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "RegionGetByID" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDRegion(r.db, id)
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
func (r *RegionRepo) GetAll() ([]entity.Region, error) {
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []entity.Region
	if cacheControl == "false" {
		data, _ = getAllRegion(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "RegionGetAll"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRegion(r.db)
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

func getAllRegion(db *gorm.DB) ([]entity.Region, error) {
	repo := repository.RegionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// GetAllPaginate
func (r *RegionRepo) GetAllPaginate(postsPerPage int, offset int) ([]entity.Region, error) {
	access := optionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []entity.Region
	if cacheControl == "false" {
		data, _ = getAllPaginate(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginate_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllPaginate(r.db, postsPerPage, offset)
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

// getAllPaginate
func getAllPaginate(db *gorm.DB, postsPerPage int, offset int) ([]entity.Region, error) {
	repo := repository.RegionRepositoryInit(db)
	data, _ := repo.GetAllPaginate(postsPerPage, offset)
	return data, nil
}

// GetAllPaginateCount
func (r *RegionRepo) GetAllPaginateCount(returnValue *int64) {
	var count int64
	repo := repository.RegionRepositoryInit(r.db)
	repo.GetAllPaginateCount(&count)
	*returnValue = count
}

// Save data
func (r *RegionRepo) Save(data *entity.Region) (*entity.Region, map[string]string) {
	repo := repository.RegionRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *RegionRepo) Update(data *entity.Region) (*entity.Region, map[string]string) {
	repo := repository.RegionRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err

}

// Delete delete data
func (r *RegionRepo) Delete(id uint64) error {
	repo := repository.RegionRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
