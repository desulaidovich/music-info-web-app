package usecase_test

import (
	"app/v1/mocks"
	"app/v1/models"
	"app/v1/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUsecaseHandler_GetAll(t *testing.T) {
	type args struct {
		song        *models.Song
		inputCount  string
		inputOffset string
	}
	tests := []struct {
		name    string
		u       *usecase.UsecaseHandler
		args    args
		want    *[]models.Song
		wantErr bool
	}{
		{
			name: "add new item",
			args: args{
				song: &models.Song{
					ID:          0,
					Group:       "Тестовая группа",
					Song:        "Название песни",
					Text:        `[]`,
					ReleaseDate: "03.10.2024",
					Link:        "localhost:1234",
				},
				inputCount:  "10",
				inputOffset: "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("GetAll", mock.Anything, 10, 2).Return(&[]models.Song{}, nil)

			tt.u = &usecase.UsecaseHandler{
				Repo: repo,
			}

			_, err := tt.u.GetAll(tt.args.song, tt.args.inputCount, tt.args.inputOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseHandler.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
