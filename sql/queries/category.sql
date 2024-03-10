
-- name: CreateCategory :one
INSERT INTO "categories" ("name")
VALUES ($1)
RETURNING *;

-- name: ListCategory :many
SELECT * FROM "categories"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: CountCategory :one
SELECT COUNT(*) FROM "categories"
WHERE "deleted_at" IS NULL;

-- name: GetOneCategoryById :one
SELECT * FROM "categories"
WHERE "id" = $1 AND "deleted_at" IS NULL;

-- name: UpdateOneCategoryById :one
UPDATE "categories"
SET "name" = $1
WHERE "id" = $2 AND "deleted_at" IS NULL
RETURNING *;

-- name: DeleteOneCategoryById :one
UPDATE "categories"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;