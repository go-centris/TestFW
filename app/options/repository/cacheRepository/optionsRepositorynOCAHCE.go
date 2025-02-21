package cacheRepository

import (
	"stncCms/app/options/entity"
	// repository "stncCms/app/domain/repository/dbRepository"
	optionRepository "stncCms/app/options/repository/dbRepository"

)

//AddOption data (save)
func (r *OptionRepositoryRepo) AddOption(data *entity.Options) (*entity.Options, map[string]string) {
	repo := optionRepository.OptionRepositoryInit(r.db)
	datas, err := repo.AddOption(data)
	return datas, err
}

//SetOption upate data
func (r *OptionRepositoryRepo) SetOption(name, value string) {
	repo := optionRepository.OptionRepositoryInit(r.db)
	repo.SetOption(name, value)

}

//GetOptionID get data
func (r *OptionRepositoryRepo) GetOptionID(name string) (returnValue int) {
	repo := optionRepository.OptionRepositoryInit(r.db)
	return repo.GetOptionID(name)
}

//GetOption get data
func (r *OptionRepositoryRepo) GetOption(name string) string {
	repo := optionRepository.OptionRepositoryInit(r.db)
	return repo.GetOption(name)
}

//DeleteOptionID data
func (r *OptionRepositoryRepo) DeleteOptionID(id uint64) error {

	repo := optionRepository.OptionRepositoryInit(r.db)
	err := repo.DeleteOptionID(id)
	return err
}

//DeleteOption data
func (r *OptionRepositoryRepo) DeleteOption(value string) error {
	repo := optionRepository.OptionRepositoryInit(r.db)
	err := repo.DeleteOption(value)
	return err
}
