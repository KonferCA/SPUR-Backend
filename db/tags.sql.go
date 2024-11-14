// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tags.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
    name
) VALUES (
    $1
)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateTag(ctx context.Context, name string) (Tag, error) {
	row := q.db.QueryRow(ctx, createTag, name)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteTag, id)
	return err
}

const getTag = `-- name: GetTag :one
SELECT id, name, created_at, updated_at FROM tags
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTag(ctx context.Context, id pgtype.UUID) (Tag, error) {
	row := q.db.QueryRow(ctx, getTag, id)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTags = `-- name: ListTags :many
SELECT id, name, created_at, updated_at FROM tags
ORDER BY name
`

func (q *Queries) ListTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.Query(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
