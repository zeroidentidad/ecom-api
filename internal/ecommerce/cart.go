package ecommerce

import (
	"context"
	"fmt"
)

// representation of cart
// as structure for service
type Cart struct {
	UserId    string `json:"userid" db:"userid"`
	ProductId string `json:"productid" db:"productid"`
}

func (s *Service) PostItemCart(ctx context.Context, c Cart) (Cart, error) {
	c, err := s.StoreService.PostItemCart(ctx, c)
	if err != nil {
		return Cart{}, err
	}

	return c, nil
}

func (s *Service) GetItemsCart(ctx context.Context, userid string) ([]Cart, error) {
	c, err := s.StoreService.GetItemsCart(ctx, userid)
	if err != nil {
		fmt.Println(err)                 // handle as internal logging
		return []Cart{}, ErrFetchingCart // expose to client
	}
	return c, nil
}

func (s *Service) DeleteItemCart(ctx context.Context, userid, productid string) error {
	return s.StoreService.DeleteItemCart(ctx, userid, productid)
}
