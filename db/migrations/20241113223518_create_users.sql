-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT,
    name TEXT,
    surname TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
