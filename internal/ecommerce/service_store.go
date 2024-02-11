package ecommerce

import (
	"context"
)

// all methods that service needs to operate
type Store interface {
	UpsertProduct(context.Context, Product) (Product, error)
	GetProduct(context.Context, string) (Product, error)
	GetProducts(context.Context, string, string) ([]Product, error)
	DeleteProduct(context.Context, string) error

	PostItemCart(context.Context, Cart) (Cart, error)
	GetItemsCart(context.Context, string) ([]Cart, error)
	DeleteItemCart(context.Context, string, string) error
}

// struct where logic will be built on top of
type Service struct {
	StoreService Store
}

// return a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		StoreService: store,
	}
}
