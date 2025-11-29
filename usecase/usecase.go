package usecase

import "github.com/VysMax/links-list-server/entities"

type LinksUsecase interface {
	SaveLinks(entities.LinksToSave) error
	GetLinks(entities.LinksNumRequest) error
}

type LinksService struct {
	repo LinksUsecase
}

func New(repo LinksUsecase) *LinksService {
	return &LinksService{
		repo: repo,
	}
}

func (s *LinksService) SaveLinks(linksList entities.LinksToSave) error {
	err := s.repo.SaveLinks(linksList)
	if err != nil {
		return err
	}
	return nil
}

func (s *LinksService) GetLinks(linkNums entities.LinksNumRequest) error {
	err := s.repo.GetLinks(linkNums)
	if err != nil {
		return err
	}
	return nil
}
