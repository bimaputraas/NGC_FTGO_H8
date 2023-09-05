package usecase

import "ngc_echo/internal/repository"

type Usecase struct {
	repository *repository.Repository
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		repository: repo,
	}
}
