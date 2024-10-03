package repository

import (
	"app/v1/models"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}

func (r *PostgresRepository) Add(model *models.Song) error {
	rows, err := r.DB.NamedQuery(`INSERT INTO public.songs
		("group", song, "text", release_date, link) VALUES
		(:group, :song, :text, :release_date, :link) RETURNING *;`, model)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err = rows.StructScan(model); err != nil {
			return err
		}
	}

	return nil
}

func (r *PostgresRepository) Detele(model *models.Song) error {
	_, err := r.DB.NamedQuery(`DELETE FROM public.songs WHERE
	id=:id;`, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetAll(song *models.Song, count, offset int) (*[]models.Song, error) {
	songs := []models.Song{}

	filter := map[string]any{
		"group":        song.Group,
		"song":         song.Song,
		"release_date": song.ReleaseDate,
		"offset":       offset,
		"limit":        count,
	}

	nstmt, err := r.DB.PrepareNamed(`SELECT * FROM public.songs WHERE
		id > :offset AND (
			"group"=:group OR
			song=:song OR
			release_date=:release_date
		) IS NOT NULL
		ORDER BY id ASC
		LIMIT :limit;
	`)
	if err != nil {
		return nil, err
	}

	if err := nstmt.Select(&songs, filter); err != nil {
		return nil, err
	}

	return &songs, nil
}

func (r *PostgresRepository) Get(model *models.Song) (*models.Song, error) {
	rows, err := r.DB.NamedQuery(`SELECT * FROM public.songs WHERE
		song=:song AND "group"=:group;`, &model)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.StructScan(&model); err != nil {
			return nil, err
		}
	}

	return model, nil
}

func (r *PostgresRepository) Update(song *models.Song) error {
	_, err := r.DB.NamedQuery(`UPDATE public.songs
		SET
		"group"=:group,
		song=:song,
		"text"=:text,
		release_date=:release_date,
		link=:link
		WHERE id=:id;`, &song)
	if err != nil {
		return err
	}

	return nil
}
