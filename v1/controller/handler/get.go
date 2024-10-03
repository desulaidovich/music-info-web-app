package handler

import (
	"log/slog"
	"net/http"

	"app/internal/logger"
	"app/internal/render"
	"app/middleware"
	v1 "app/v1"
	"app/v1/models"
	"app/v1/usecase"

	"github.com/jmoiron/sqlx"
)

type getSongHandler struct {
	uc     v1.Usecase
	logger *logger.Logger
}

func NewGetSongHandler(db *sqlx.DB, logger *logger.Logger) *getSongHandler {
	return &getSongHandler{
		uc:     usecase.NewUsecaseHandler(db),
		logger: logger,
	}
}

// @Summary 	Получить данные по песне
// @Tags 		API для музыкальной библиотеки
// @Description	Получить данные по песне
// @ID 			get-song-info
// @Accept  	json
// @Produce  	json
// @Param 		verse	query 		integer false 		"Номер куплета от 1. Пусто - все куплеты"
// @Param 		group	query 		string  true 		"Название группы"
// @Param 		song	query 		string  true 		"Название песни"
// @Success 	200 	{object}	domain.SongDetail	"Данные"
// @Failure 	400 	{string} 	string				"Bad request"
// @Failure 	500 	{string} 	string 				"Internal error"
// @Router 		/info	[get]
func (h *getSongHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := middleware.RequestID(r.Context())

		verse := r.URL.Query().Get("verse")
		group := r.URL.Query().Get("group")
		song := r.URL.Query().Get("song")

		if group == "" || song == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(id,
				slog.String("Error", "empty params"),
				slog.Group("Params",
					slog.String("Group", group),
					slog.String("Song", song),
				),
			)
			return
		}

		model := &models.Song{
			Group: group,
			Song:  song,
		}

		detail, err := h.uc.Get(model, verse)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Any("Model", model), slog.String("Param verse", verse))
			return
		}

		if err := render.RenderAs(detail, http.StatusOK, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Any("Song", detail))
			return
		}

		h.logger.Info(id,
			slog.String("handler", h.Name()),
			slog.Any("Detail", detail),
		)
	}
}

func (h *getSongHandler) Name() string {
	return "Get song"
}

func (h *getSongHandler) Method() string {
	return http.MethodGet
}

func (h *getSongHandler) Pattern() string {
	return "/info"
}
