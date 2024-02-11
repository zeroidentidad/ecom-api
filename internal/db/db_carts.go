package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName
)

type CartRow struct {
	UserId    sql.NullString
	ProductId sql.NullString
}

func cartRowToCart(c CartRow) ecommerce.Cart {
	return ecommerce.Cart{
		UserId:    c.UserId.String,
		ProductId: c.ProductId.String,
	}
}

func (d *Database) PostItemCart(ctx context.Context, c ecommerce.Cart) (ecommerce.Cart, error) {
	row := CartRow{
		UserId:    sql.NullString{String: c.UserId, Valid: true},
		ProductId: sql.NullString{String: c.ProductId, Valid: true},
	}
	sqlQuery := `INSERT INTO carts (userid, productid) VALUES (:userid, :productid)`

	rows, err := d.Client.NamedQueryContext(
		ctx,
		sqlQuery,
		row,
	)
	if err != nil {
		return ecommerce.Cart{}, fmt.Errorf("failed to insert item cart: %w", err)
	}
	if err := rows.Close(); err != nil {
		return ecommerce.Cart{}, fmt.Errorf("failed to close rows insert: %w", err)
	}

	return c, nil
}

func (d *Database) GetItemsCart(ctx context.Context, userid string) ([]ecommerce.Cart, error) {
	var cartItems []ecommerce.Cart
	sqlQuery := `SELECT userid, productid FROM carts WHERE userid=$1`

	rows, err := d.Client.QueryContext(ctx, sqlQuery, userid)
	if err != nil {
		return []ecommerce.Cart{}, fmt.Errorf("error getting the cart by userid: %w", err)
	}

	for rows.Next() {
		cartItem := CartRow{}
		if err = rows.Scan(&cartItem.UserId, &cartItem.ProductId); err != nil {
			return []ecommerce.Cart{}, fmt.Errorf("error getting the cart items: %w", err)
		}
		cartItems = append(cartItems, cartRowToCart(cartItem))
	}
	rows.Close()

	return cartItems, nil
}

func (d *Database) DeleteItemCart(ctx context.Context, userid, productid string) error {
	sqlQuery := `DELETE FROM carts WHERE userid=$1 AND productid=$2`
	_, err := d.Client.ExecContext(
		ctx,
		sqlQuery,
		userid,
		productid,
	)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	return nil
}
