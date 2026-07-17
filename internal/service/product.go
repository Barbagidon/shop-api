package service

import "github.com/Barbagidon/shop-api/internal/domain"

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.repo.GetAll()
}
