package service

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/repository"
)

type TokoService struct {
	repository *repository.TokoRepository
}

func NewTokoService(repository *repository.TokoRepository) *TokoService {
	return &TokoService{repository: repository}
}

func (t TokoService) Create(toko *model.Toko) bool {
	err := t.repository.Create(toko)

	if err != nil {
		return false
	}

	return true
}

func (t TokoService) FindByUser(user model.User) (*model.Toko, bool) {
	toko, err := t.repository.FindByUser(user)

	if err != nil {
		return nil, false
	}

	return toko, true
}
