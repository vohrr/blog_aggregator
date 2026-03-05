-- +goose Up
CREATE TABLE users (
	id uuid not null PRIMARY KEY,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null,
	name TEXT not null
);

-- +goose Down

DROP TABLE users;
