package usecase_test

import (
	"app/v1/domain"
	"app/v1/mocks"
	"app/v1/models"
	"app/v1/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUsecaseHandler_Get(t *testing.T) {
	type args struct {
		model   *models.Song
		verseID string
	}
	tests := []struct {
		name    string
		u       *usecase.UsecaseHandler
		args    args
		want    *domain.SongDetail
		wantErr bool
	}{
		{
			name: "get song info",
			args: args{
				model: &models.Song{
					Group:       "Тестовая группа",
					Song:        "Название песни",
					Text:        `[{"verse":"Без Куплета!","chorus":"Россия, Россия — в этом слове огонь и сила\nВ этом слове победы пламя\nПоднимаем России знамя\nРоссия, Россия — в этом слове огонь и сила\nВ этом слове победы пламя\nПоднимаем России знамя"}]`,
					ReleaseDate: "03.10.2024",
					Link:        "localhost:1234",
				},
				verseID: "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("Get", mock.Anything).Return(tt.args.model, nil)

			tt.u = &usecase.UsecaseHandler{
				Repo: repo,
			}

			_, err := tt.u.Get(tt.args.model, tt.args.verseID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseHandler.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
