package ecommerce

import (
	"context"
	"fmt"
)

// representation of product
// as structure for service
type Product struct {
	ID       string  `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Price    float64 `json:"price" db:"price"`
	ImageUrl string  `json:"imageurl" db:"imageurl"`
}

func (p Product) ToDecimal() string {
	return fmt.Sprintf("%.2f", p.Price)
}

func (s *Service) UpsertProduct(ctx context.Context, p Product) (Product, error) {
	p, err := s.StoreService.UpsertProduct(ctx, p)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func (s *Service) GetProduct(ctx context.Context, id string) (Product, error) {
	p, err := s.StoreService.GetProduct(ctx, id)
	if err != nil {
		fmt.Println(err)                     // handle as internal logging
		return Product{}, ErrFetchingProduct // expose to client
	}
	return p, nil
}

func (s *Service) GetProducts(ctx context.Context, order, search string) ([]Product, error) {
	p, err := s.StoreService.GetProducts(ctx, order, search)
	if err != nil {
		fmt.Println(err)                       // handle as internal logging
		return []Product{}, ErrFetchingProduct // expose to client
	}
	return p, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	return s.StoreService.DeleteProduct(ctx, id)
}
