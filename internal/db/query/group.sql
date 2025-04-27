-- name: CreateGroup :one
INSERT INTO public."group" (title, descriptions)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroup :one
SELECT * FROM public."group"
WHERE id = $1;

-- name: UpdateGroup :one
UPDATE public."group"
SET
    title = $2,
    descriptions = $3
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM public."group"
WHERE id = $1;
