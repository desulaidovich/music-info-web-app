package models

// type SongDetail struct {
// 	ReleaseDate string `json:"releaseDate"`
// 	Text        string `json:"text"`
// 	Patronymic  string `json:"patronymic,omitempty"`
// }

type Song struct {
	ID          int    `db:"id"`
	Group       string `db:"group"`
	Song        string `db:"song"`
	Text        string `db:"text"`
	ReleaseDate string `db:"release_date"`
	Link        string `db:"link"`
}
