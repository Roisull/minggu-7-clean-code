package repository

import (
	"belajar-go-echo/app/model"

	"gorm.io/gorm"
)

// Implementasi UserRepository menggunakan gorm
type UserRepositoryImpl struct {
	db *gorm.DB
}

// GetAllUsers implements UserRepository.
func (r *UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	// panic("unimplemented")

	return users, nil
}

// // GetAllUsers implements UserRepository.
// func (*UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
// 	panic("unimplemented")
// }

// membuat instance UserRepositoryImpl
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// CreateUser membuat user baru
func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAuthenticatedUsers implements UserRepository.
func (r *UserRepositoryImpl) GetAuthenticatedUsers() ([]model.User, error) {
	// logika untuk mendapatkan pengguna yang memerlukan otentikasi
	// menggunakan kondisi tertentu seperti status otentikasi di dalam query
	var users []model.User
	err := r.db.Where("token IS NOT NULL").Find(&users).Error
	if err != nil {
	   return nil, err
	}
	return users, nil
 }
