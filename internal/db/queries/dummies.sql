-- name: GetDummy :one
SELECT *
FROM dummies
WHERE id = $1;

-- name: GetAllDummies :many
SELECT *
FROM dummies;

-- name: CreateDummy :one
INSERT INTO dummies (name)
VALUES ($1)
RETURNING *;