-- name: CreateContact :one
INSERT INTO public.contact (username, given_name, email, birthday_date, group_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetContact :one
SELECT * FROM public.contact
WHERE id = $1;

-- name: UpdateContact :one
UPDATE public.contact
SET
    username = $2,
    given_name = $3,
    email = $4,
    birthday_date = $5,
    group_id = $6
WHERE id = $1
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM public.contact
WHERE id = $1;
