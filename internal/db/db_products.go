package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName

	uuid "github.com/satori/go.uuid"
)

type ProductRow struct {
	ID       string
	Name     sql.NullString
	Price    sql.NullString
	ImageUrl sql.NullString
}

func productRowToProduct(p ProductRow) ecommerce.Product {
	priceFloat, _ := strconv.ParseFloat(p.Price.String, 64)
	return ecommerce.Product{
		ID:       p.ID,
		Name:     p.Name.String,
		Price:    priceFloat,
		ImageUrl: p.ImageUrl.String,
	}
}

func (d *Database) UpsertProduct(ctx context.Context, p ecommerce.Product) (ecommerce.Product, error) {
	sqlQuery := `UPDATE products SET name=:name, price=:price, imageurl=:imageurl WHERE id=:id`
	if p.ID == "" {
		p.ID = uuid.NewV4().String()
		sqlQuery = `INSERT INTO products (id, name, price, imageurl) VALUES (:id, :name, :price, :imageurl)`
	}

	row := ProductRow{
		ID:       p.ID,
		Name:     sql.NullString{String: p.Name, Valid: true},
		Price:    sql.NullString{String: p.ToDecimal(), Valid: true},
		ImageUrl: sql.NullString{String: p.ImageUrl, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		sqlQuery,
		row,
	)
	if err != nil {
		return ecommerce.Product{}, fmt.Errorf("failed to upsert product: %w", err)
	}
	if err := rows.Close(); err != nil {
		return ecommerce.Product{}, fmt.Errorf("failed to close rows upsert: %w", err)
	}

	return productRowToProduct(row), nil
}

func (d *Database) GetProduct(ctx context.Context, uuid string) (ecommerce.Product, error) {
	var row ProductRow
	sqlQuery := `SELECT id, name, price, imageurl FROM products WHERE id=$1`

	rows := d.Client.QueryRowContext(ctx, sqlQuery, uuid)
	err := rows.Scan(&row.ID, &row.Name, &row.Price, &row.ImageUrl)
	if err != nil {
		return ecommerce.Product{}, fmt.Errorf("error getting the product by uuid: %w", err)
	}

	return productRowToProduct(row), nil
}

func (d *Database) GetProducts(ctx context.Context, order, search string) ([]ecommerce.Product, error) {
	var products []ecommerce.Product

	ord := ""
	if order != "" {
		ord = " ORDER BY price " + order
	}
	srch := ""
	if search != "" {
		srch = " WHERE LOWER(name) LIKE LOWER('" + search + "%')"
	}

	sqlQuery := "SELECT id, name, price, imageurl FROM products" + srch + ord

	rows, err := d.Client.QueryContext(ctx, sqlQuery)
	if err != nil {
		return []ecommerce.Product{}, fmt.Errorf("error getting the products: %w", err)
	}

	for rows.Next() {
		product := ProductRow{}
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageUrl); err != nil {
			return []ecommerce.Product{}, fmt.Errorf("error getting the products: %w", err)
		}
		products = append(products, productRowToProduct(product))
	}
	rows.Close()

	return products, nil
}

func (d *Database) DeleteProduct(ctx context.Context, id string) error {
	sqlQuery := `DELETE FROM products WHERE id=$1`
	_, err := d.Client.ExecContext(
		ctx,
		sqlQuery,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
