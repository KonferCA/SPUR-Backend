-- name: CreateStartup :one
INSERT INTO startups (
    owner_id,
    name,
    status
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetStartups :many
SELECT * FROM startups
ORDER BY updated_at DESC;