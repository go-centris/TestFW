package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	repository "stncCms/app/domain/repository/dbRepository"
	"time"

	"github.com/jinzhu/gorm"
)

// GruplarRepo struct
type GruplarRepo struct {
	db *gorm.DB
}

// GruplarRepositoryInit initial
func GruplarRepositoryInit(db *gorm.DB) *GruplarRepo {
	return &GruplarRepo{db}
}

//GruplarRepo implements the repository.KurbanRepository interface
// var _ interfaces.PostAppInterface = &GruplarRepo{}

// GetByID get data
func (r *GruplarRepo) GetByID(id uint64) (*entity.SacrificeGruplar, error) {
	var data *entity.SacrificeGruplar

	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDGroups(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByIDGroups_" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDGroups(r.db, id)
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

func getByIDGroups(db *gorm.DB, id uint64) (*entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// GetByIDAllRelations get //TODO:  data hayvanı atanmış olnaları getirir
func getByIDAllRelationsGroups(db *gorm.DB, id uint64) (*entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByIDAllRelations(id)
	return datas, nil
}

// GetByIDAllRelations get //TODO:  data hayvanı atanmış olnaları getirir
func (r *GruplarRepo) GetByIDAllRelations(id uint64) (*entity.SacrificeGruplar, error) {
	var data *entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDAllRelationsGroups(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByIDAllRelationsGroups_" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDAllRelationsGroups(r.db, id)
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

// GetByIDAllRelations get //TODO:  data hayvanı atanmamış olanları getirir
func getByIDAllRelationsHayvanOlmayanlarGroups(db *gorm.DB, id uint64) (*entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByIDAllRelationsHayvanOlmayanlar(id)
	return datas, nil
}

// GetByIDAllRelations get //TODO:  data hayvanı atanmamış olanları getirir
func (r *GruplarRepo) GetByIDAllRelationsHayvanOlmayanlar(id uint64) (*entity.SacrificeGruplar, error) {
	var data *entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDAllRelationsHayvanOlmayanlarGroups(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetByIDAllRelationsHayvanOlmayanlarGroups_" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDAllRelationsHayvanOlmayanlarGroups(r.db, id)
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

// GetByAllRelationsLimit listeleme
func getByAllRelationsLimitGroups(db *gorm.DB) ([]dto.GruplarExcelandIndex, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByAllRelationsLimit()
	return datas, nil
}

// GetByAllRelationsLimit listeleme
func (r *GruplarRepo) GetByAllRelationsLimit() ([]dto.GruplarExcelandIndex, error) {
	var data []dto.GruplarExcelandIndex
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByAllRelationsLimitGroups(r.db)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByAllRelationsLimitGroups_"

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByAllRelationsLimitGroups(r.db)
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

// GetByAllRelations listeleme
func getByAllRelationsGroup(db *gorm.DB, year string) ([]dto.GruplarExcelandIndex, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByAllRelations(year)
	return datas, nil

}

// GetByAllRelations listeleme
func (r *GruplarRepo) GetByAllRelations(year string) ([]dto.GruplarExcelandIndex, error) {
	var data []dto.GruplarExcelandIndex
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByAllRelationsGroup(r.db, year)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetByAllRelationsGroup_" + year

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByAllRelationsGroup(r.db, year)
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

// GetByAllRelationsByID meditaor
func getByAllRelationsByIDGroup(db *gorm.DB, id uint64, year string) (*dto.GruplarExcelandIndex, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetByAllRelationsByID(id, year)
	return datas, nil
}

// GetByAllRelationsByID listeleme
func (r *GruplarRepo) GetByAllRelationsByID(id uint64, year string) (*dto.GruplarExcelandIndex, error) {
	var data *dto.GruplarExcelandIndex
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByAllRelationsByIDGroup(r.db, id, year)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByAllRelationsByIDGroup_" + stnccollection.Uint64toString(id) + year

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByAllRelationsByIDGroup(r.db, id, year)
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

func GetAllStatusFindGroup(db *gorm.DB, durum int) ([]entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetAllStatusFind(durum)
	return datas, nil
}

func (r *GruplarRepo) GetAllStatusFind(durum int) ([]entity.SacrificeGruplar, error) {
	var data []entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = GetAllStatusFindGroup(r.db, durum)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetAllStatusFindGroup_" + stnccollection.IntToString(durum)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = GetAllStatusFindGroup(r.db, durum)
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

func GetAllStatusFindAndAgirlikTipiGroup(db *gorm.DB, durum int, agirlikTipi int) ([]entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetAllStatusFindAndAgirlikTipi(durum, agirlikTipi)
	return datas, nil
}

// GetAllStatusFind all data -- TODO: buradaki yil olayi degisecek
func (r *GruplarRepo) GetAllStatusFindAndAgirlikTipi(durum int, agirlikTipi int) ([]entity.SacrificeGruplar, error) {
	var data []entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = GetAllStatusFindAndAgirlikTipiGroup(r.db, durum, agirlikTipi)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetAllStatusFindAndAgirlikTipiGroup_" + stnccollection.IntToString(durum) + stnccollection.IntToString(agirlikTipi)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = GetAllStatusFindAndAgirlikTipiGroup(r.db, durum, agirlikTipi)
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

func getAllGroups(db *gorm.DB) ([]entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetAll()
	return datas, nil
}

func (r *GruplarRepo) GetAll() ([]entity.SacrificeGruplar, error) {
	var data []entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getAllGroups(r.db)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getAllGroups"

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getAllGroups(r.db)
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
func getAllpGroups(db *gorm.DB, postsPerPage int, offset int) ([]entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetAllPagination(postsPerPage, offset)
	return datas, nil
}

func (r *GruplarRepo) GetAllPagination(postsPerPage int, offset int) ([]entity.SacrificeGruplar, error) {
	var data []entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getAllpGroups(r.db, postsPerPage, offset)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getAllpGroups_" + stnccollection.IntToString(postsPerPage) + stnccollection.IntToString(offset)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getAllpGroups(r.db, postsPerPage, offset)
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

func kurbanFiyatiGroup(db *gorm.DB, kurbanID uint64) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.KurbanFiyati(kurbanID)
}

// KurbanFiyati kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) KurbanFiyati(kurbanID uint64) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = kurbanFiyatiGroup(r.db, kurbanID)
	} else {
		redisClient := cache.RedisDBInit()

		key := "kurbanFiyatiGroup_" + stnccollection.Uint64toString(kurbanID)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = kurbanFiyatiGroup(r.db, kurbanID)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) GetGrupIDTotalPayment(grupID uint64) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = GetGrupIDTotalPayment(r.db, grupID)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetGrupIDTotalPayment_" + stnccollection.Uint64toString(grupID)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = GetGrupIDTotalPayment(r.db, grupID)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

func GetGrupIDTotalPayment(db *gorm.DB, grupID uint64) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.GetGrupIDTotalPayment(grupID)
}

func getGrupTotalPayment(db *gorm.DB) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.GetGrupTotalPayment()
}

func (r *GruplarRepo) GetGrupTotalPayment() float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = getGrupTotalPayment(r.db)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getGrupTotalPayment"

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = getGrupTotalPayment(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func grupKalanBorclar(db *gorm.DB, grupID uint64) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.GrupKalanBorclar(grupID)
}

func (r *GruplarRepo) GrupKalanBorclar(grupID uint64) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = grupKalanBorclar(r.db, grupID)
	} else {
		redisClient := cache.RedisDBInit()

		key := "grupKalanBorclar" + stnccollection.Uint64toString(grupID)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = grupKalanBorclar(r.db, grupID)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

// kasaBorcu kurban id ye gore odenen miktar toplamı
func kasaBorcu(db *gorm.DB, grupID uint64) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.KasaBorcu(grupID)
}

// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) KasaBorcu(grupID uint64) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = kasaBorcu(r.db, grupID)
	} else {
		redisClient := cache.RedisDBInit()

		key := "KasaBorcu"

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = kasaBorcu(r.db, grupID)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

// SatisFiyatTuru kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) SatisFiyatTuru(hayvanBilgisiID uint64) int {
	var data int
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = satisFiyatTuru(r.db, hayvanBilgisiID)
	} else {
		redisClient := cache.RedisDBInit()

		key := "satisFiyatTuru" + stnccollection.Uint64toString(hayvanBilgisiID)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = satisFiyatTuru(r.db, hayvanBilgisiID)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

func satisFiyatTuru(db *gorm.DB, hayvanBilgisiID uint64) int {
	repo := repository.GruplarRepositoryInit(db)
	return repo.SatisFiyatTuru(hayvanBilgisiID)
}

//func toplamOdemeler(db *gorm.DB) float64 {
//	repo := repository.GruplarRepositoryInit(db)
//	return repo.ToplamOdemeler()
//}
//
//func (r *GruplarRepo) ToplamOdemeler() float64 {
//	var data float64
//	access := repository.OptionRepositoryInit(r.db)
//	cacheControl := access.GetOption("cache_open_close")
//
//	if cacheControl == "false" {
//		data = toplamOdemeler(r.db)
//	} else {
//		redisClient := cache.RedisDBInit()
//
//		key := "toplamOdemeler"
//
//		cachedProducts, err := redisClient.GetKey(key)
//
//		if err != nil {
//			data = toplamOdemeler(r.db)
//			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
//			fmt.Println("key olustur")
//			if err != nil {
//				fmt.Println("hata baş")
//			}
//			return data
//		}
//		err = json.Unmarshal(cachedProducts, &data)
//		if err != nil {
//			fmt.Println("hata son")
//		}
//	}
//	return data
//}

//-************************-----burada---------------------------*****************************

func toplamOdemelerForYear(db *gorm.DB, year string) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.ToplamOdemelerForYear(year)
}
func (r *GruplarRepo) ToplamOdemelerForYear(year string) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = toplamOdemelerForYear(r.db, year)
	} else {
		redisClient := cache.RedisDBInit()

		key := "toplamOdemelerForYear" + year

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = toplamOdemelerForYear(r.db, year)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}

//-************************-----burada---------------------------*****************************

// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func kalanBorclarForYear(db *gorm.DB, year string) float64 {
	repo := repository.GruplarRepositoryInit(db)
	return repo.KalanBorclarForYear(year)
}

// KalanBorclarForYear
func (r *GruplarRepo) KalanBorclarForYear(year string) float64 {
	var data float64
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data = kalanBorclarForYear(r.db, year)
	} else {
		redisClient := cache.RedisDBInit()

		key := "kalanBorclarForYear" + year

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data = kalanBorclarForYear(r.db, year)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data
}
