// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createImageset = `-- name: CreateImageset :one
INSERT INTO imagesets (
  id, name, description, created_at, updated_at
) VALUES (
  uuid_generate_v4(), $1, $2, NOW(), NOW() 
)
RETURNING db_id, id, name, description, created_at, updated_at
`

type CreateImagesetParams struct {
	Name        string
	Description string
}

func (q *Queries) CreateImageset(ctx context.Context, arg CreateImagesetParams) (Imageset, error) {
	row := q.db.QueryRow(ctx, createImageset, arg.Name, arg.Description)
	var i Imageset
	err := row.Scan(
		&i.DbID,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getImageset = `-- name: GetImageset :one
SELECT db_id, id, name, description, created_at, updated_at FROM imagesets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetImageset(ctx context.Context, id uuid.UUID) (Imageset, error) {
	row := q.db.QueryRow(ctx, getImageset, id)
	var i Imageset
	err := row.Scan(
		&i.DbID,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listImageset = `-- name: ListImageset :many
SELECT db_id, id, name, description, created_at, updated_at FROM imagesets
ORDER BY name
`

func (q *Queries) ListImageset(ctx context.Context) ([]Imageset, error) {
	rows, err := q.db.Query(ctx, listImageset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Imageset
	for rows.Next() {
		var i Imageset
		if err := rows.Scan(
			&i.DbID,
			&i.ID,
			&i.Name,
			&i.Description,
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
