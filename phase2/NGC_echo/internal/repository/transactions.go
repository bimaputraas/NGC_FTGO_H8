package repository

import (
	"context"
	"errors"
	"ngc_echo/internal/model"
)

func (r *Repository) InsertTransaction(ctx context.Context, user *model.Users, trx *model.Transactions) (*model.Transactions, error) {
	// init tx for transaction by gorm
	tx := r.db.WithContext(ctx).Begin()

	// find product
	var product model.Products
	result := tx.Where("id = ?", trx.ProductID).First(&product)
	if result.Error != nil {
		tx.Rollback()
		return nil, errors.New("Data does not exist")
	}

	// check if stock unavailable
	if product.Stock < trx.Quantity {
		tx.Rollback()
		return nil, errors.New("Stock unavailable")
	}

	// if stock available
	product.Stock -= trx.Quantity
	result = tx.Save(&product)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// check if user fund less than total amount transaction
	trx.TotalAmount = product.Price * float64(trx.Quantity)
	if user.DepositAmount < trx.TotalAmount {
		tx.Rollback()
		return nil, errors.New("The fund is insufficient")
	}

	// if fund is sufficient
	user.DepositAmount -= trx.TotalAmount
	result = tx.Save(user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// create transaction information
	trx.UserID = user.ID
	result = tx.Create(&trx)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// success purchase
	tx.Commit()
	return trx, nil
}
