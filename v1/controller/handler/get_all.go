package handler

import (
	"log/slog"
	"net/http"

	"app/internal/logger"
	"app/internal/render"
	"app/middleware"
	v1 "app/v1"
	"app/v1/domain"
	"app/v1/models"
	"app/v1/usecase"

	"github.com/jmoiron/sqlx"
)

type (
	getAllSongsHandler struct {
		uc     v1.Usecase
		logger *logger.Logger
	}
	response struct {
		Count int            `json:"count"`
		Songs *[]models.Song `json:"songs"`
	}
)

func NewGetAllSongsHandler(db *sqlx.DB, logger *logger.Logger) *getAllSongsHandler {
	return &getAllSongsHandler{
		uc:     usecase.NewUsecaseHandler(db),
		logger: logger,
	}
}

// @Summary 	Все данные библиотеки
// @Tags 		API для музыкальной библиотеки
// @Description	Все данные библиотеки
// @ID 			get-all-songs
// @Accept  	json
// @Produce  	json
// @Param 		count	query 		integer true 						"Кол-во записей"
// @Param 		offset	query 		integer true 						"Смещение по записям"
// @Param 		input 	body 		domain.SongDetail false 			"Фильтрация"
// @Success 	200 	{object}	handler.response{Songs=models.Song}	"Без некоторых данных"
// @Failure 	400 	{string} 	string								"Bad request"
// @Failure 	500 	{string} 	string 								"Internal error"
// @Router 		/		[get]
func (h *getAllSongsHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := middleware.RequestID(r.Context())

		songFiler, err := render.BindAs[domain.SongDetail](r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.String("Body", "Incorrent JSON body"))
			return
		}

		model := &models.Song{
			Group:       songFiler.Group,
			Song:        songFiler.Song,
			ReleaseDate: songFiler.ReleaseDate,
			Link:        songFiler.Link,
		}

		songs, err := h.uc.GetAll(model, r.URL.Query().Get("count"), r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Group("Params",
				slog.String("count", r.URL.Query().Get("count")),
				slog.String("offset", r.URL.Query().Get("offset")),
			))
			return
		}

		if err := render.RenderAs(&response{
			Count: len(*songs),
			Songs: songs,
		}, http.StatusOK, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			h.logger.Debug(id, slog.String("Error", err.Error()), slog.Any("Song array", songs))
			return
		}

		h.logger.Info(id,
			slog.String("handler", h.Name()),
			slog.Int("Count", len(*songs)),
			slog.Any("Songs", songs),
		)
	}
}

func (h *getAllSongsHandler) Name() string {
	return "Get all songs"
}

func (h *getAllSongsHandler) Method() string {
	return http.MethodGet
}

func (h *getAllSongsHandler) Pattern() string {
	return "/"
}
