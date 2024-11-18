-- name: InsertEmployee :exec
INSERT INTO
  employees (name, address, nip, created_at, updated_at)
VALUES
  ($1, $2, $3, now(), now());

-- name: SelectAllEmployee :many
SELECT
  id,
  name,
  address,
  nip,
  created_at
FROM
  employees;

-- name: RemoveEmployeeByID :exec
DELETE FROM
  employees
WHERE
  id = $1;