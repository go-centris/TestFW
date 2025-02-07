package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/helpers/stnccollection"
	repository "stncCms/app/domain/repository/dbRepository"
	"time"

	"github.com/jinzhu/gorm"
)

// ReportRepo struct
type ReportRepo struct {
	db *gorm.DB
}

// ReportRepositoryInit initial
func ReportRepositoryInit(db *gorm.DB) *ReportRepo {
	return &ReportRepo{db}
}

func getAllUsersWhoAddedMostSacrifeAndBranch(db *gorm.DB, postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	repo := repository.ReportRepositoryInit(db)
	data, _ := repo.GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage, offset)
	return data, nil
}

// GetAllUsersWhoAddedMostSacrifeAndBranch pagination all data
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []dto.UsersWhoAddedMostSacrifeAndBranch
	if cacheControl == "false" {
		data, _ = getAllUsersWhoAddedMostSacrifeAndBranch(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "UsersWhoAddedMostSacrifeAndBranchList_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllUsersWhoAddedMostSacrifeAndBranch(r.db, postsPerPage, offset)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data, nil
}

// GetAllUsersWhoAddedMostSacrifeAndBranchCount fat
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndBranchCount(totalCount *int64) {
	var count int64
	repo := repository.ReportRepositoryInit(r.db)
	repo.GetAllUsersWhoAddedMostSacrifeAndBranchCount(&count)
	*totalCount = count
}

//-------------------////

func getAllUsersWhoAddedMostSacrifeAndUser(db *gorm.DB, postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	repo := repository.ReportRepositoryInit(db)
	data, _ := repo.GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage, offset)
	return data, nil
}

// GetAllUsersWhoAddedMostSacrifeAndUser pagination all data
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []dto.UsersWhoAddedMostSacrifeAndBranch
	if cacheControl == "false" {
		data, _ = getAllUsersWhoAddedMostSacrifeAndUser(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllUsersWhoAddedMostSacrifeAndUser_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllUsersWhoAddedMostSacrifeAndUser(r.db, postsPerPage, offset)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data, nil
}

// GetAllUsersWhoAddedMostSacrifeAndUserCount fat
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndUserCount(totalCount *int64) {
	var count int64
	repo := repository.ReportRepositoryInit(r.db)
	repo.GetAllUsersWhoAddedMostSacrifeAndUserCount(&count)
	*totalCount = count
}

// *** En cok kurban kestiren hayirsever ***/////

// getAllCharitableWhoAddedMostSacrife En cok kurban kestiren hayirsever
func getAllCharitableWhoAddedMostSacrife(db *gorm.DB, postsPerPage int, offset int) ([]dto.CharitableWhoAddedMostSacrife, error) {
	repo := repository.ReportRepositoryInit(db)
	data, _ := repo.GetAllCharitableWhoAddedMostSacrife(postsPerPage, offset)
	return data, nil
}

// GetAllCharitableWhoAddedMostSacrife  En cok kurban kestiren hayirsever
func (r *ReportRepo) GetAllCharitableWhoAddedMostSacrife(postsPerPage int, offset int) ([]dto.CharitableWhoAddedMostSacrife, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []dto.CharitableWhoAddedMostSacrife
	if cacheControl == "false" {
		data, _ = getAllCharitableWhoAddedMostSacrife(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllCharitableWhoAddedMostSacrife_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllCharitableWhoAddedMostSacrife(r.db, postsPerPage, offset)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data, nil
}

// GetAllCharitableWhoAddedMostSacrifeCount  En cok kurban kestiren hayirsever count
func (r *ReportRepo) GetAllCharitableWhoAddedMostSacrifeCount(totalCount *int64) {
	var count int64
	repo := repository.ReportRepositoryInit(r.db)
	repo.GetAllCharitableWhoAddedMostSacrifeCount(&count)
	*totalCount = count
}
