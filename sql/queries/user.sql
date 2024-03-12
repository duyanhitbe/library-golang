
-- name: CreateUser :one
INSERT INTO "users" ("username", "password", "role")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListUser :many
SELECT * FROM "users"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: CountUser :one
SELECT COUNT(*) FROM "users"
WHERE "deleted_at" IS NULL;

-- name: GetOneUserById :one
SELECT * FROM "users"
WHERE "id" = $1 AND "deleted_at" IS NULL;

-- name: UpdateOneUserById :one
UPDATE "users"
SET "username" = $2, "role" = $3
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;

-- name: DeleteOneUserById :one
UPDATE "users"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;