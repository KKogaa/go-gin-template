package repository

import (
	"github.com/webtoon/internal/app/model"
	"github.com/webtoon/internal/app/request"
	"gorm.io/gorm"
)

type ManhwaRepository interface {
	GetManhwas() ([]*model.Manhwa, error)
	GetManhwaByTitle(title string) (*model.Manhwa, error)
}

type manhwaRepository struct {
	db *gorm.DB
}

func NewManhwaRepository(db *gorm.DB) ManhwaRepository {
	return &manhwaRepository{
		db: db,
	}
}

func (r *manhwaRepository) GetManhwas() ([]*model.Manhwa, error) {
	var manhwas []*model.Manhwa
	if err := r.db.Find(&manhwas).Error; err != nil {
		return nil, err
	}
	return manhwas, nil
}

func (r *manhwaRepository) GetManhwaByTitle(title string) (*model.Manhwa, error) {
	var manhwa model.Manhwa
	if err := r.db.Where("title = ?", title).First(&manhwa).Error; err != nil {
		return nil, err
	}
	return &manhwa, nil
}

func (r *manhwaRepository) CreateManhwa(request *request.CreateManhwaRequest) (*model.Manhwa, error) {

}
