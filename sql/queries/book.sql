
-- name: CreateBook :one
INSERT INTO "books" ("category_id", "book_info_id")
VALUES ($1, $2)
RETURNING *;

-- name: ListBook :many
SELECT * FROM "books"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: CountBook :one
SELECT COUNT(*) FROM "books"
WHERE "deleted_at" IS NULL;

-- name: GetOneBookById :one
SELECT * FROM "books"
WHERE "id" = $1 AND "deleted_at" IS NULL;

-- name: UpdateOneBookById :one
UPDATE "books"
SET "category_id" = $2
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;

-- name: DeleteOneBookById :one
UPDATE "books"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;