package usecase

import (
	v1 "app/v1"
	"app/v1/domain"
	"app/v1/models"
	"app/v1/repository"
	"encoding/json"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type UsecaseHandler struct {
	Repo v1.Repository
}

func NewUsecaseHandler(db *sqlx.DB) *UsecaseHandler {
	return &UsecaseHandler{
		Repo: repository.NewPostgresRepository(db),
	}
}

func (u *UsecaseHandler) Add(item *domain.SongDetail) (*models.Song, error) {
	text, err := json.Marshal(item.Text)
	if err != nil {
		return nil, err
	}

	model := &models.Song{
		Group:       item.Group,
		Song:        item.Song,
		Text:        string(text),
		ReleaseDate: item.ReleaseDate,
		Link:        item.Link,
	}

	if err := u.Repo.Add(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (u *UsecaseHandler) Delete(model *models.Song) error {
	if err := u.Repo.Detele(model); err != nil {
		return err
	}

	return nil
}

func (u *UsecaseHandler) GetAll(song *models.Song, inputCount, inputOffset string) (*[]models.Song, error) {
	count, err := strconv.Atoi(inputCount)
	if err != nil {
		return nil, err
	}

	offset, err := strconv.Atoi(inputOffset)
	if err != nil {
		return nil, err
	}

	songs, err := u.Repo.GetAll(song, count, offset)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (u *UsecaseHandler) Get(model *models.Song, verseID string) (*domain.SongDetail, error) {
	song, err := u.Repo.Get(model)
	if err != nil {
		return nil, err
	}

	if verseID != "" {
		id, err := strconv.Atoi(verseID)
		if err != nil {
			return nil, err
		}

		text := []domain.SongDetailText{}

		if err := json.Unmarshal([]byte(song.Text), &text); err != nil {
			return nil, err
		}

		size := len(text)

		if id > size {
			id = size - 1
		} else if id < 1 {
			id = 0
		} else {
			id -= 1
		}

		return &domain.SongDetail{
			TextString:  text[id].Verse,
			ReleaseDate: song.ReleaseDate,
			Link:        song.Link,
		}, nil

	}

	return &domain.SongDetail{
		TextString:  song.Text,
		ReleaseDate: song.ReleaseDate,
		Link:        song.Link,
	}, nil
}

func (u *UsecaseHandler) Update(item *domain.SongDetail) error {
	text, err := json.Marshal(item.Text)
	if err != nil {
		return err
	}

	model := &models.Song{
		ID:          item.ID,
		Group:       item.Group,
		Song:        item.Song,
		Text:        string(text),
		ReleaseDate: item.ReleaseDate,
		Link:        item.Link,
	}

	if err := u.Repo.Update(model); err != nil {
		return err
	}

	return nil
}
