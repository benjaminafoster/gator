-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
	$1, 
	$2, 
	$3, 
	$4, 
	$5, 
	$6, 
	$7, 
	$8
) RETURNING *;

-- name: GetPostsByUserID :many
WITH feed_follows_by_user_id AS (
    SELECT
        feed_follows.*
    FROM feed_follows
    INNER JOIN users ON feed_follows.user_id = users.id
    WHERE users.id = $1
)
SELECT
    posts.*
FROM posts
INNER JOIN feed_follows_by_user_id ON posts.feed_id = feed_follows_by_user_id.feed_id
LIMIT $2;