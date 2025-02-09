package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stnchelper"
	repository "stncCms/app/domain/repository/dbRepository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// KisilerRepo struct
type KisilerRepo struct {
	db *gorm.DB
}

// KisilerRepositoryInit initial
func KisilerRepositoryInit(db *gorm.DB) *KisilerRepo {
	return &KisilerRepo{db}
}

type JsonResponse struct {
	Data   []entity.Kisiler `json:"data"`
	Source string           `json:"source"`
}

func getByID(db *gorm.DB, id uint64) (*entity.Kisiler, error) {
	repo := repository.KisilerRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// getByIDRel get data
func getByIDRel(db *gorm.DB, id uint64) (*entity.Kisiler, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.GetByIDRel(id)
	return data, nil
}

// getAll all data
// TODO: pointer look
func getByReferansID(db *gorm.DB, id uint64) (*entity.Kisiler, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.GetByReferansID(id)
	return data, nil
}

// getAll all data
func getAll(db *gorm.DB) ([]entity.Kisiler, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// listDataTable upate data
func listDataTable(c *gin.Context, db *gorm.DB) (dto.KisilerDataTableResult, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.ListDataTable(c)
	return data, nil
}
func getAllP(db *gorm.DB, postsPerPage int, offset int) ([]entity.Kisiler, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.GetAllPagination(postsPerPage, offset)
	return data, nil
}

// search
func searchData(search string, db *gorm.DB) ([]entity.KisilerDTO, error) {
	repo := repository.KisilerRepositoryInit(db)
	data, _ := repo.Search(search)
	return data, nil
}

// GetByID get data
func (r *KisilerRepo) GetByID(id uint64) (*entity.Kisiler, error) {
	// data := entity.Kisiler{ID: 1}
	// return &data, nil
	var data *entity.Kisiler
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByID(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "PersonsGetByID" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByID(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

// GetByIDRel get data
func (r *KisilerRepo) GetByIDRel(id uint64) (*entity.Kisiler, error) {
	var data *entity.Kisiler
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDRel(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetByIDRelPersons_" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDRel(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

// GetByReferans referans data
// TODO: tam olarak verecek belli değil kullanımda değil
func (r *KisilerRepo) GetByReferansID(id uint64) (*entity.Kisiler, error) {
	var data *entity.Kisiler
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByReferansID(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByReferansID_Persons" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByReferansID(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

func (r *KisilerRepo) GetAll() ([]entity.Kisiler, error) {

	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	var data []entity.Kisiler
	if cacheControl == "false" {
		data, _ = getAll(r.db)
	} else {

		redisClient := cache.RedisDBInit()

		key := "personsGetAll"

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getAll(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

func (r *KisilerRepo) ListDataTable(c *gin.Context) (dto.KisilerDataTableResult, error) {

	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	data := dto.KisilerDataTableResult{}

	if cacheControl == "false" {
		data, _ = listDataTable(c, r.db)
	} else {

		redisClient := cache.RedisDBInit()
		search := c.QueryMap("search")

		key := "persons_" + stnccollection.IntToString(stnchelper.QueryOffset(c)) + string(search["value"])
		fmt.Println("key")
		fmt.Println(key)
		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = listDataTable(c, r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

// GetAllPagination pagination all data
func (r *KisilerRepo) GetAllPagination(postsPerPage int, offset int) ([]entity.Kisiler, error) {

	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	var data []entity.Kisiler
	if cacheControl == "false" {
		data, _ = getAllP(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPagination_persons_" + stnccollection.IntToString(postsPerPage) + "_" + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllP(r.db, postsPerPage, offset)
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

// Search   data
func (r *KisilerRepo) Search(search string) ([]entity.KisilerDTO, error) {

	var data []entity.KisilerDTO
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = searchData(search, r.db)
	} else {

		redisClient := cache.RedisDBInit()
		key := "personSearch_" + search
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = searchData(search, r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			// fmt.Println("key olustur")
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

// func (r *KisilerRepo) GetAllcache() ([]entity.Kisiler, error) {
// 	data := []entity.Kisiler{}
// 	// var data []entity.Kisiler
// 	// var data []entity.Kisiler

// 	redisClient := cache.RedisDBInit()
// 	key := "data"
// 	cachedProducts, err := redisClient.GetKey(key)

// 	if err != nil {
// 		repo := repository.KisilerRepositoryInit(r.db)
// 		data, _ := repo.GetAll()
// 		err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
// 		if err != nil {
// 			return nil, err
// 		}
// 		return data, nil
// 	}

// 	err = json.Unmarshal(cachedProducts, data)

// 	if err != nil {
// 		fmt.Println("hata son")
// 		return nil, err
// 	}

// 	return data, nil
// }

//-------------------------------------------------------------------

//Search   data -- bunda map donusunde status ve null gibi degerler var ondan dolayi beklesin silme

/*
func (r *KisilerRepo) Search(search string) (map[string]string, []entity.KisilerDTO, error) {
	var data []entity.KisilerDTO
	var err error
	var count int64
	dbdriver := os.Getenv("DB_DRIVER")

	if dbdriver == "mysql" {
		err = r.db.Debug().Table(dto.KisilerTableName).Where("ad_soyad LIKE ?", "%"+search+"%").Find(&data).Error
		r.db.Debug().Table(dto.KisilerTableName).Where("ad_soyad LIKE ?", "%"+search+"%").Model(&data).Count(&count)
	} else if dbdriver == "postgres" {
		err = r.db.Debug().Table(dto.KisilerTableName).Where("ad_soyad ILIKE ?", "%"+search+"%").Find(&data).Error
		r.db.Debug().Table(dto.KisilerTableName).Where("ad_soyad ILIKE ?", "%"+search+"%").Model(&data).Count(&count)
	}

	if err != nil {
		fmt.Println("err nil")
		return map[string]string{"status": "error"}, nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("null nil")
		return map[string]string{"status": "error"}, nil, errors.New("post not found")
	}

	if count == 0 {
		fmt.Println("data nil")
		mapD := map[string]string{"status": "not found", "html": "<span></span>"}
		//mapB, _ := json.Marshal(mapD)
		//fmt.Println(string(mapB))
		return mapD, nil, err
	} else {
		return map[string]string{"status": "ok"}, data, nil
	}

}
*/

// func (r *KisilerRepo) queryOrder(c *gin.Context) string {
// 	columnMap := map[string]string{
// 		"0": "id",
// 		"1": "ad_soyad",
// 		"2": "telefon",
// 		"3": "adres",
// 	}

// 	column := c.DefaultQuery("order[0][column]", "1")
// 	dir := c.DefaultQuery("order[0][dir]", "desc")
// 	orderString := columnMap[column] + " " + dir

// 	return orderString
// }
