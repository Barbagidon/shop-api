package memory

import (
	"errors"

	"github.com/Barbagidon/shop-api/internal/domain"
)

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

func (r *ProductRepository) GetByID(id int64) (domain.Product, error) {
	for _, value := range r.products {
		if value.ID == id {
			return value, nil
		}
	}

	return domain.Product{}, errors.New("product not found")
}
