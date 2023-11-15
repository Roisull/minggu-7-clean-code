package usecase

import (
	"belajar-go-echo/app/model"
	"belajar-go-echo/app/repository"
)

// implementasi usecase untuk entitas User
type UserUsecase struct {
	userRepository repository.UserRepository
}

// membuat instance UserUsecase
func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase{
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// kembalikan semua user
func (u *UserUsecase) GetAllUsers() ([]model.User, error){
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// membuat user baru
func (u *UserUsecase) CreateUser(user *model.User) (*model.User, error){
	createdUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// usecase untuk interaksi get all user with auth
func (u *UserUsecase) GetAuthenticatedUsers() ([]model.User, error) {
	// logika bisnis untuk mendapatkan pengguna yang memerlukan otentikasi
	users, err := u.userRepository.GetAuthenticatedUsers()
	if err != nil{
		return nil, err
	}
	// Panggil metode dari userRepository atau repositori data
	return users, nil
 }
