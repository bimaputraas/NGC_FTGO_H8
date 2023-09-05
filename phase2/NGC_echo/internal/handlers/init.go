package handlers

import (
	"ngc_echo/internal/usecase"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}
