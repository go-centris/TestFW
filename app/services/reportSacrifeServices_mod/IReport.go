package reportSacrifeServices_mod

import (
	"stncCms/app/domain/dto"
)

// ReportAppInterface interface
type ReportAppInterface interface {
	GetAllUsersWhoAddedMostSacrifeAndBranch(int, int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) // GetAllUsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
	GetAllUsersWhoAddedMostSacrifeAndBranchCount(*int64)                                               //GetAllUsersWhoAddedMostSacrifeAndBranchCount En cok kurban ekleyen  subemiz count

	GetAllUsersWhoAddedMostSacrifeAndUser(int, int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) // GetAllUsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
	GetAllUsersWhoAddedMostSacrifeAndUserCount(*int64)                                               //GetAllUsersWhoAddedMostSacrifeAndUserCount  En cok kurban ekleyen hocamiz  count

	GetAllCharitableWhoAddedMostSacrife(int, int) ([]dto.CharitableWhoAddedMostSacrife, error) // GetAllUsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
	GetAllCharitableWhoAddedMostSacrifeCount(*int64)                                           //GetAllUsersWhoAddedMostSacrifeAndUserCount  En cok kurban ekleyen hocamiz  count

}
type reportApp struct {
	request ReportAppInterface
}

var _ ReportAppInterface = &reportApp{}

// GetAllUsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen hocamiz ve subemiz
func (f *reportApp) GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	return f.request.GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage, offset)
}

// GetAllUsersWhoAddedMostSacrifeAndBranchCount En cok kurban ekleyen  subemiz count
func (f *reportApp) GetAllUsersWhoAddedMostSacrifeAndBranchCount(count *int64) {
	f.request.GetAllUsersWhoAddedMostSacrifeAndBranchCount(count)
}

// GetAllUsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz ve subemiz
func (f *reportApp) GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage int, offset int) ([]dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	return f.request.GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage, offset)
}

// GetAllUsersWhoAddedMostSacrifeAndUserCount  En cok kurban ekleyen hocamiz  count
func (f *reportApp) GetAllUsersWhoAddedMostSacrifeAndUserCount(count *int64) {
	f.request.GetAllUsersWhoAddedMostSacrifeAndUserCount(count)
}

// GetAllUsersWhoAddedMostSacrifeAndUser  En cok kurban kestiren hayirsever
func (f *reportApp) GetAllCharitableWhoAddedMostSacrife(postsPerPage int, offset int) ([]dto.CharitableWhoAddedMostSacrife, error) {
	return f.request.GetAllCharitableWhoAddedMostSacrife(postsPerPage, offset)
}

// GetAllUsersWhoAddedMostSacrifeAndUserCount    En cok kurban kestiren hayirsever Count
func (f *reportApp) GetAllCharitableWhoAddedMostSacrifeCount(count *int64) {
	f.request.GetAllCharitableWhoAddedMostSacrifeCount(count)
}
