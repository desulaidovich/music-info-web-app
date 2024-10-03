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
	"app/v1/usecase"

	"github.com/jmoiron/sqlx"
)

type updateSongHandler struct {
	uc     v1.Usecase
	logger *logger.Logger
}

func NewUpdateSongHandler(db *sqlx.DB, logger *logger.Logger) *updateSongHandler {
	return &updateSongHandler{
		uc:     usecase.NewUsecaseHandler(db),
		logger: logger,
	}
}

// @Summary 	Обновить данные песни
// @Tags 		API для музыкальной библиотеки
// @Description	Обновить данные песни
// @ID 			update-song-info
// @Accept  	json
// @Produce  	json
// @Param 		songID	query 		integer true 		"ИД песни"
// @Success 	200 	{object}	domain.SongDetail	"Данные"
// @Failure 	400 	{string} 	string				"Bad request"
// @Failure 	500 	{string} 	string 				"Internal error"
// @Router 		/update	[put]
func (h *updateSongHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := middleware.RequestID(r.Context())

		songID := r.URL.Query().Get("songID")
		id, err := strconv.Atoi(songID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(rid, slog.String("Error", err.Error()), slog.String("Param", songID))
			return
		}

		songReqBody, err := render.BindAs[domain.SongDetail](r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(rid, slog.String("JSON request body", err.Error()))
			return
		}

		song := &domain.SongDetail{
			ID:          id,
			Group:       songReqBody.Group,
			Text:        songReqBody.Text,
			TextString:  songReqBody.TextString,
			Song:        songReqBody.Song,
			ReleaseDate: songReqBody.ReleaseDate,
			Link:        songReqBody.Link,
		}

		if err := h.uc.Update(song); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(rid, slog.String("Error", err.Error()), slog.Any("Song", song))
			return
		}

		if err := render.RenderAs(&map[string]string{
			"song": "updated",
		}, http.StatusOK, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(rid, slog.String("Map error", err.Error()))
			return
		}

		h.logger.Info(rid,
			slog.String("handler", h.Name()),
			slog.Int("Song ID", id),
		)
	}
}

func (h *updateSongHandler) Name() string {
	return "Update song"
}

func (h *updateSongHandler) Method() string {
	return http.MethodPut
}

func (h *updateSongHandler) Pattern() string {
	return "/update"
}
