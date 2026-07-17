package memory

import "github.com/Barbagidon/shop-api/internal/domain"

type ProductRepository struct {
	products []domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []domain.Product{

			{
				ID:    1,
				Name:  "iPhone",
				Price: 1000,
			},
			{
				ID:    2,
				Name:  "MacBook",
				Price: 2500,
			},
		},
	}
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	return r.products, nil
}
