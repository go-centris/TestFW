package cacheRepository

import (
	"stncCms/app/post/entity"
	// repository "stncCms/app/domain/repository/dbRepository"
	postRepository "stncCms/app/post/repository/dbRepository"
)

// Save data
func (r *CatPostRepo) Save(data *entity.CategoryPosts) (*entity.CategoryPosts, map[string]string) {
	repo := postRepository.CatPostRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *CatPostRepo) Update(data *entity.CategoryPosts) (*entity.CategoryPosts, map[string]string) {
	repo := postRepository.CatPostRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

// Delete delete data
func (r *CatPostRepo) Delete(id uint64) error {
	repo := postRepository.CatPostRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}

// DeleteForPostID delete data
func (r *CatPostRepo) DeleteForPostID(postID uint64) error {
	repo := postRepository.CatPostRepositoryInit(r.db)
	err := repo.DeleteForPostID(postID)
	return err
}

// DeleteForCatID delete data
func (r *CatPostRepo) DeleteForCatID(CatID uint64) error {
	repo := postRepository.CatPostRepositoryInit(r.db)
	err := repo.DeleteForCatID(CatID)
	return err
}
