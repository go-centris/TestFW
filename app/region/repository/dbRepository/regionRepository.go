package dbRepository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"stncCms/app/region/entity"
	"strings"
)

//repost repo ve dahsboard repoda ornekler var

// RegionRepo struct
type RegionRepo struct {
	db *gorm.DB
}

// RegionRepositoryInit initial
func RegionRepositoryInit(db *gorm.DB) *RegionRepo {
	return &RegionRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatAppInterface = &RegionRepo{}

// Save data
func (r *RegionRepo) Save(cat *entity.Region) (*entity.Region, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&cat).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "post title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return cat, nil
}

// Update upate data
func (r *RegionRepo) Update(cat *entity.Region) (*entity.Region, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&cat).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return cat, nil
}

// GetByID get data
func (r *RegionRepo) GetByID(id uint64) (*entity.Region, error) {
	var cat entity.Region
	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")

	}
	return &cat, nil
}

// GetAll all data
func (r *RegionRepo) GetAll() ([]entity.Region, error) {
	var cat []entity.Region
	err := r.db.Debug().Order("id asc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return cat, nil
}

// GetAllPaginate paginatin list
func (r *RegionRepo) GetAllPaginate(postsPerPage int, offset int) ([]entity.Region, error) {
	var regionList []entity.Region
	err := r.db.Debug().Order("id asc").Limit(postsPerPage).Offset(offset).Find(&regionList).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return regionList, nil
}

// GetAllPaginateCount
func (r *RegionRepo) GetAllPaginateCount(returnValue *int64) {
	var table entity.Region
	var count int64
	r.db.Debug().Model(table).Count(&count)
	*returnValue = count
}

// Delete delete data
func (r *RegionRepo) Delete(id uint64) error {
	var cat entity.Region
	err := r.db.Debug().Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
