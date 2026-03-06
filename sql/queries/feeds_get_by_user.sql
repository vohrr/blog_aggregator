-- name: GetFeedsForUser :many
SELECT * FROM feeds
WHERE user_id = $1;
