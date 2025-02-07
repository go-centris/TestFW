package dbRepository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//var KurbanBorcDurumIlkEklenenFiyat string = stnccollection.IntToString(entity.KurbanBorcDurumIlkEklenenFiyat)

// DashboardRepo struct
type DashboardRepo struct {
	db *gorm.DB
}

// DashboardRepositoryInit initial
func DashboardRepositoryInit(db *gorm.DB) *DashboardRepo {
	return &DashboardRepo{db}
}

//// KalanUcret kurbanın kalan ücret bilgisi
//func (r *DashboardRepo) KalanUcret(kurbanID uint64) float64 {
//	var result float64
//	row := r.db.Debug().Table(entity.KurbanTableName).Select("alacak").Where("id = ?", kurbanID).Row()
//	row.Scan(&result)
//	return result
//}

//// GetKurbanKisiVarmi grupta kac kisi var
//func (r *DashboardRepo) GetKurbanKisiVarmi(kurbanID uint64, kisiID *int) {
//	var kisi int
//	row := r.db.Debug().Table(entity.KurbanTableName).Select("kisi_id").Where("id = ? ", kurbanID).Row()
//	row.Scan(&kisi)
//	*kisiID = kisi
//}

// TotalSacrife toplam kesilen kurban adeti
func (r *DashboardRepo) TotalSacrife(returnValue *int64) {
	var table entity.SacrificeKurbanlar
	var count int64
	r.db.Debug().Model(table).Where("durum<>1 AND durum<>2 AND durum<>3 AND durum<>9  ").Count(&count)
	*returnValue = count
}

// TotalPrice genel toplam kesilen kurban parasi
func (r *DashboardRepo) TotalPrice(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(bakiye) as totalPrice").Where("borc_durum<>2 AND borc_durum<>1 ").Row().Scan(&data)
	*returnValue = data
}

// SacrifeSharedPriceTotal sadece hisseli kurban paRASI
func (r *DashboardRepo) SacrifeSharedPriceTotal(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(bakiye) as total").Where("borc_durum<>2 AND borc_durum<>1 AND kurban_turu=9 ").Row().Scan(&data)
	*returnValue = data
}

// RemainingDebt Toplam Kalan Borc
func (r *DashboardRepo) RemainingDebt(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(alacak) as total").Where("borc_durum=3 ").Row().Scan(&data)
	*returnValue = data
}

// SharedSacrifeTotal   Hisseli Kesilen Toplam Kurban
func (r *DashboardRepo) SharedSacrifeTotal(returnValue *int64) {
	var table entity.SacrificeKurbanlar
	var count int64
	r.db.Debug().Model(table).Where("durum<>1 AND durum<>2 AND durum<>3 AND durum<>9").Where("kurban_turu=8 OR kurban_turu=9 ").Count(&count)
	*returnValue = count
}

// SharedSacrifeRemainingDebt sadece hisseli  Kalan toplam Borc
func (r *DashboardRepo) SharedSacrifeRemainingDebt(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(alacak) as total").Where("borc_durum=3 ").Where("kurban_turu=8 OR kurban_turu=9").Row().Scan(&data)
	*returnValue = data
}

// ShareSacrifeCount2021  sadece hisseli kurban miktari yil 2021
func (r *DashboardRepo) ShareSacrifeCount2021(returnValue *int64) {
	var table entity.SacrificeKurbanlar
	var data int64
	//r.db.Debug().Model(table).Select("count(id) as total").Where(" borc_durum<>2 AND borc_durum<>1 AND kurban_turu=9 AND kurban_bayrami_yili=2021").Count(&data)
	r.db.Debug().Model(table).Select("count(id) as total").Where(" borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2021").Count(&data)
	*returnValue = data
}

// ShareSacrifeCount2022  sadece hisseli kurban miktari yil 2021
func (r *DashboardRepo) ShareSacrifeCount2022(returnValue *int64) {
	var table entity.SacrificeKurbanlar
	var data int64
	r.db.Debug().Model(table).Select("count(id) as total").Where(" borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2022").Count(&data)
	*returnValue = data
}

// ShareSacrifeCount2023  sadece hisseli kurban miktari yil 2021
func (r *DashboardRepo) ShareSacrifeCount2023(returnValue *int64) {
	var table entity.SacrificeKurbanlar
	var data int64
	r.db.Debug().Model(table).Select("count(id) as total").Where(" borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2023").Count(&data)
	*returnValue = data
}

// SharedSacrifeRemainingDebt2021 sadece hisseli  Kalan toplam Borc 2021
func (r *DashboardRepo) SharedSacrifeRemainingDebt2021(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(alacak) as total").Where("borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2021").Row().Scan(&data)
	*returnValue = data
}

// SharedSacrifeRemainingDebt2022 sadece hisseli  Kalan toplam Borc 2022
func (r *DashboardRepo) SharedSacrifeRemainingDebt2022(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(alacak) as total").Where("borc_durum<>2 AND borc_durum<>1  AND kurban_bayrami_yili=2022").Row().Scan(&data)
	*returnValue = data
}

// SharedSacrifeRemainingDebt2023 sadece hisseli  Kalan toplam Borc 2023
func (r *DashboardRepo) SharedSacrifeRemainingDebt2023(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(alacak) as total").Where("borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2023").Row().Scan(&data)
	*returnValue = data
}

// SacrifeSharedPriceTotal2021 sadece hisseli kurban paRASI 2021
func (r *DashboardRepo) SacrifeSharedPriceTotal2021(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(bakiye) as total").Where("borc_durum<>2 AND borc_durum<>1 AND kurban_bayrami_yili=2021").Row().Scan(&data)
	*returnValue = data
}

// SacrifeSharedPriceTotal2022 sadece hisseli kurban paRASI 2021
func (r *DashboardRepo) SacrifeSharedPriceTotal2022(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(bakiye) as total").Where("borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2022").Row().Scan(&data)
	*returnValue = data
}

// SacrifeSharedPriceTotal2023 sadece hisseli kurban paRASI 2021
func (r *DashboardRepo) SacrifeSharedPriceTotal2023(returnValue *float64) {
	var table entity.SacrificeKurbanlar
	var data float64
	r.db.Debug().Model(table).Select("sum(bakiye) as total").Where("borc_durum<>2 AND borc_durum<>1 AND  kurban_bayrami_yili=2023").Row().Scan(&data)
	*returnValue = data
}

//// GetByID2 get data for id
//func (r *DashboardRepo) GetByID2(id uint64) (*entity.Branches, error) {
//	var cat entity.Branches
//	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
//
//	//	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
//
//	if err != nil {
//		return nil, errors.New("database error, please try again")
//	}
//	if gorm.IsRecordNotFoundError(err) {
//		return nil, errors.New("post not found")
//
//	}
//	return &cat, nil
//}

// CharitableWhoAddedMostSacrife  En cok kurban kestiren hayirsever
func (r *DashboardRepo) CharitableWhoAddedMostSacrife() (*dto.CharitableWhoAddedMostSacrife, error) {
	var charitableWhoAddedMostSacrife dto.CharitableWhoAddedMostSacrife
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(`sacrifice_kurbanlar.user_id,ki.id as kisi_id, COUNT(*) AS counts,ki.ad_soyad AS name_lastname`)
	query = query.Joins("INNER JOIN sacrifice_kisiler AS ki ON ki.id=sacrifice_kurbanlar.kisi_id")
	query = query.Group("sacrifice_kurbanlar.user_id,ki.ad_soyad,ki.id ")
	query = query.Having("ki.ad_soyad <> 'Boş GRUP, Kişi kaydı için tıklayınız'")
	query = query.Order(" counts DESC")
	err = query.Take(&charitableWhoAddedMostSacrife).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &charitableWhoAddedMostSacrife, nil
}

// UsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
func (r *DashboardRepo) UsersWhoAddedMostSacrifeAndBranch() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var usersWhoAddedMostSacrifeAndBranch dto.UsersWhoAddedMostSacrifeAndBranch
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(
		`br.id AS branch_id,COUNT(br) AS counts,br.title AS branch `)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches AS br ON br.id=users.branch_id")
	query = query.Group("br.id")
	query = query.Order("counts DESC")
	err = query.Take(&usersWhoAddedMostSacrifeAndBranch).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &usersWhoAddedMostSacrifeAndBranch, nil
}

// UsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
func (r *DashboardRepo) UsersWhoAddedMostSacrifeAndUser() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var usersWhoAddedMostSacrifeAndBranch dto.UsersWhoAddedMostSacrifeAndBranch
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(
		`sacrifice_kurbanlar.user_id AS sacrife_user_id, COUNT(users) AS counts,users.first_name,
     users.last_name,users.email,br.title AS branch`)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches AS br ON br.id=users.branch_id")
	query = query.Group("users.id,sacrife_user_id,br.id")
	query = query.Order(" counts DESC")
	err = query.Take(&usersWhoAddedMostSacrifeAndBranch).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &usersWhoAddedMostSacrifeAndBranch, nil
}
