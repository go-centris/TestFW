package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

func (r *KisilerRepo) Save(data *entity.Kisiler) (*entity.Kisiler, map[string]string) {
	repo := repository.KisilerRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *KisilerRepo) Update(data *entity.Kisiler) (*entity.Kisiler, map[string]string) {
	repo := repository.KisilerRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

// Delete delete data
func (r *KisilerRepo) Delete(id uint64) error {
	repo := repository.KisilerRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}

// Count fat
func (r *KisilerRepo) Count(totalCount *int64) {
	var count int64
	repo := repository.KisilerRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count

}

// comment data
func (r *KisilerRepo) SetAciklama(kisiID uint64, value string) {
	repo := repository.KisilerRepositoryInit(r.db)
	repo.SetAciklama(kisiID, value)
}

// HasKisiKurban kişi kurban listesinde var mı
func (r *KisilerRepo) HasKisiKurban(id uint64, postTotalCount *int) {
	var kurbanData entity.SacrificeKurbanlar
	var count int
	r.db.Debug().Model(kurbanData).Where("kisi_id = ?", id).Count(&count)
	*postTotalCount = count
}

// HasKisiReferans kişi referans listesinde var mı?
func (r *KisilerRepo) HasKisiReferans(id uint64, totalCount *int) {
	var count int
	repo := repository.KisilerRepositoryInit(r.db)
	repo.HasKisiReferans(id, &count)
	*totalCount = count
}

func (r *KisilerRepo) GetPersonComment(id uint64, aciklamaOut *string) {
	var aciklama string
	repo := repository.KisilerRepositoryInit(r.db)
	repo.GetPersonComment(id, &aciklama)
	*aciklamaOut = aciklama
}
