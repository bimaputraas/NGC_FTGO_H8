package handler

import (
	"ms-paylater/repository"

	"gorm.io/gorm"
)

type Handler struct {
	Repository repository.Repository
}

func InitHandler(db *gorm.DB) *Handler {
	return &Handler{
		Repository: repository.Repository{
			DB: db,
		},
	}
}
