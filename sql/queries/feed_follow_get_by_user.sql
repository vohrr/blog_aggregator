-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, users.name as user_name, feeds.name as feed_name FROM feed_follows  
	INNER JOIN users ON users.id = feed_follows.user_id
	INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;
