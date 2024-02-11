package ecommerce

import "errors"

var (
	ErrNotImplemented  = errors.New("not implemented")
	ErrCreatingProduct = errors.New("failed to create product")
	ErrFetchingProduct = errors.New("failed to fetch product")
	ErrCreatingCart    = errors.New("failed to create cart")
	ErrFetchingCart    = errors.New("failed to fetch cart")
)
