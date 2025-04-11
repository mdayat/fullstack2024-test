-- name: InsertMyClient :one
INSERT INTO my_client (name, slug, is_project, self_capture, client_prefix, client_logo, address, phone_number, city, created_at, updated_at, deleted_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *;

-- name: SelectMyClientById :one
SELECT * FROM my_client WHERE id = $1;

-- name: UpdateMyClient :one
UPDATE my_client
SET
  name = $1,
  slug = $2,
  is_project = $3,
  self_capture = $4,
  client_prefix = $5,
  client_logo = $6,
  address = $7,
  phone_number = $8,
  city = $9,
  updated_at = $10,
  created_at = COALESCE(sqlc.narg(created_at), created_at)
WHERE id = $1 RETURNING *;

-- name: DeleteMyClient :execrows
UPDATE my_client SET deleted_at = $2 WHERE id = $1;
