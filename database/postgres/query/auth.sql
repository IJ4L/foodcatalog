-- name: InsertUser :exec
INSERT INTO
  auth (email, password)
VALUES
  ($1, $2);

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