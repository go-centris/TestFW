package regionServices_mod

import (
	"stncCms/app/region/entity"
)

// RegionAppInterface service
type RegionAppInterface interface {
	Save(*entity.Region) (*entity.Region, map[string]string)
	GetByID(uint64) (*entity.Region, error)
	GetAll() ([]entity.Region, error)
	Update(*entity.Region) (*entity.Region, map[string]string)
	Delete(uint64) error
	GetAllPaginate(postsPerPage int, offset int) ([]entity.Region, error)
	GetAllPaginateCount(returnValue *int64)
}

// RegionApp struct  init
type RegionApp struct {
	request RegionAppInterface
}

var _ RegionAppInterface = &RegionApp{}

// Save service init
func (f *RegionApp) Save(Cat *entity.Region) (*entity.Region, map[string]string) {
	return f.request.Save(Cat)
}

// GetByID service init
func (f *RegionApp) GetByID(catID uint64) (*entity.Region, error) {
	return f.request.GetByID(catID)
}

// GetAll service init
func (f *RegionApp) GetAll() ([]entity.Region, error) {
	return f.request.GetAll()
}

// Update service init
func (f *RegionApp) Update(cat *entity.Region) (*entity.Region, map[string]string) {
	return f.request.Update(cat)
}

// Delete service init
func (f *RegionApp) Delete(catID uint64) error {
	return f.request.Delete(catID)
}

// GetAllPaginate list
func (f *RegionApp) GetAllPaginate(postsPerPage int, offset int) ([]entity.Region, error) {
	return f.request.GetAllPaginate(postsPerPage, offset)
}

// GetAllPaginateCount
func (f *RegionApp) GetAllPaginateCount(returnValue *int64) {
	f.request.GetAllPaginateCount(returnValue)
}
