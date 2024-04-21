-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one 
INSERT INTO users (
  email, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMoviesForUser :many
SELECT * FROM movies 
where user_id = $1
Order by "watchedDate" desc;

-- name: CreateMovie :one
INSERT INTO movies (
  user_id, name, "watchedDate", "posterUrl", diary, description, "locationWatched", "releaseDate"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: EditMovie :one 
UPDATE movies 
SET "locationWatched" = $2,
    diary = $3,
    "watchedDate" = $4 
WHERE id = $1 
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies 
WHERE id = $1;