// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  user_id, name, "watchedDate", "posterUrl", diary, description, "locationWatched", "releaseDate"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, user_id, name, "watchedDate", "posterUrl", diary, description, "locationWatched", "releaseDate"
`

type CreateMovieParams struct {
	UserID          int64
	Name            string
	WatchedDate   	time.Time
	PosterUrl       string
	Diary           string
	Description     string
	LocationWatched string
	ReleaseDate     string
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
		arg.UserID,
		arg.Name,
		arg.WatchedDate,
		arg.PosterUrl,
		arg.Diary,
		arg.Description,
		arg.LocationWatched,
		arg.ReleaseDate,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.WatchedDate,
		&i.PosterUrl,
		&i.Diary,
		&i.Description,
		&i.LocationWatched,
		&i.ReleaseDate,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, name
) VALUES (
  $1, $2
)
RETURNING id, name, email
`

type CreateUserParams struct {
	Email string
	Name  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Name)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies 
WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, id)
	return err
}

const editMovie = `-- name: EditMovie :one
UPDATE movies 
SET "locationWatched" = $2,
    diary = $3,
    "watchedDate" = $4 
WHERE id = $1 
RETURNING id, user_id, name, "watchedDate", "posterUrl", diary, description, "locationWatched", "releaseDate"
`

type EditMovieParams struct {
	ID              int64
	LocationWatched string 
	Diary           string
	WatchedDate     time.Time
}

func (q *Queries) EditMovie(ctx context.Context, arg EditMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, editMovie,
		arg.ID,
		arg.LocationWatched,
		arg.Diary,
		arg.WatchedDate,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.WatchedDate,
		&i.PosterUrl,
		&i.Diary,
		&i.Description,
		&i.LocationWatched,
		&i.ReleaseDate,
	)
	return i, err
}

const getMoviesForUser = `-- name: GetMoviesForUser :many
SELECT id, user_id, name, "watchedDate", "posterUrl", diary, description, "locationWatched", "releaseDate" FROM movies 
where user_id = $1
Order by "watchedDate" desc
`

func (q *Queries) GetMoviesForUser(ctx context.Context, userID int64) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMoviesForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.WatchedDate,
			&i.PosterUrl,
			&i.Diary,
			&i.Description,
			&i.LocationWatched,
			&i.ReleaseDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}
