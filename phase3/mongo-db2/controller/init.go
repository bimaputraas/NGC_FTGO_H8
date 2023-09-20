package controller

import "ngc2_p3/repository"

type Controller struct {
	Repository *repository.Repository
}

func NewController(repository *repository.Repository) *Controller {
	return &Controller{
		Repository: repository,
	}
}