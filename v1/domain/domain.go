package domain

type (
	Song struct {
		ID    int    `json:"songID,omitempty"`
		Group string `json:"group,omitempty"`
		Song  string `json:"song,omitempty"`
	}

	SongDetailText struct {
		Verse  string `json:"verse"`
		Chorus string `json:"chorus"`
	}

	SongDetail struct {
		ID          int              `json:"songID,omitempty"`
		Group       string           `json:"group,omitempty"`
		Song        string           `json:"song,omitempty"`
		Text        []SongDetailText `json:"sourceText,omitempty"`
		TextString  string           `json:"text,omitempty"`
		ReleaseDate string           `json:"releaseDate,omitempty"`
		Link        string           `json:"link,omitempty"`
	}
)
