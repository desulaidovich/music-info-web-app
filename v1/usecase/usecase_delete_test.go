package usecase_test

import (
	"app/v1/mocks"
	"app/v1/models"
	"app/v1/usecase"
	"testing"
)

func TestUsecaseHandler_Delete(t *testing.T) {
	type args struct {
		model *models.Song
	}
	tests := []struct {
		name    string
		u       *usecase.UsecaseHandler
		args    args
		wantErr bool
	}{
		{
			name: "add new item",
			args: args{
				model: &models.Song{
					ID: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		repo := mocks.NewRepository(t)
		repo.On("Detele", &models.Song{}).Return(nil)

		tt.u = &usecase.UsecaseHandler{
			Repo: repo,
		}

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Delete(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("UsecaseHandler.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
