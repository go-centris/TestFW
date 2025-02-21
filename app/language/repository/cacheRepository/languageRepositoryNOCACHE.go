package cacheRepository

import (
	"stncCms/app/language/entity"
	
)

//Save data
func (r *LanguageRepo) Save(data *entity.Languages) (*entity.Languages, map[string]string) {
	repo := LanguageRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *LanguageRepo) Update(data *entity.Languages) (*entity.Languages, map[string]string) {
	repo := LanguageRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

//Delete data
func (r *LanguageRepo) Delete(id uint64) error {
	repo := LanguageRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}