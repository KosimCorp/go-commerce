package service

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/repository"
)

type ProdukService struct {
	produkRespository *repository.ProdukRepository
}

func NewProdukService(produkRespository *repository.ProdukRepository) *ProdukService {
	return &ProdukService{produkRespository: produkRespository}
}

func (p ProdukService) FindAll() (*[]model.Produk, bool) {
	data, err := p.produkRespository.FindAll()

	if err != nil {
		return nil, false
	}

	return &data, true
}

func (p ProdukService) Find(id int) (*model.Produk, bool) {
	data, err := p.produkRespository.Find(id)

	if err != nil {
		return nil, false
	}

	return &data, true
}

func (p ProdukService) Create(produk *model.Produk) bool {
	err := p.produkRespository.Create(produk)

	if err != nil {
		return false
	}

	return true
}
