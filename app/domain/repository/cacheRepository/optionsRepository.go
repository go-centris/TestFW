package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/pkg/cache"
	"time"

	"github.com/jinzhu/gorm"
)

var optionTableName string = "options"

// OptionRepositoryRepo struct
type OptionRepositoryRepo struct {
	db *gorm.DB
}

// OptionRepositoryInit initial
func OptionRepositoryInit(db *gorm.DB) *OptionRepositoryRepo {
	return &OptionRepositoryRepo{db}
}

func getAllOptions(db *gorm.DB) ([]entity.Options, error) {
	repo := repository.OptionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// GetAll all data
func (r *OptionRepositoryRepo) GetAll() ([]entity.Options, error) {
	var data []entity.Options
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllOptions(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllOptions"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllOptions(r.db)
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
