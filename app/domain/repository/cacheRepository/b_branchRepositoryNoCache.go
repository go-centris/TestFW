package cacheRepository

import (
	// "stncCms/app/domain/entity"
	// repository "stncCms/app/domain/repository/dbRepository"
		// repository "stncCms/app/domain/repository/dbRepository"
	branchRepository "stncCms/app/branch/repository/dbRepository"
	branchEntity "stncCms/app/branch/entity"
)

//Save data
func (r *BranchRepo) Save(data *branchEntity.Branches) (*branchEntity.Branches, map[string]string) {
	repo := branchRepository.BranchRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *BranchRepo) Update(data *branchEntity.Branches) (*branchEntity.Branches, map[string]string) {
	repo := branchRepository.BranchRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err

}

//Delete delete data
func (r *BranchRepo) Delete(id uint64) error {
	repo := branchRepository.BranchRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
