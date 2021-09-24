package service

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRespository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRespository: userRepository}
}

func (u UserService) Login(username string, password string) (*model.User, bool) {
	user, err := u.userRespository.FindByCredentials(username)

	if err != nil {
		return nil, false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, false
	}

	return &user, true
}
