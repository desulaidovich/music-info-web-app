-- +goose Up
-- +goose StatementBegin
CREATE TABLE Songs (
    id SERIAL PRIMARY KEY,
    "group" TEXT NOT NULL,
    song TEXT NOT NULL,
    "text" JSON NOT NULL,
    release_date TEXT NOT NULL,
    link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Songs CASCADE;
-- +goose StatementEnd
