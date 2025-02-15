package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/post/entity"
	PostRepository "stncCms/app/post/repository/dbRepository"
	OptionRepository "stncCms/app/options/repository/dbRepository"
	"stncCms/pkg/cache"
	"stncCms/pkg/helpers/stnccollection"
	"time"

	"github.com/jinzhu/gorm"
)

// PostRepo struct
type PostRepo struct {
	db *gorm.DB
}

// PostRepositoryInit initial
func PostRepositoryInit(db *gorm.DB) *PostRepo {
	return &PostRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.PostAppInterface = &PostRepo{}

// GetByID get data
func (r *PostRepo) GetByID(id uint64) (*entity.Post, error) {
	var data *entity.Post
	access := OptionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDPost(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDPost_" + stnccollection.Uint64toString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDPost(r.db, id)
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

func getByIDPost(db *gorm.DB, id uint64) (*entity.Post, error) {
	repo := PostRepository.PostRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

// GetAll all data
func (r *PostRepo) GetAll() ([]entity.Post, error) {
	var data []entity.Post
	access := OptionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationost(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationost"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationost(r.db)
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

// GetAll all data
func GetAllPaginationost(db *gorm.DB) ([]entity.Post, error) {
	repo := PostRepository.PostRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// GetAllPagination pagination all data
func GetAllPaginationPost(db *gorm.DB, perPage int, offset int) ([]entity.Post, error) {
	repo := PostRepository.PostRepositoryInit(db)
	data, _ := repo.GetAllPagination(perPage, offset)
	return data, nil
}

// GetAllPagination pagination all data
func (r *PostRepo) GetAllPagination(perPage int, offset int) ([]entity.Post, error) {

	var data []entity.Post
	access := OptionRepository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationPost(r.db, perPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationpost_" + stnccollection.IntToString(perPage) + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationPost(r.db, perPage, offset)
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
