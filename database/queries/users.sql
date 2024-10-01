-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetTotalUser :one
SELECT COUNT(*) FROM users;

-- name: InsertUser :exec
INSERT IGNORE INTO users (email, name, role)
VALUES (?, ?, ?);

-- name: UpdateUserRole :exec
UPDATE users
SET
    role = ?
WHERE
    user_id = ?;
