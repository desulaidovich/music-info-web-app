package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"app/internal/logger"
	"app/internal/render"
	"app/middleware"
	v1 "app/v1"
	"app/v1/domain"
	"app/v1/models"
	"app/v1/usecase"

	"github.com/jmoiron/sqlx"
)

type deleteSongHandler struct {
	uc     v1.Usecase
	logger *logger.Logger
}

func NewDeleteSongHandler(db *sqlx.DB, logger *logger.Logger) *deleteSongHandler {
	return &deleteSongHandler{
		uc:     usecase.NewUsecaseHandler(db),
		logger: logger,
	}
}

// @Summary 	Запрос на удаление песни
// @Tags 		API для музыкальной библиотеки
// @Description	Запрос на удаление песни
// @ID 			delete-song
// @Accept  	json
// @Produce  	json
// @Param 		songID	query 		integer true	"ИД"
// @Success 	200 	{object}	domain.Song		"Без некоторых данных"
// @Failure 	400 	{string} 	string			"Bad request"
// @Failure 	500 	{string} 	string 			"Internal error"
// @Router 		/delete	[delete]
func (h *deleteSongHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := middleware.RequestID(r.Context())

		paramSongID := r.URL.Query().Get("songID")
		id, err := strconv.Atoi(paramSongID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(rid, slog.String("Error", err.Error()), slog.String("URL param", paramSongID))
			return
		}

		model := &models.Song{
			ID: id,
		}

		if err := h.uc.Delete(model); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(rid, slog.String("Error", err.Error()), slog.Any("Model", model))
			return
		}

		if err := render.RenderAs(&domain.Song{
			Song:  model.Song,
			Group: model.Group,
		}, http.StatusOK, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(rid, slog.String("Error", err.Error()),
				slog.Group("Model",
					slog.Int("ID", model.ID),
					slog.String("Song", model.Song),
					slog.String("Group", model.Group),
				),
			)
			return
		}

		h.logger.Info(rid,
			slog.String("handler", h.Name()),
			slog.Int("Song ID", model.ID),
		)
	}
}

func (h *deleteSongHandler) Name() string {
	return "Delete song by ID"
}

func (h *deleteSongHandler) Method() string {
	return http.MethodDelete
}

func (h *deleteSongHandler) Pattern() string {
	return "/delete"
}
