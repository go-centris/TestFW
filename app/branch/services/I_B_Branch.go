package regionServices_mod

import (
	"stncCms/app/branch/entity"
)

// BranchAppInterface service
type BranchAppInterface interface {
	Save(*entity.Branches) (*entity.Branches, map[string]string)
	GetByID(uint64) (*entity.Branches, error)
	GetByRegionID(regionID uint64) ([]entity.Branches, error)
	GetAll() ([]entity.Branches, error)
	Update(*entity.Branches) (*entity.Branches, map[string]string)
	Delete(uint64) error
	GetAllPaginate(postsPerPage int, offset int) ([]entity.Branches, error)
	GetAllPaginateCount(returnValue *int64)
}

// BranchApp struct  init
type BranchApp struct {
	request BranchAppInterface
}

var _ BranchAppInterface = &BranchApp{}

// Save service init
func (f *BranchApp) Save(Cat *entity.Branches) (*entity.Branches, map[string]string) {
	return f.request.Save(Cat)
}

// GetAll service init
func (f *BranchApp) GetAll() ([]entity.Branches, error) {
	return f.request.GetAll()
}

// GetByID service init
func (f *BranchApp) GetByID(CatID uint64) (*entity.Branches, error) {
	return f.request.GetByID(CatID)
}

// GetByRegionID service init
func (f *BranchApp) GetByRegionID(regionID uint64) ([]entity.Branches, error) {
	return f.request.GetByRegionID(regionID)
}

// Update service init
func (f *BranchApp) Update(Cat *entity.Branches) (*entity.Branches, map[string]string) {
	return f.request.Update(Cat)
}

// Delete service init
func (f *BranchApp) Delete(CatID uint64) error {
	return f.request.Delete(CatID)
}

// GetAllPaginate
func (f *BranchApp) GetAllPaginate(postsPerPage int, offset int) ([]entity.Branches, error) {
	return f.request.GetAllPaginate(postsPerPage, offset)
}

// GetAllPaginateCount
func (f *BranchApp) GetAllPaginateCount(returnValue *int64) {
	f.request.GetAllPaginateCount(returnValue)
}
