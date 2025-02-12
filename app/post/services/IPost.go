package cmsServices_mod

import (
	"stncCms/app/post/entity"
)

// PostAppInterface interface
type PostAppInterface interface {
	Save(*entity.Post) (*entity.Post, map[string]string)
	GetByID(uint64) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
	GetAllPagination(int, int) ([]entity.Post, error)
	Update(*entity.Post) (*entity.Post, map[string]string)
	Count(*int64)
	Delete(uint64) error
}
type postApp struct {
	request PostAppInterface
}

var _ PostAppInterface = &postApp{}

func (f *postApp) Save(post *entity.Post) (*entity.Post, map[string]string) {
	return f.request.Save(post)
}

func (f *postApp) GetByID(postID uint64) (*entity.Post, error) {
	return f.request.GetByID(postID)
}

func (f *postApp) GetAll() ([]entity.Post, error) {
	return f.request.GetAll()
}

func (f *postApp) GetAllPagination(postsPerPage int, offset int) ([]entity.Post, error) {
	return f.request.GetAllPagination(postsPerPage, offset)
}

func (f *postApp) Update(post *entity.Post) (*entity.Post, map[string]string) {
	return f.request.Update(post)
}

func (f *postApp) Count(postTotalCount *int64) {
	f.request.Count(postTotalCount)
}

func (f *postApp) Delete(postID uint64) error {
	return f.request.Delete(postID)
}
