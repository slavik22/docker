-- name: CreateTutorial :one
INSERT INTO "tutorial" (
    user_id,
    material,
    title
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTutorials :many
SELECT * FROM "tutorial";

-- name: GetTutorial :one
SELECT * FROM "tutorial"
WHERE id = $1;

-- name: GetTutorialsByUser :many
SELECT * FROM "tutorial"
WHERE user_id = $1;

-- name: DeleteTutorial :exec
DELETE FROM "tutorial"
WHERE id = $1;

-- name: UpdateTutorial :one
UPDATE "tutorial"
SET
    title = $2,
    material = $3
WHERE id = $1
    RETURNING *;