package usecase_test

import (
	"app/v1/domain"
	"app/v1/mocks"
	"app/v1/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUsecaseHandler_Update(t *testing.T) {
	type args struct {
		item *domain.SongDetail
	}
	tests := []struct {
		name    string
		u       *usecase.UsecaseHandler
		args    args
		wantErr bool
	}{
		{
			name: "update song",
			args: args{
				&domain.SongDetail{
					ID:          0,
					Group:       "Тестовая группа",
					Song:        "Название песни",
					TextString:  `[]`,
					ReleaseDate: "03.10.2024",
					Link:        "localhost:1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("Update", mock.Anything).Return(nil)

			tt.u = &usecase.UsecaseHandler{
				Repo: repo,
			}

			if err := tt.u.Update(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("UsecaseHandler.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
