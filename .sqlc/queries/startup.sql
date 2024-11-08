-- name: CreateStartup :one
INSERT INTO startups (
    owner_id,
    name,
    status,
    created_at
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetStartups :many
SELECT * FROM startups
ORDER BY created_at DESC;