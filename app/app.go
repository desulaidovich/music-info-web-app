package app

import (
	"context"
	"net/http"

	"app/config"
	"app/internal/logger"
	"app/middleware"
	v1 "app/v1"
	"app/v1/controller/handler"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type application struct {
	config *config.Config
}

func NewApp(options ...func(*application)) *application {
	app := &application{}

	for _, option := range options {
		option(app)
	}

	return app
}

func WithConfig(cfg *config.Config) func(*application) {
	return func(a *application) {
		a.config = cfg
	}
}

func WithMigrate(down bool) func(*application) {
	return func(a *application) {
		if a.config.DB == nil {
			panic("env DATABASE_URL not found")
		}

		db, err := goose.OpenDBWithDriver("pgx", a.config.DB.URL)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}()

		cmd := "up"

		if !down {
			cmd = "down"
		}

		if err := goose.RunContext(context.Background(), cmd, db, "./migrations", "postgres"); err != nil {
			panic(err)
		}
	}
}

func (a *application) Run(logger *logger.Logger) error {
	db, err := sqlx.Connect("pgx", a.config.DB.URL)
	if err != nil {
		return err
	}

	mux := mux.NewRouter().StrictSlash(true)

	routes := v1.SignatureList{
		handler.NewAddSongHandler(db, logger),
		handler.NewDeleteSongHandler(db, logger),
		handler.NewGetAllSongsHandler(db, logger),
		handler.NewGetSongHandler(db, logger),
		handler.NewUpdateSongHandler(db, logger),
	}

	for _, rr := range routes {
		mux.
			Methods(rr.Method()).
			Path(rr.Pattern()).
			Name(rr.Name()).
			Handler(rr.Handler())
	}

	s := &http.Server{
		Addr:    ":" + a.config.HTTP.Port,
		Handler: middleware.Logger(mux, logger),
	}

	return s.ListenAndServe()
}
