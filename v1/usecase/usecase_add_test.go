package usecase_test

import (
	"app/v1/domain"
	"app/v1/mocks"
	"app/v1/models"
	"app/v1/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUsecaseHandler_Add(t *testing.T) {
	type args struct {
		item *domain.SongDetail
	}
	tests := []struct {
		name    string
		u       *usecase.UsecaseHandler
		args    args
		want    *models.Song
		wantErr bool
	}{
		{
			name: "add new item",
			args: args{
				item: &domain.SongDetail{
					Group: "Тестовая группа",
					Song:  "Название песни",
					TextString: `[
					]`,
					ReleaseDate: "03.10.2024",
					Link:        "localhost:1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("Add", mock.Anything).Return(nil)

			tt.u = &usecase.UsecaseHandler{
				Repo: repo,
			}

			_, err := tt.u.Add(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseHandler.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
