-- name: GetAverageColor :one
SELECT * FROM average_colors
WHERE id = $1 LIMIT 1;

-- name: ListAverageColors :many
SELECT * FROM average_colors
ORDER BY file_name;

-- name: CreateAverageColor :one
INSERT INTO average_colors (
  id, imageset_id, file_name, r, g, b, a, created_at, updated_at
) VALUES (
  uuid_generate_v4(), $1, $2, $3, $4, $5, $6, NOW(), NOW() 
)
RETURNING *;
