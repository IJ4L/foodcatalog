-- name: InsertEmployee :one
INSERT INTO
  employees (name, address, nip, created_at, updated_at)
VALUES
  ($1, $2, $3, now(), now()) RETURNING *;

-- name: SelectAllEmployee :many
SELECT
  id,
  name,
  address,
  nip,
  created_at,
  updated_at
FROM
  employees;

-- name: SelectEmployeeByID :one
SELECT
  id,
  name,
  address,
  nip,
  created_at,
  updated_at
FROM
  employees
WHERE
  id = $1;

-- name: RemoveEmployeeByID :exec
DELETE FROM
  employees
WHERE
  id = $1;