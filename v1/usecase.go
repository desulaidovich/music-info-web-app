package v1

import (
	"app/v1/domain"
	"app/v1/models"
)

type Usecase interface {
	Add(item *domain.SongDetail) (*models.Song, error)
	Delete(model *models.Song) error
	GetAll(song *models.Song, inputCount, inputOffset string) (*[]models.Song, error)
	Get(model *models.Song, verseID string) (*domain.SongDetail, error)
	Update(item *domain.SongDetail) error
}
