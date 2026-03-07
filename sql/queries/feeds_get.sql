-- name: GetFeeds :many
SELECT feeds.*, users.Name as user_name FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;
