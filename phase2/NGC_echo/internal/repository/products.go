package repository

import (
	"context"
	"errors"
	"ngc_echo/internal/model"
)

func (r *Repository) GetAllProducts(ctx context.Context) ([]*model.Products, error) {
	var products []*model.Products
	result := r.db.WithContext(ctx).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Data does not exist")
	}

	return products, nil
}
