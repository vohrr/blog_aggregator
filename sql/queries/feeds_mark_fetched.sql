-- name: MarkFeedFetched :exec 
UPDATE feeds 
SET last_fetched_at = NOW(),
created_at = NOW()
WHERE id = $1;

