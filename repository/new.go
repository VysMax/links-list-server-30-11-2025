package repository

import "github.com/VysMax/links-list-server/usecase"

type Repository struct{}

func New() usecase.LinksUsecase {
	return &Repository{}
}
