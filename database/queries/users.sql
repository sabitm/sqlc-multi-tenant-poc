-- name: GetAllUsers :many
SELECT * FROM xxxx.users;

-- name: GetUserByID :one
SELECT * FROM xxxx.users
WHERE user_id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM xxxx.users
WHERE email = ? LIMIT 1;

-- name: GetTotalUser :one
SELECT COUNT(*) FROM xxxx.users;

-- name: InsertUser :exec
INSERT IGNORE INTO xxxx.users (email, name, role)
VALUES (?, ?, ?);

-- name: UpdateUserRole :exec
UPDATE xxxx.users
SET
    role = ?
WHERE
    user_id = ?;
