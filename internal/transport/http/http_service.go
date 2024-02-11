package http

import (
	"context"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName
)

type EcommerceService interface {
	UpsertProduct(context.Context, ecommerce.Product) (ecommerce.Product, error)
	GetProduct(context.Context, string) (ecommerce.Product, error)
	GetProducts(context.Context, string, string) ([]ecommerce.Product, error)
	DeleteProduct(context.Context, string) error

	PostItemCart(context.Context, ecommerce.Cart) (ecommerce.Cart, error)
	GetItemsCart(context.Context, string) ([]ecommerce.Cart, error)
	DeleteItemCart(context.Context, string, string) error
}

// representation of cart
// to generate doc
type Cart struct {
	UserId    string `json:"userid"`
	ProductId string `json:"productid"`
}

// representation of product
// to generate doc
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageurl"`
}

// representation of default response
// to generate doc
type message struct {
	Message string `json:"message"`
}
