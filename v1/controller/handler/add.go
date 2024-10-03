package handler

import (
	"log/slog"
	"net/http"

	"app/internal/logger"
	"app/internal/render"
	"app/middleware"
	v1 "app/v1"
	"app/v1/domain"
	"app/v1/usecase"

	"github.com/jmoiron/sqlx"
)

type addSongHandler struct {
	uc     v1.Usecase
	logger *logger.Logger
}

func NewAddSongHandler(db *sqlx.DB, logger *logger.Logger) *addSongHandler {
	return &addSongHandler{
		uc:     usecase.NewUsecaseHandler(db),
		logger: logger,
	}
}

// @Summary 	Запрос на добавление новой песни
// @Tags 		API для музыкальной библиотеки
// @Description	Запрос на добавление новой песни
// @ID 			add-new-song
// @Accept  	json
// @Produce  	json
// @Param 		input 	body 		domain.SongDetail true	"Данные по песне"
// @Success 	200 	{object} 	domain.Song 			"ИД"
// @Failure 	400 	{string} 	string					"Bad request"
// @Failure 	500 	{string} 	string 					"Internal error"
// @Router 		/add	[post]
func (h *addSongHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := middleware.RequestID(r.Context())

		song, err := render.BindAs[domain.SongDetail](r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Any("Song model", song))
			return
		}

		model, err := h.uc.Add(song)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Any("Model", model))
			return
		}

		if err := render.RenderAs(&domain.Song{
			ID: model.ID,
		}, http.StatusOK, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Int("Render model id", model.ID))
			return
		}

		h.logger.Info(id,
			slog.String("handler", h.Name()),
			slog.Int("Song ID", model.ID),
		)
	}
}

func (h *addSongHandler) Name() string {
	return "Add"
}

func (h *addSongHandler) Method() string {
	return http.MethodPost
}

func (h *addSongHandler) Pattern() string {
	return "/add"
}
