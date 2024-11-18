-- name: InsertUser :exec
INSERT INTO
  auth (email, password, created_at, updated_at)
VALUES
  ($1, $2, $3, $4);

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