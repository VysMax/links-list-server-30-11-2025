package controller

import (
	"github.com/VysMax/links-list-server/usecase"
)

type LinksHandler struct {
	service *usecase.LinksService
}

func New(s *usecase.LinksService) *LinksHandler {
	return &LinksHandler{
		service: s,
	}
}
