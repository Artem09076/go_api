-- name: CreatePhone :one
INSERT INTO public.phone (country_code, operator, "number", contact_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPhone :one
SELECT * FROM public.phone
WHERE id = $1;

-- name: UpdatePhone :one
UPDATE public.phone
SET
    country_code = $2,
    operator = $3,
    "number" = $4,
    contact_id = $5
WHERE id = $1
RETURNING *;

-- name: DeletePhone :exec
DELETE FROM public.phone
WHERE id = $1;
