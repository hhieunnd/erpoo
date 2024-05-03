
-- name: GetTeam :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY name;

-- name: CreateTeam :one
INSERT INTO teams
(id, "name", description)
VALUES($1, $2, $3)
RETURNING *;

-- name: UpdateTeam :one
UPDATE teams
SET "name"=$1, description=$2
WHERE id=$3
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id=$1;