package service

import (
	"github.com/webtoon/internal/app/model"
	"github.com/webtoon/internal/app/repository"
	"github.com/webtoon/internal/app/request"
)

type ManhwaService interface {
	GetManhwaByTitle(title string) (*model.Manhwa, error)
	GetManhwas() ([]*model.Manhwa, error)
}

type manhwaService struct {
	repository repository.ManhwaRepository
}

func NewManhwaService(repository repository.ManhwaRepository) ManhwaService {
	return &manhwaService{
		repository: repository,
	}
}

func (s *manhwaService) GetManhwas() ([]*model.Manhwa, error) {
	return s.repository.GetManhwas()
}

func (s *manhwaService) GetManhwaByTitle(title string) (*model.Manhwa, error) {
	return s.repository.GetManhwaByTitle(title)
}

func (s *manhwaService) CreateManhwa(request *request.CreateManhwaRequest) (*model.Manhwa, error) {
	return s.repository.CreateManhwa()

}
