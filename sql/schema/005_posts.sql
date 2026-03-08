-- +goose Up
CREATE TABLE posts (
	id uuid not null PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL,
	url TEXT NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
	feed_id uuid not null REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE posts;
