package middleware

import (
	"ms-paylater/repository"

	"gorm.io/gorm"
)

type Auth struct {
	Repository repository.Repository
}

func InitAuth(db *gorm.DB) *Auth{
	return &Auth{
		Repository: repository.Repository{
			DB: db,
		},
	}
}