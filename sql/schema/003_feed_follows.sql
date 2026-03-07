-- +goose Up
CREATE TABLE feed_follows (
	id uuid PRIMARY KEY not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null,
	user_id uuid not null REFERENCES users(id) ON DELETE CASCADE, 
	feed_id uuid not null REFERENCES feeds(id) ON DELETE CASCADE,
	UNIQUE(user_id, feed_id)
);


-- +goose Down
DROP TABLE feed_follows;
