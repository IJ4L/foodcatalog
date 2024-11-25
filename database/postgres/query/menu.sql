-- name: InsertMenu :one
INSERT INTO
  menus (
    name,
    category,
    description,
    price,
    created_at,
    updated_at
  )
VALUES
  ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: SelectAllMenu :many
SELECT
  id,
  name,
  category,
  description,
  price,
  created_at,
  updated_at
FROM
  menus;
  
-- name: SelectMenuByID :one
SELECT
  id,
  name,
  category,
  description,
  price,
  created_at,
  updated_at
FROM
  menus
WHERE
  id = $1;