package dbRepository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

// ReportRepo struct
type ReportRepo struct {
	db *gorm.DB
}

// ReportRepositoryInit initial
func ReportRepositoryInit(db *gorm.DB) *ReportRepo {
	return &ReportRepo{db}
}

// GetAllUsersWhoAddedMostSacrifeAndBranchCount En cok kurban ekleyen  subemiz count
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndBranchCount(totalCount *int64) {
	var data int64
	query := r.db.Debug().Table(entity.KurbanTableName).Select(`br.id AS branch_id,COUNT(br) AS counts`)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches  ON br_branches.id=users.branch_id")
	query = query.Group("br.id")
	query.Row().Scan(&data)
	*totalCount = data
}

// GetAllUsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var usersWhoAddedMostSacrifeAndBranch []dto.UsersWhoAddedMostSacrifeAndBranch
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(
		`br.id AS branch_id,COUNT(br) AS counts,br.title AS branch `)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches AS br ON br.id=users.branch_id")
	query = query.Group("br.id")
	query = query.Order("counts DESC")
	query = query.Limit(postsPerPage)
	query = query.Offset(offset)
	err = query.Find(&usersWhoAddedMostSacrifeAndBranch).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return usersWhoAddedMostSacrifeAndBranch, nil
}

// GetAllUsersWhoAddedMostSacrifeAndUserCount  En cok kurban ekleyen hocamiz Count
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndUserCount(totalCount *int64) {
	var data int64
	query := r.db.Debug().Table(entity.KurbanTableName).Select(`COUNT(sacrifice_kurbanlar.user_id) AS d`)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches  ON br_branches.id=users.branch_id")
	query = query.Group("users.id")
	query.Row().Scan(&data)
	*totalCount = data
}

// GetAllUsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
func (r *ReportRepo) GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	var usersWhoAddedMostSacrifeAndBranch []dto.UsersWhoAddedMostSacrifeAndBranch
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(`sacrifice_kurbanlar.user_id AS sacrife_user_id, COUNT(users) AS counts,users.first_name,
     users.last_name,users.email,br.title AS branch`)
	query = query.Joins(" JOIN users ON users.id=sacrifice_kurbanlar.user_id")
	query = query.Joins("JOIN br_branches AS br ON br.id=users.branch_id")
	query = query.Group("users.id,sacrife_user_id,br.id")
	query = query.Order(" counts DESC")
	query = query.Limit(postsPerPage)
	query = query.Offset(offset)
	err = query.Find(&usersWhoAddedMostSacrifeAndBranch).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return usersWhoAddedMostSacrifeAndBranch, nil
}

// *** En cok kurban kestiren hayirsever ***/////

// GetAllCharitableWhoAddedMostSacrife  En cok kurban kestiren hayirsever
func (r *ReportRepo) GetAllCharitableWhoAddedMostSacrife(postsPerPage int, offset int) ([]dto.CharitableWhoAddedMostSacrife, error) {
	var charitableWhoAddedMostSacrife []dto.CharitableWhoAddedMostSacrife
	var err error
	query := r.db.Debug().Table(entity.KurbanTableName).Select(`sacrifice_kurbanlar.user_id,ki.id as kisi_id, COUNT(*) AS counts,ki.ad_soyad AS name_lastname`)
	query = query.Joins("INNER JOIN sacrifice_kisiler AS ki ON ki.id=sacrifice_kurbanlar.kisi_id")
	query = query.Group("sacrifice_kurbanlar.user_id,ki.ad_soyad,ki.id ")
	query = query.Having("ki.ad_soyad <> 'Boş GRUP, Kişi kaydı için tıklayınız'")
	query = query.Order(" counts DESC")
	query = query.Limit(postsPerPage)
	query = query.Offset(offset)
	err = query.Find(&charitableWhoAddedMostSacrife).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return charitableWhoAddedMostSacrife, nil
}

// GetAllCharitableWhoAddedMostSacrifeCount   En cok kurban kestiren hayirsever Count
func (r *ReportRepo) GetAllCharitableWhoAddedMostSacrifeCount(totalCount *int64) {
	var data int64
	query := r.db.Debug().Table(entity.KurbanTableName).Select(` COUNT(*) AS counts`)
	query = query.Joins("INNER JOIN sacrifice_kisiler AS ki ON ki.id=sacrifice_kurbanlar.kisi_id")
	query = query.Group("sacrifice_kurbanlar.user_id,ki.ad_soyad,ki.id ")
	query = query.Having("ki.ad_soyad <> 'Boş GRUP, Kişi kaydı için tıklayınız'")
	query = query.Order(" counts DESC")
	query.Row().Scan(&data)
	*totalCount = data
}
