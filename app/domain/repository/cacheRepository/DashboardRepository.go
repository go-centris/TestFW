package cacheRepository

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/dto"
	repository "stncCms/app/domain/repository/dbRepository"
	"time"
)

// KurbanRepo struct
type DashboardRepo struct {
	db *gorm.DB
}

// KurbanRepositoryInit initial
func DashboardRepositoryInit(db *gorm.DB) *DashboardRepo {
	return &DashboardRepo{db}
}

// TotalSacrife  toplam kesilen kurban
func (r *DashboardRepo) TotalSacrife(returnValue *int64) {
	var totalSacrife int64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.TotalSacrife(&totalSacrife)
	*returnValue = totalSacrife
}

// TotalPrice   genel toplam kesilen kurban parasi
func (r *DashboardRepo) TotalPrice(returnValue *float64) {
	var totalPrice float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.TotalPrice(&totalPrice)
	*returnValue = totalPrice
}

// SacrifeSharedPriceTotal   sadece hisseli kurban miktari
func (r *DashboardRepo) SacrifeSharedPriceTotal(returnValue *float64) {
	var total float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SacrifeSharedPriceTotal(&total)
	*returnValue = total
}

// RemainingDebt Toplam Kalan Borc
func (r *DashboardRepo) RemainingDebt(returnValue *float64) {
	var debt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.RemainingDebt(&debt)
	*returnValue = debt
}

// SharedSacrifeTotal   Hisseli Kesilen Toplam Kurban
func (r *DashboardRepo) SharedSacrifeTotal(returnValue *int64) {
	var sharedSacrifeTotal int64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SharedSacrifeTotal(&sharedSacrifeTotal)
	*returnValue = sharedSacrifeTotal
}

// SharedSacrifeRemainingDebt sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SharedSacrifeRemainingDebt(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SharedSacrifeRemainingDebt(&remainingDebt)
	*returnValue = remainingDebt
}

// ShareSacrifeCount2021  sadece hisseli kurban miktari yil 2021
func (r *DashboardRepo) ShareSacrifeCount2021(returnValue *int64) {
	var sharedSacrifeTotal int64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.ShareSacrifeCount2021(&sharedSacrifeTotal)
	*returnValue = sharedSacrifeTotal
}

// ShareSacrifeCount2022  sadece hisseli kurban miktari yil 2022
func (r *DashboardRepo) ShareSacrifeCount2022(returnValue *int64) {
	var sharedSacrifeTotal int64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.ShareSacrifeCount2022(&sharedSacrifeTotal)
	*returnValue = sharedSacrifeTotal
}

// ShareSacrifeCount2023  sadece hisseli kurban miktari yil 2023
func (r *DashboardRepo) ShareSacrifeCount2023(returnValue *int64) {
	var sharedSacrifeTotal int64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.ShareSacrifeCount2023(&sharedSacrifeTotal)
	*returnValue = sharedSacrifeTotal
}

// SharedSacrifeRemainingDebt2021 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SharedSacrifeRemainingDebt2021(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SharedSacrifeRemainingDebt2021(&remainingDebt)
	*returnValue = remainingDebt
}

// SharedSacrifeRemainingDebt2022 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SharedSacrifeRemainingDebt2022(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SharedSacrifeRemainingDebt2022(&remainingDebt)
	*returnValue = remainingDebt
}

// SharedSacrifeRemainingDebt2023 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SharedSacrifeRemainingDebt2023(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SharedSacrifeRemainingDebt2023(&remainingDebt)
	*returnValue = remainingDebt
}

// SacrifeSharedPriceTotal2021 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SacrifeSharedPriceTotal2021(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SacrifeSharedPriceTotal2021(&remainingDebt)
	*returnValue = remainingDebt
}

// SacrifeSharedPriceTotal2022 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SacrifeSharedPriceTotal2022(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SacrifeSharedPriceTotal2022(&remainingDebt)
	*returnValue = remainingDebt
}

// SacrifeSharedPriceTotal2023 sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SacrifeSharedPriceTotal2023(returnValue *float64) {
	var remainingDebt float64
	repo := repository.DashboardRepositoryInit(r.db)
	repo.SacrifeSharedPriceTotal2023(&remainingDebt)
	*returnValue = remainingDebt
}

// CharitableWhoAddedMostSacrife  En cok kurban kestiren hayirsever
func (r *DashboardRepo) CharitableWhoAddedMostSacrife() (*dto.CharitableWhoAddedMostSacrife, error) {
	var data *dto.CharitableWhoAddedMostSacrife
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = charitableWhoAddedMostSacrife(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "CharitableWhoAddedMostSacrife_"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = charitableWhoAddedMostSacrife(r.db)
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

// charitableWhoAddedMostSacrife En cok kurban kestiren hayirsever
func charitableWhoAddedMostSacrife(db *gorm.DB) (*dto.CharitableWhoAddedMostSacrife, error) {
	repo := repository.DashboardRepositoryInit(db)
	data, _ := repo.CharitableWhoAddedMostSacrife()
	return data, nil
}

// UsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
func (r *DashboardRepo) UsersWhoAddedMostSacrifeAndBranch() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var data *dto.UsersWhoAddedMostSacrifeAndBranch
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = usersWhoAddedMostSacrifeAndBranch(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "usersWhoAddedMostSacrifeAndBranch"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = usersWhoAddedMostSacrifeAndBranch(r.db)
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

// usersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
func usersWhoAddedMostSacrifeAndBranch(db *gorm.DB) (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	repo := repository.DashboardRepositoryInit(db)
	data, _ := repo.UsersWhoAddedMostSacrifeAndBranch()
	return data, nil
}

// UsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
func (r *DashboardRepo) UsersWhoAddedMostSacrifeAndUser() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var data *dto.UsersWhoAddedMostSacrifeAndBranch
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = usersWhoAddedMostSacrifeAndUser(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "usersWhoAddedMostSacrifeAndUser"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = usersWhoAddedMostSacrifeAndUser(r.db)
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

// usersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen hocamiz
func usersWhoAddedMostSacrifeAndUser(db *gorm.DB) (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	repo := repository.DashboardRepositoryInit(db)
	data, _ := repo.UsersWhoAddedMostSacrifeAndUser()
	return data, nil
}
