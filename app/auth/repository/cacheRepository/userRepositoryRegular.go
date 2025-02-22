package cacheRepository

import (
	authDto "stncCms/app/auth/dto"
	authEntity "stncCms/app/auth/entity"
	authRepository "stncCms/app/auth/repository/dbRepository"
)

func (r *UserRepo) SaveUser(data *authEntity.Users) (*authEntity.Users, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	datas, err := repo.SaveUser(data)
	return datas, err
}

// Save data
func (r *UserRepo) Save(data *authEntity.Users) (*authEntity.Users, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *UserRepo) SaveDto(data *authDto.User) (*authDto.User, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	datas, err := repo.SaveDto(data)
	return datas, err
}

func (r *UserRepo) UpdateDto(data *authDto.User) (*authDto.User, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	datas, err := repo.UpdateDto(data)
	return datas, err
}

// Update upate data
func (r *UserRepo) Update(data *authEntity.Users) (*authEntity.Users, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

// Count fat
func (r *UserRepo) Count(totalCount *int64) {
	var count int64
	repo := authRepository.UserRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count
}

// Delete data
func (r *UserRepo) Delete(id uint64) error {
	repo := authRepository.UserRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}


func (r *UserRepo) SetUserStatusUpdate(id uint64, status int) {
	repo := authRepository.UserRepositoryInit(r.db)
	repo.SetUserStatus(id, status)
}

// for api 
func (r *UserRepo) GetUserByEmailAndPassword(u *authEntity.Users) (*authEntity.Users, map[string]string) {
	repo := authRepository.UserRepositoryInit(r.db)
	data, _ := repo.GetUserByEmailAndPassword(u)
	return data, nil
}
func (r *UserRepo) GetUserByEmailAndPassword2(email string, InputPassword string) (*authEntity.Users, bool) {
	repo := authRepository.UserRepositoryInit(r.db)
	data, result := repo.GetUserByEmailAndPassword2(email, InputPassword)
	return data, result
}
func (r *UserRepo) SetUserPassword(id uint64, password string) {
	repo := authRepository.UserRepositoryInit(r.db)
	repo.SetUserPassword(id, password)
}

func (r *UserRepo) GetByUserForBranchID(branchID int) (*authEntity.UsersGetByUserForBranchIDDTO, error) {
	repo := authRepository.UserRepositoryInit(r.db)
	data, _ := repo.GetByUserForBranchID(branchID)
	return data, nil
}
