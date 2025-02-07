package dbRepository

import (
	"errors"
	"fmt"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	"strings"

	"github.com/jinzhu/gorm"
)

// GruplarRepo struct
type GruplarRepo struct {
	db *gorm.DB
}

var gruplarDurumKurbanBayramiDisiKesim string = stnccollection.IntToString(entity.GruplarDurumKurbanBayramiDisiKesim)
var kurbanDurumKurbanKesildiTeslimEdildi string = stnccollection.IntToString(entity.KurbanDurumKurbanKesildiTeslimEdildi) //durum <> 7
var kurbanBorcDurumKasaBorcluDurumda string = stnccollection.IntToString(entity.KurbanBorcDurumKasaBorcluDurumda)         //borc_durum <> 3
var kurbanBorcDurumIlkEklenenFiyat string = stnccollection.IntToString(entity.KurbanBorcDurumIlkEklenenFiyat)
var kurbanBorcDurumHesapKapandi string = stnccollection.IntToString(entity.KurbanBorcDurumHesapKapandi)

// GruplarRepositoryInit initial
func GruplarRepositoryInit(db *gorm.DB) *GruplarRepo {
	return &GruplarRepo{db}
}

//GruplarRepo implements the repository.KurbanRepository interface
// var _ interfaces.PostAppInterface = &GruplarRepo{}

// Save data
func (r *GruplarRepo) Save(data *entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

// Update upate data
func (r *GruplarRepo) Update(post *entity.SacrificeGruplar) (*entity.SacrificeGruplar, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&post).Error
	//db.Table("libraries").Where("id = ?", id).Update(postData)

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return post, nil
}

// Delete data
func (r *GruplarRepo) Delete(id uint64) error {
	var post entity.SacrificeGruplar
	err := r.db.Debug().Where("id = ?", id).Delete(&post).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

// GetByID get data
func (r *GruplarRepo) GetByID(id uint64) (*entity.SacrificeGruplar, error) {
	var post entity.SacrificeGruplar
	err := r.db.Debug().Where("id = ? and durum <> 1  ", id).Preload("Kurban").Take(&post).Error
	fmt.Printf("%+v\n", post)
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &post, nil
}

// GetByIDAllRelations get //TODO:  data hayvanı atanmış olnaları getirir
func (r *GruplarRepo) GetByIDAllRelations(id uint64) (*entity.SacrificeGruplar, error) {
	var gruplarList entity.SacrificeGruplar
	err := r.db.Debug().Where("id = ? and  durum <> "+gruplarDurumKurbanBayramiDisiKesim+" and hayvan_bilgisi_id <> 0 ", id).Preload("Kurban").Take(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &gruplarList, nil
}

// GetByIDAllRelations get //TODO:  data hayvanı atanmamış olanları getirir
func (r *GruplarRepo) GetByIDAllRelationsHayvanOlmayanlar(id uint64) (*entity.SacrificeGruplar, error) {
	var gruplarList entity.SacrificeGruplar
	// var err error
	err := r.db.Debug().Where("id = ? and  durum <> "+gruplarDurumKurbanBayramiDisiKesim+" and hayvan_bilgisi_id = 0 ", id).Preload("Kurban").Take(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &gruplarList, nil
}

// GetByAllRelations listeleme
func (r *GruplarRepo) GetByAllRelationsLimit() ([]dto.GruplarExcelandIndex, error) {
	var gruplarList []dto.GruplarExcelandIndex

	// err = r.db.Debug().Where("  durum <> 1  ").Preload("Kurban").Find(&gruplarList).Error //entity.gruplara gore
	// err := r.db.Debug().Table("sacrifice_gruplar").Where("  durum <> 1  ").Limit(3).Order("kesim_sira_no asc").Find(&gruplarList).Error
	err := r.db.Debug().Table("sacrifice_gruplar").Where("  durum <> " + gruplarDurumKurbanBayramiDisiKesim).Order("kesim_sira_no asc").Find(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return gruplarList, nil
}

// GetByAllRelations listeleme
func (r *GruplarRepo) GetByAllRelations(year string) ([]dto.GruplarExcelandIndex, error) {
	var gruplarList []dto.GruplarExcelandIndex

	// err = r.db.Debug().Where("  durum <> 1  ").Preload("Kurban").Find(&gruplarList).Error //entity.gruplara gore
	err := r.db.Debug().Table("sacrifice_gruplar").Where("  kurban_bayrami_yili = ? AND  durum <> "+gruplarDurumKurbanBayramiDisiKesim, year).Order("kesim_sira_no asc").Find(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return gruplarList, nil
}

// TODO: burayı iptal et kullanan yok sanırım
// GetAllStatusFind get data
func (r *GruplarRepo) GetAllStatusFind(durum int) ([]entity.SacrificeGruplar, error) {
	var gruplarList []entity.SacrificeGruplar
	err := r.db.Debug().Where("  durum = ?  ", durum).Find(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return gruplarList, nil
}

// GetAllStatusFind all data
// TODO: buradaki yil olayi degisecek
func (r *GruplarRepo) GetAllStatusFindAndAgirlikTipi(durum int, agirlikTipi int) ([]entity.SacrificeGruplar, error) {
	var datas []entity.SacrificeGruplar
	err := r.db.Debug().Where("durum = ? and agirlik_tipi= ? and kurban_bayrami_yili <> 2023", durum, agirlikTipi).Preload("Kurban").Order("id asc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

// GetAll all data
func (r *GruplarRepo) GetAll() ([]entity.SacrificeGruplar, error) {
	var rows []entity.SacrificeGruplar
	err := r.db.Debug().Where("  durum <> " + gruplarDurumKurbanBayramiDisiKesim).Order("kesim_sira_no asc").Find(&rows).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return rows, nil
}

// GetAllPagination pagination all data
func (r *GruplarRepo) GetAllPagination(postsPerPage int, offset int) ([]entity.SacrificeGruplar, error) {
	var posts []entity.SacrificeGruplar
	// var err error
	err := r.db.Debug().Where(" durum <> " + gruplarDurumKurbanBayramiDisiKesim).Limit(postsPerPage).Offset(offset).Order("created_at desc").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return posts, nil
}

// TODO: durum sogrulari entity e gore olacak
// KurbanFiyati kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) KurbanFiyati(kurbanID uint64) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("kurban_fiyati").Where("id = ? AND borc_durum ="+kurbanBorcDurumIlkEklenenFiyat, kurbanID).Row()
	row.Scan(&result)
	return result
}

// TODO: durum sogrulari entity e gore olacak sorguyu kisalt alt sanira insin
// ToplamOdemeler grup  id ye gore odenen miktar toplamı
func (r *GruplarRepo) GetGrupIDTotalPayment(GrupID uint64) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("sum(bakiye) as total").Where("grup_id = ? AND borc_durum <> "+kurbanBorcDurumKasaBorcluDurumda+" AND durum <>  "+kurbanDurumKurbanKesildiTeslimEdildi, GrupID).Row()
	row.Scan(&result)
	return result
}

// TODO: durum sogrulari entity e gore olacak
func (r *GruplarRepo) GetGrupTotalPayment() float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("sum(bakiye) as total").Where(" borc_durum <> " + kurbanBorcDurumKasaBorcluDurumda + " AND durum <>  " + kurbanDurumKurbanKesildiTeslimEdildi).Row()
	row.Scan(&result)
	return result
}

// TODO: durum sogrulari entity e gore olacak
// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) GrupKalanBorclar(GrupID uint64) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("sum(alacak) as total").Where("grup_id = ? AND borc_durum <> "+kurbanBorcDurumKasaBorcluDurumda, GrupID).Row()
	row.Scan(&result)
	return result
}

// TODO: durum sogrulari entity e gore olacak
// ToplamOdemeler kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) KasaBorcu(GrupID uint64) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("sum(kasa_borcu) as total").Where("grup_id = ? AND borc_durum = "+kurbanBorcDurumKasaBorcluDurumda+"   AND durum <> "+kurbanBorcDurumHesapKapandi, GrupID).Row()

	row.Scan(&result)
	return result
}

// SatisFiyatTuru kurban id ye gore odenen miktar toplamı
func (r *GruplarRepo) SatisFiyatTuru(hayvanBilgisiID uint64) int {
	var result int
	row := r.db.Debug().Table(entity.GruplarTableName).Select("satis_fiyat_turu").Where("hayvan_bilgisi_id = ? ", hayvanBilgisiID).Row()
	row.Scan(&result)
	return result
}

// Count fat
func (r *GruplarRepo) Count(postTotalCount *int64) {
	var post entity.SacrificeGruplar
	var count int64
	r.db.Debug().Model(post).Where(" durum <> 1 ").Count(&count)
	*postTotalCount = count
}

// SetGrupLideri upate data
func (r *GruplarRepo) SetGrupID(grupID uint64, kurbanID uint64, grupLideri int) {
	r.db.Debug().Table(entity.KurbanTableName).Where(" id = ?", kurbanID).Update(entity.SacrificeKurbanlar{GrupID: grupID, GrupLideri: grupLideri})
}

//// TODO: durum sogrulari entity e gore olacak
//func (r *GruplarRepo) ToplamOdemeler() float64 {
//	var result float64
//	row := r.db.Debug().Table(entity.KurbanTableName).
//		Select("sum(bakiye) as total").Where(" borc_durum <> " + kurbanBorcDurumKasaBorcluDurumda + "  AND durum <>  " + kurbanDurumKurbanKesildiTeslimEdildi).Row()
//	row.Scan(&result)
//	return result
//}

// ToplamOdemelerForYear
func (r *GruplarRepo) ToplamOdemelerForYear(year string) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).
		Select("sum(bakiye) as total").Where("kurban_bayrami_yili = ? AND  borc_durum <> "+kurbanBorcDurumKasaBorcluDurumda, year).Row()
	row.Scan(&result)
	return result
}

// KalanBorclarForYear
func (r *GruplarRepo) KalanBorclarForYear(year string) float64 {
	var result float64
	row := r.db.Debug().Table(entity.KurbanTableName).Select("sum(alacak) as total").Where("kurban_bayrami_yili = ? AND   borc_durum <>  "+kurbanBorcDurumKasaBorcluDurumda, year).Row()
	row.Scan(&result)
	return result
}

// GetByAllRelationsByID listeleme
func (r *GruplarRepo) GetByAllRelationsByID(id uint64, year string) (*dto.GruplarExcelandIndex, error) {
	var gruplarList dto.GruplarExcelandIndex

	// err = r.db.Debug().Where("  durum <> 1  ").Preload("Kurban").Find(&gruplarList).Error //entity.gruplara gore
	err := r.db.Debug().Table("sacrifice_gruplar").Where("  kurban_bayrami_yili = ? AND id = ? AND  "+"durum <> "+gruplarDurumKurbanBayramiDisiKesim, year, id).Order("kesim_sira_no asc").Take(&gruplarList).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &gruplarList, nil
}
