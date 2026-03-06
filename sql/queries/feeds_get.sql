-- name: GetFeeds :many
SELECT feeds.*, users.Name as UserName FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;
