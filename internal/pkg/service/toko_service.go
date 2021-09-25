package service

import "ecommerce/internal/pkg/repository"

type TokoService struct {
	repository *repository.TokoRepository
}

func NewTokoService(repository *repository.TokoRepository) *TokoService {
	return &TokoService{repository: repository}
}
