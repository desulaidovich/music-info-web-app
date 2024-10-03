package v1

import "app/v1/models"

//go:generate go run github.com/vektra/mockery/v2@v2.46.0 --name Repository
type Repository interface {
	Add(model *models.Song) error
	Detele(model *models.Song) error
	GetAll(song *models.Song, count, offset int) (*[]models.Song, error)
	Get(song *models.Song) (*models.Song, error)
	Update(song *models.Song) error
}
