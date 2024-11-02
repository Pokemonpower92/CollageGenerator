// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: collage.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createCollage = `-- name: CreateCollage :one
INSERT INTO collages (
  id, name, description, image_set_id, target_image_id, created_at, updated_at
) VALUES (
  uuid_generate_v4(), $1, $2, $3, $4, NOW(), NOW() 
)
RETURNING db_id, id, name, description, image_set_id, target_image_id, created_at, updated_at
`

type CreateCollageParams struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ImageSetID    uuid.UUID `json:"image_set_id"`
	TargetImageID uuid.UUID `json:"target_image_id"`
}

func (q *Queries) CreateCollage(ctx context.Context, arg CreateCollageParams) (*Collage, error) {
	row := q.db.QueryRow(ctx, createCollage,
		arg.Name,
		arg.Description,
		arg.ImageSetID,
		arg.TargetImageID,
	)
	var i Collage
	err := row.Scan(
		&i.DbID,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageSetID,
		&i.TargetImageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getCollage = `-- name: GetCollage :one
SELECT db_id, id, name, description, image_set_id, target_image_id, created_at, updated_at FROM collages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCollage(ctx context.Context, id uuid.UUID) (*Collage, error) {
	row := q.db.QueryRow(ctx, getCollage, id)
	var i Collage
	err := row.Scan(
		&i.DbID,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageSetID,
		&i.TargetImageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listCollages = `-- name: ListCollages :many
SELECT db_id, id, name, description, image_set_id, target_image_id, created_at, updated_at FROM collages
ORDER BY name
`

func (q *Queries) ListCollages(ctx context.Context) ([]*Collage, error) {
	rows, err := q.db.Query(ctx, listCollages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Collage
	for rows.Next() {
		var i Collage
		if err := rows.Scan(
			&i.DbID,
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageSetID,
			&i.TargetImageID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
