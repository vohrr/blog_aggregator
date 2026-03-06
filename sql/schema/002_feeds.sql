-- +goose Up
CREATE TABLE feeds (
	id uuid not null PRIMARY KEY,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null,
	name TEXT not null,
	url TEXT unique not null,
	user_id uuid not null REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
