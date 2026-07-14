package service

import "github.com/Barbagidon/shop-api/internal/domain"

type ProductService struct {
	products []domain.Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []domain.Product{},
	}
}

func (s *ProductService) GetProducts() []domain.Product {
	return s.products
}
