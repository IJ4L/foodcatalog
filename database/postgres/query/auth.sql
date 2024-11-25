-- name: InsertUser :one
INSERT INTO
  auth (email, password)
VALUES
  ($1, $2) RETURNING id, email, created_at, updated_at;

-- name: SelectUserByEmail :one
SELECT
  id,
  email,
  password,
  created_at,
  updated_at
FROM
  auth
WHERE
  email = $1;