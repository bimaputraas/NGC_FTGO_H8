package usecase

import (
	"context"
	"errors"
	"ngc_echo/internal/model"
	"ngc_echo/pkg"
)

func (u *Usecase) Register(ctx context.Context, usr *model.Users) error {
	hashedPass, err := pkg.HashPassword(usr.Password)
	if err != nil {
		return errors.New("Failed hash")
	}
	usr.Password = hashedPass

	return u.repository.InsertUser(ctx, usr)
}

func (u *Usecase) Login(ctx context.Context, username string, pass string) (*model.Users, string, error) {
	hashedPass, err := pkg.HashPassword(pass)
	if err != nil {
		return nil, "", errors.New("Failed hash")
	}

	user, err := u.repository.GetUser(ctx, username, hashedPass)
	if err != nil {
		return nil, "", err
	}

	// compare hash
	if !pkg.CheckPasswordHash(pass, hashedPass) {
		return nil, "", errors.New("Failed compare hash")
	}

	// generate token
	tokenString, err := pkg.GenerateToken(int(user.ID))
	if err != nil {
		return nil, "", errors.New("Failed generate token")
	}

	return user, tokenString, nil
}

func (u *Usecase) GetAllProducts(ctx context.Context) ([]*model.Products, error) {
	return u.repository.GetAllProducts(ctx)
}

func (u *Usecase) InsertTransaction(ctx context.Context, usr *model.Users, trx *model.Transactions) (*model.Transactions, error) {
	return u.repository.InsertTransaction(ctx, usr, trx)
}
