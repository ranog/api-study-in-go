-- name: GetUserByID :one
SELECT * FROM users u WHERE u.id = $1;
