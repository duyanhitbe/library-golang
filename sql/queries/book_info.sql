
-- name: CreateBookInfo :one
INSERT INTO "book_infos" ("name", "author", "publication_date")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListBookInfo :many
SELECT * FROM "book_infos"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: CountBookInfo :one
SELECT COUNT(*) FROM "book_infos"
WHERE "deleted_at" IS NULL;

-- name: GetOneBookInfoById :one
SELECT * FROM "book_infos"
WHERE "id" = $1 AND "deleted_at" IS NULL;

-- name: UpdateOneBookInfoById :one
UPDATE "book_infos"
SET "name" = $2, "author" = $3, "publication_date" = $4
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;

-- name: DeleteOneBookInfoById :one
UPDATE "book_infos"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;