package repository

import (
	"context"
	"errors"
	"ngc_echo/internal/model"
)

func (r *Repository) InsertUser(ctx context.Context, usr *model.Users) error {
	result := r.db.WithContext(ctx).Create(usr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetUser(ctx context.Context, username string, pass string) (*model.Users, error) {
	// find from db by username
	var usr model.Users
	result := r.db.Where("username = ?", username).Find(&usr)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Data does not exist")
	}

	return &usr, nil
}
